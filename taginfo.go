package lwdx

// TODO: Use Go EMBED directive for DTDs and LwDITA samples

import (
	"fmt"
)

// Tags with differing Modes:
// map: html = image map, lwdita = ToC
// body: lwdita = topic
// video:

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
