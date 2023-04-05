package lwdx

import "fmt"

// https://docs.oasis-open.org/dita/dita/v1.3/os/part1-base/non-normative/elementsMerged.html

var AllXditaTags map[string]BlockInlineOther

func init() {
	AllXditaTags = make(map[string]BlockInlineOther)

	AllXditaBlockTags = make(map[string]void)
	for _, s := range allXditaBlockTags {
		AllXditaBlockTags[s] = aVoid
		AddToAllXditaTags(s, BIO(true, false))
	}
	AllXditaInlineTags = make(map[string]void)
	for _, s := range allXditaInlineTags {
		AllXditaInlineTags[s] = aVoid
		AddToAllXditaTags(s, BIO(false, true))
	}
}

func AddToAllXditaTags(tag string, bio BlockInlineOther) {
	prevBio, ok := AllXditaTags[tag]
	if !ok {
		// Not in the map yet
		AllXditaTags[tag] = bio
	} else {
		// Is already in map
		// Is value a dupe ?
		if bio.IsBlock == prevBio.IsBlock &&
			bio.IsInline == prevBio.IsInline {
			fmt.Printf("Xdita dupeVal <%s>\t %t,%t \n",
				tag, bio.IsBlock, prevBio.IsBlock)
		} else {
			fmt.Printf("Xdita conflict <%s>\t B,%t,%t I,%t,%t \n", tag,
				bio.IsBlock, prevBio.IsBlock,
				bio.IsInline, prevBio.IsInline)
		}
	}
}

/*
<-- lw-topic.mod: common content models (" <!ENTITY % ") -->
simple-blox  "p|ul|ol|dl| pre|audio|video| example| note|%data;">
flbk-blox    "p|ul|ol|dl| pre|note|image|alt|%data;">
fn-block     "p|ul|ol|dl| %data;">
all-blox     "p|ul|ol|dl| pre|audio|video| example| simpletable| fig|note|%data;">
list-blox    "p|ul|ol|dl| pre|audio|video| example| simpletable| fig|note|%data;">
fig-blox     "p|ul|ol|dl| pre|audio|video| example| simpletable| %data;">
example-blox "p|ul|ol|dl| pre|audio|video|          simpletable| fig|note|%data;">
*/

var AllXditaBlockTags map[string]void
var allXditaBlockTags = []string{
	"alt", "body", "dd", "desc", "dl", "div", "dlentry", "dt",
	"example", "fig", "fn", "image", "keydef", "li", "linktext",
	"map", "navtitle", "note", "ol", "p", "pre", "prolog",
	"section", "shortdesc", "simpletable", "sthead", "strong", "strow",
	"title", "topic", "topicmeta", "topicref", "ul",
}

var AllXditaInlineTags map[string]void
var allXditaInlineTags = []string{
	"b", "em", "i" /*"image",*/, "ph",
	"sub", "sup", "u", "xref",
}

var AllXditaHiddenTags map[string]void

/*
ALL reused elements:
https://lists.oasis-open.org/archives/dita-lightweight-dita/202210/msg00004/dita-lw-dita-reuse.pdf
*/
