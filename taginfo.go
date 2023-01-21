package lwdx

// TODO: Use Go EMBED directive for DTDs and LwDITA samples

import (
	"fmt"
	"sort"
	// SU "github.com/fbaube/stringutils"
)

/*
Tags with differing Modes:

map: html = image map, lwdita = ToC
body: lwdita = topic
video:
*/

type TagMode string

var TagModes = []TagMode{
	"BLCK",
	"INLN",
	"BOTH", // Both
	"EMBD", // Embedded, i.e. not rendered, or N/A
}

func (tm TagMode) S() string {
	return string(tm)
}

func (tm TagMode) String() string {
	return string(tm)
}

func (tm TagMode) IsBlock() bool {
	if tm == "BLCK" || tm == "BOTH" {
		return true
	}
	if tm == "INLN" || tm == "EMBD" {
		return false
	}
	panic("lwx.IsBlock:" + tm.S())
}

func (tm TagMode) IsInline() bool {
	if tm == "INLN" || tm == "BOTH" {
		return true
	}
	if tm == "BLCK" || tm == "EMBD" {
		return false
	}
	panic("lwx.IsInline:" + tm.S())
}

func (tm TagMode) IsEmbed() bool {
	return tm == "EMBD"
}

// TagSummary is a set of booleans that quickly characterizes a tag, no
// matter what kind of "common" XML file it is found in - HTML5, LwDITA,
// maybe also DITA. There's a bit of mix & match going on between LwDITA
// and HTML5, and we want to be liberal about accepting near-misses (like
// B & I v EMPH & STRONG), so this approach kinda makes sense.
// .
type TagSummary struct {
	// IsHtml5  bool
	// IsLwdita bool
	// TagMode
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
	TagInfo["test1"] = TagSummary{"BOTH", "BOTH"}
	TagInfo["test2"] = TagSummary{"BLCK", "INLN"}
	setSchemaAndBLKorINL(HtmlBlockTags, false, "BLCK")
	setSchemaAndBLKorINL(HtmlInlineTags, false, "INLN")
	setSchemaAndBLKorINL(HtmlEmbedTags, false, "EMBD")
	setSchemaAndBLKorINL(LwDitaBlockTags, true, "BLCK")
	setSchemaAndBLKorINL(LwDitaInlineTags, true, "INLN")
	setSchemaAndBLKorINL(LwDitaModelessTags, true, "INLN")

	var BLCKs, INLNs, BOTHs, EMBDs []string
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
			fmt.Printf("DISagreement: <%s> H:%s L:%s \n",
				tag, Hmode, Lmode)
			continue
		} else {
			panic(fmt.Sprintf("oh shit:<%s>  H:%s L:%s",
				tag, Hmode, Lmode))
		}
		var s string
		s = (bothMode.S() + soleMode.S())
		if len(s) < 4 || len(s) > 6 {
			panic(fmt.Sprintf("oh shit: <%s> H:%s L:%s 2:%s 1:%s",
				tag, Hmode, Lmode, bothMode, soleMode))
		}
		tag += " "
		switch s[len(s)-4:] {
		case "BLCK":
			BLCKs = append(BLCKs, tag)
		case "INLN":
			INLNs = append(INLNs, tag)
		case "BOTH":
			BOTHs = append(BOTHs, tag)
		case "EMBD":
			EMBDs = append(EMBDs, tag)
		case "":
			fmt.Printf("DISagreement on: %s \n", tag)
		default:
			fmt.Printf("WTF: <%s> both<%s> sole<%s> \n",
				tag, bothMode, soleMode)
		}
	}
	sort.Strings(BLCKs)
	sort.Strings(INLNs)
	sort.Strings(BOTHs)
	sort.Strings(EMBDs)
	fmt.Printf("BLOX: %v \n", BLCKs)
	fmt.Printf("INLN: %v \n", INLNs)
	fmt.Printf("BOTH: %v \n", BOTHs)
	fmt.Printf("EMBD: %v \n", EMBDs)
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
				fmt.Printf("Tag ERROR: " +
					"multiple Html5 entries for: " + s)
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
				// panic("FAIL on BLK v INL: " + s)
				fmt.Printf("Tag OOPS: <%s>: "+
					"Html5 <%s> != LwDITA <%s> \n",
					s, TS.Html5Mode, TS.LwditaMode)
			}
		}
		// TS.TagMode = TagMode(tagMode)
		TagInfo[s] = TS
	}
}

// An inline element cannot contain a block-level element!

// == HTML ==

// HtmlBlockTags is TODO
var HtmlBlockTags = []string{
	"address", "article", "aside", "blockquote", "body",
	"canvas", "dd", "div", "dl", "dt", "fieldset",
	"figcaption", "figure", "footer", "form",
	"h1", "h2", "h3", "h4", "h5", "h6", /* "head", */
	"header", "hr" /* "html", */, "li", "main", "nav",
	"noscript", "ol", "output", "p", "pre", "section",
	"table", "tfoot" /* "title", */, "ul", "video"}

// HtmlInlineTags is TODO
var HtmlInlineTags = []string{
	"a", "abbr", "acronym", "b", "bdo",
	"big", "br", "button", "cite", "code",
	"dfn", "em", "i", "img", "input",
	"kbd", "label" /* "link", */, "map",
	"object", /* "output", */
	"q", "samp", "script", "select", "small",
	"span", "strong", "sub", "sup", "textarea",
	"time", "tt", "var"}

// HtmlEmbedTags is TODO
var HtmlEmbedTags = []string{
	"html", "head", "link", "title"}

// HtmlSelfClosingTags is TOTO
var HtmlSelfClosingTags = []string{
	"area", "base", "br", "col", "embed", "hr", "img", "input",
	"link", "meta", "param", "source", "track", "wbr"}

// == LwDITA ==

//  "To reuse block-level content, authors will use @conref.
// For phrase-level content, authors will use @keyref."

// LwDitaBlockTags is TODO
var LwDitaBlockTags = []string{
	"dd", "dlentry", "dl", "dt", "fig", "li", "map", "ol", "p", "pre",
	"section", "shortdesc", "simpletable", "title", "topic", "ul"}

// LwDitaInlineTags is TODO
var LwDitaInlineTags = []string{
	"b", "i", "ph", "sup", "sub", "u",
	"xref", "fn", "image", "linktext", "span", "topicref"}

// LwDitaModelessTags is TODO
var LwDitaModelessTags = []string{
	"alt", "body", "data", "desc", "keydef", "note", "navtitle",
	"prolog", "stentry", "sthead", "strow", "topicmeta",
	// XML elements in LwDITA that are new
	"audio", "media-autoplay", "media-controls", "media-loop",
	"media-muted", "video-poster", "media-source", "media-track",
	"video"}

// var TTinline = TagSummary{false, false, "INLN", false, false}
// var TTblock = TagSummary{false, false, "BLCK", false, false}
// var TTembed = TagSummary{false, false, "EMBD", false, false}

// TagTypeMapper is a type created for var gxml.TagTypes
type TagTypeMapper map[string]TagSummary

/*

// PredefinedTagTypes is a singleton for quick characterization
// of all LwDITA tags and common HTML5 tags.
var PredefinedTagTypes = TagTypeMapper{
	// INLINE TAGS
	"b":      {false, false, "INLN", false, false},
	"i":      {false, false, "INLN", false, false},
	"u":      {false, false, "INLN", false, false},
	"em":     {false, false, "INLN", false, false},
	"strong": {false, false, "INLN", false, false},
	"ph":     {false, false, "INLN", false, false},
	"span":   {false, false, "INLN", false, false},
	// BLOCK TAGS
	"p":         {false, false, "BLCK", false, false},
	"topic":     {false, false, "BLCK", false, false},
	"title":     {false, false, "BLCK", false, false},
	"shortdesc": {false, false, "BLCK", false, false},
	"body":      {false, false, "BLCK", false, false},
	"section":   {false, false, "BLCK", false, false},
	"ul":        {false, false, "BLCK", false, false},
	"ol":        {false, false, "BLCK", false, false},
	"li":        {false, false, "BLCK", false, false},
	"taskbody":  {false, false, "BLCK", false, false},
	"context":   {false, false, "BLCK", false, false},
	"prereq":    {false, false, "BLCK", false, false},
	"steps":     {false, false, "BLCK", false, false},
	"step":      {false, false, "BLCK", false, false},
	"table":     {false, false, "BLCK", false, false},
	"hr":        {false, false, "BLCK", false, false},
	"h1":        {false, false, "BLCK", false, false},
	"h2":        {false, false, "BLCK", false, false},
	"h3":        {false, false, "BLCK", false, false},
	"h4":        {false, false, "BLCK", false, false},
	"h5":        {false, false, "BLCK", false, false},
	"author":    {false, false, "BLCK", false, false},
	"keydef":    {false, false, "BLCK", false, false},
	"keyword":   {false, false, "BLCK", false, false},
	"keywords":  {false, false, "BLCK", false, false},
	"map":       {false, false, "BLCK", false, false},
	"topicmeta": {false, false, "BLCK", false, false},
	"topicref":  {false, false, "BLCK", false, false},
	"navtitle":  {false, false, "BLCK", false, false},
}

*/
