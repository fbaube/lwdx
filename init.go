package lwdx

// TODO: Use Go EMBED directive for DTDs and LwDITA samples

import (
	"fmt"
	"sort"
	// SU "github.com/fbaube/stringutils"
)

// init uses the slices of tags defined
// elsewhere to initialize descriptors.
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
