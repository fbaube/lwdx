package lwdx

// NOTE:410 TODO Use EMBED !!! For Image, Video, and maybe all the Modeless

import (
	"fmt"

	SU "github.com/fbaube/stringutils"
)

// Package tags describes LwDITA tags and also many HTML5 tags.

type TagMode string

var TagModes = []TagMode{
	"BLCK",
	"INLN",
	"BOTH",
	"EMBD",
}

func (tm TagMode) IsBlock() bool {
	return tm == "BLCK" || tm == "BOTH"
}

func (tm TagMode) IsInline() bool {
	return tm == "INLN" || tm == "BOTH"
}

func (tm TagMode) IsEmbed() bool {
	return tm == "EMBD"
}

// TagSummary is a set of booleans that quickly characterizes a tag,
// no matter what kind of "common" XML file it is found in - HTML5,
// LwDITA, also ?DITA. There's a bit of mix & match going on between
// LwDITA and HTML5, so this approach kinda makes sense.
type TagSummary struct {
	IsHtml5  bool
	IsLwdita bool
	TagMode
	IsSelfclosing bool
	IsHidden      bool
}

// TagTypes is a singleton for quick characterization
// of all LwDITA tags and common HTML5 tags.
var TagSummaries TagTypeMapper

// init uses the slices of tags defined below to initialize descriptors.
func init() {
	TagSummaries = make(map[string]TagSummary)
	TagSummaries["test1"] = TagSummary{true, true, "BOTH", true, false}
	TagSummaries["test2"] = TagSummary{true, true, "BOTH", false, false}
	setSchemaAndIsBlockOrInline(HtmlBlockTags, true, false, "BLCK")
	setSchemaAndIsBlockOrInline(HtmlInlineTags, true, false, "INLN")
	setSchemaAndIsBlockOrInline(LwDitaBlockTags, false, true, "BLCK")
	setSchemaAndIsBlockOrInline(LwDitaInlineTags, false, true, "INLN")
	setSchemaAndIsBlockOrInline(LwDitaModelessTags, false, true, "INLN")
	/*
		for k, v := range TagTypes {
			fmt.Fprintf(os.Stdout, "%s:      \t%s \n", k, v)
		}
	*/
}

func (TT TagSummary) String() string {
	return fmt.Sprintf("%s,Html:%s,LwDita:%s",
		TT.TagMode, SU.Yn(TT.IsHtml5), SU.Yn(TT.IsLwdita))
}

func setSchemaAndIsBlockOrInline(tags []string,
	isHtml bool, isLwd bool, tagMode string) {
	var TT TagSummary
	var ok bool
	for _, s := range tags {
		if TT, ok = TagSummaries[s]; !ok {
			TT = *new(TagSummary)
		}
		if isHtml {
			TT.IsHtml5 = true
		}
		if isLwd {
			TT.IsLwdita = true
		}
		TT.TagMode = TagMode(tagMode)
		TagSummaries[s] = TT
	}
}

// HtmlBlockTags is TODO
var HtmlBlockTags = []string{
	"address", "article", "aside", "blockquote", "body", "canvas", "dd",
	"div", "dl", "dt", "fieldset", "figcaption", "figure", "footer", "form",
	"h1", "h2", "h3", "h4", "h5", "h6", "head", "header", "hr", "html", "li",
	"main", "nav", "noscript", "ol", "output", "p", "pre", "section", "table",
	"tfoot", "title", "ul", "video"}

// HtmlInlineTags is TODO
var HtmlInlineTags = []string{
	"a", "abbr", "acronym", "b", "bdo", "big", "br", "button",
	"cite", "code", "dfn", "em", "i", "img", "input", "kbd",
	"label", "link", "map", "object", "q", "samp", "script",
	"select", "small", "span", "strong", "sub", "sup",
	"textarea", "time", "tt", "var"}

// HtmlSelfClosingTags is TOTO
var HtmlSelfClosingTags = []string{
	"area", "base", "br", "col", "embed", "hr", "img", "input",
	"link", "meta", "param", "source", "track", "wbr"}

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

var TTinline = TagSummary{false, false, "INLN", false, false}
var TTblock = TagSummary{false, false, "BLCK", false, false}
var TTembed = TagSummary{false, false, "EMBD", false, false}

// TagTypeMapper is a type created for var gxml.TagTypes
type TagTypeMapper map[string]TagSummary

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

