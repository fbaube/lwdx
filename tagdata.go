package lwdx

// TODO: Use Go EMBED directive for DTDs and LwDITA samples

// TagTypeMapper is a type created for var gxml.TagTypes
type TagTypeMapper map[string]TagSummary

// Tags with differing Modes:
// map: html = image map, lwdita = ToC
// body: lwdita = topic
// video:

// An inline element cannot contain a block-level element!

// Note that now <a ...> can wrap a lot.
// http://html5doctor.com/block-level-links-in-html-5/
// One new and exciting thing you can do in HTML 5 is
// wrap links round “block-level” elements.

// == HTML ==

// HtmlBlockTags is tags that we're pretty sure are block.
// That means they start a new line and occupy the full
// width. It's also a good way to organize content into
// a manageable number of maningful blocks, and points
// in the direction of something like Notion.
// .
var HtmlBlockTags = []string{
	"address", "article", "aside", "blockquote",
	"dd", "details", "dialog", "div", "dl", "dt",
	"fieldset", "figcaption", "figure", "footer",
	"form", "h1", "h2", "h3", "h4", "h5", "h6",
	"header", "hgroup", "hr", "iframe", "li",
	"main", "nav", "ol", "p", "pre", "section",
	"table", "tfoot", "ul", "video"}

// HtmlInlinTags is tags that render something visible
// but are not block.
// .
var HtmlInlinTags = []string{
	"a", "abbr", "acronym", "b", "bdo", "big",
	"br", "button", "canvas", "cite", "code",
	"col", "command", "del", "dfn", "em",
	"embed", "i", "img", "input", "ins",
	"kbd", "label", "map", "mark", "math",
	"menu", "meter", "object", "output",
	"picture", "progress", "q", "ruby", "s",
	"samp", "select", "slot", "small", "span",
	"strong", "sub", "sup", "svg", "textarea",
	"time", "tt", "u", "var", "wbr",
}

// HtmlOtherTags are tags that we shouldn't have
// to worry about when rendering a page.
// .
var HtmlOtherTags = []string{
	"area", "base", "body", "data", "datalist", "html",
	"head", "link", "meta", "noscript", "param", "script",
	"source", "style", "template", "title", "track"}

// HtmlSelfClosingTags tags are all listed elsewhere
// (mostly as "other"). We should not have to worry
// about them, because we expect the Go HTML parser
// to handle this nitpicky stuff.
// .
var HtmlSelfClosingTags = []string{
	"area", "base", "br", "col", "embed", "hr", "img", "input",
	"link", "meta", "param", "source", "track", "wbr"}

// == LwDITA ==

// "To reuse block-level content, authors will use @conref.
// To reuse phrase-level content, authors will use @keyref."

// LwditaBlockTags is TBS.
// .
var LwditaBlockTags = []string{
	"dd", "dlentry", "dl", "dt", "fig", "li",
	"map", "ol", "p", "pre", "section", "shortdesc",
	"simpletable", "title", "topic", "ul", "video"}

// LwditaInlinTags is TBS
// .
var LwditaInlinTags = []string{
	"b", "i", "ph", "sup", "sub", "u", "xref", "fn",
	"image", "linktext", "span", "topicref"}

// LwditaOtherTags is TBS
// .
var LwditaOtherTags = []string{
	"alt", "body", "data", "desc", "keydef", "note", "navtitle",
	"prolog", "stentry", "sthead", "strow", "topicmeta",
	// XML elements in LwDITA that are new
	"audio", "media-autoplay", "media-controls",
	"media-loop", "media-muted", "video-poster",
	"media-source", "media-track" /* ,"video" */}

// var TTinline= TagSummary{false, false, "INLN", false, false}
// var TTblock = TagSummary{false, false, "BLCK", false, false}
// var TTembed = TagSummary{false, false, "EMBD", false, false}
