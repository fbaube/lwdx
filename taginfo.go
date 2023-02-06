package lwdx

// TODO: Use Go EMBED directive for DTDs and LwDITA samples

import (
	"fmt"
	"sort"
	// SU "github.com/fbaube/stringutils"
)

// Tags with differing Modes:
// map: html = image map, lwdita = ToC
// body: lwdita = topic
// video:

type TagMode string

func GetTagSummaryByTagName(s string) *TagSummary {
	if s == "" {
		return nil
	}
	var TS TagSummary
	var ok bool
	TS, ok = TagInfo[s]
	if !ok {
		return nil
	}
	return &TS
}

func (tm TagMode) S() string {
	return string(tm)
}

func (tm TagMode) String() string {
	return string(tm)
}

// TagSummary is a set of booleans that quickly characterizes a tag, no
// matter what kind of "common" XML file it is found in - HTML5, LwDITA,
// maybe also DITA. There's a bit of mix & match going on between LwDITA
// and HTML5, and we want to be liberal about accepting near-misses (like
// B & I v EMPH & STRONG), so this approach kinda makes sense.
// .
type TagSummary struct {
	Html5Mode  TagMode
	LwditaMode TagMode
	// IsSelfClsg bool // self-closing, like <br/>
	// IsHidden   bool
}

// TagTypes is a singleton for quick characterization
// of all LwDITA tags and common HTML5 tags (not the
// exotic ones, whose handling will be ignored for now).
// .
var TagInfo TagTypeMapper

// init uses the slices of tags defined below to initialize descriptors.
// .
func init() {
	TagInfo = make(map[string]TagSummary)
	setSchemaAndBLKorINL(HtmlBlockTags, false, "BLOCK")
	setSchemaAndBLKorINL(HtmlInlinTags, false, "INLIN")
	setSchemaAndBLKorINL(HtmlOtherTags, false, "OTHER")
	setSchemaAndBLKorINL(LwditaBlockTags, true, "BLOCK")
	setSchemaAndBLKorINL(LwditaInlinTags, true, "INLIN")
	setSchemaAndBLKorINL(LwditaOtherTags, true, "OTHER")

	var BLOCKs, INLINs, OTHERs []string
	for tag, v := range TagInfo {
		Hmode := v.Html5Mode
		Lmode := v.LwditaMode
		if Hmode == "" && Lmode == "" {
			panic("BLEAGH")
		}
		var bothMode, soleMode TagMode
		bothMode = ""
		soleMode = ""
		// fmt.Printf("DBG: %s: \t%s \t%s \n", tag, Hmode, Lmode)
		// fmt.Printf("%s:     \t %s \n", tag, v.S())
		if Hmode == Lmode {
			bothMode = Hmode
		} else if Lmode == "" {
			soleMode = Hmode
			tag = "H:" + tag
		} else if Hmode == "" {
			soleMode = Lmode
			tag = "L:" + tag
		} else if Hmode != Lmode {
			fmt.Printf("Tag <%s> disagreement: H:%s L:%s \n",
				tag, Hmode, Lmode)
			continue
		} else {
			panic(fmt.Sprintf("oh shit:<%s>  H:%s L:%s",
				tag, Hmode, Lmode))
		}
		var s string
		s = (bothMode.S() + soleMode.S())
		if len(s) < 5 || len(s) > 7 {
			panic(fmt.Sprintf("oh shit: <%s> H:%s L:%s 2:%s 1:%s",
				tag, Hmode, Lmode, bothMode, soleMode))
		}
		tag += " "
		switch s[len(s)-5:] {
		case "BLOCK":
			BLOCKs = append(BLOCKs, tag)
		case "INLIN":
			INLINs = append(INLINs, tag)
		case "OTHER":
			OTHERs = append(OTHERs, tag)
		case "":
			// fmt.Printf("DISagreement on: %s \n", tag)
			panic("OOPS: " + tag)
		default:
			fmt.Printf("WTF: <%s> both<%s> sole<%s> \n",
				tag, bothMode, soleMode)
		}
	}
	sort.Strings(BLOCKs)
	sort.Strings(INLINs)
	sort.Strings(OTHERs)
	// fmt.Printf("BLOCK: %v \n", BLOCKs)
	// fmt.Printf("INLIN: %v \n", INLINs)
	// fmt.Printf("OTHER: %v \n", OTHERs)
}

func (TS TagSummary) String() string {
	return fmt.Sprintf("html5<%s>lwdita<%s>",
		TS.Html5Mode, TS.LwditaMode)
}

var dupedTags []string
var collidingTags []*TagSummary

func setSchemaAndBLKorINL(tags []string, isLwdita bool, tagMode TagMode) {
	var TS TagSummary
	var ok bool
	if tagMode == "" {
		panic("NO TAG MODE")
	}
	for _, s := range tags {
		if TS, ok = TagInfo[s]; !ok {
			// The tag is not listed in the map yet.
			TS = *new(TagSummary)
			// Is okay to insert
			if isLwdita {
				TS.LwditaMode = tagMode
			} else {
				TS.Html5Mode = tagMode
			}
		} else {
			// The tag IS already listed in the map, as
			// var TS, presumably in another XML schema
			// fmt.Printf("Tag is duped: %s \n", s)
			dupedTags = append(dupedTags, s)
			// Check for multiple in same schema
			if (TS.LwditaMode != "") && isLwdita {
				fmt.Printf("Tag ERROR: " +
					"multiple LwDITA entries for: " + s)
			}
			if (TS.Html5Mode != "") && !isLwdita {
				fmt.Printf("Tag ERROR: "+
					"multiple Html5 entries for: %s \n", s)
			}
			// Is okay to insert
			if isLwdita {
				TS.LwditaMode = tagMode
			} else {
				TS.Html5Mode = tagMode
			}
			// Check for [dis]agreement on TagMode.
			if (TS.Html5Mode != "") &&
				(TS.LwditaMode != "") &&
				(TS.Html5Mode != TS.LwditaMode) {
				fmt.Printf("Tag <%s>: "+
					"Html5 <%s> != LwDITA <%s> \n",
					s, TS.Html5Mode, TS.LwditaMode)
			}
		}
		// TS.TagMode = TagMode(tagMode)
		TagInfo[s] = TS
	}
}
