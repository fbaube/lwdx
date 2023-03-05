package lwdx

// TODO: Use Go EMBED directive for DTDs and LwDITA samples

// Tags with differing Modes:
// map: html = image map, lwdita = ToC
// body: lwdita = topic
// video:

// An inline element cannot contain a block-level element!

// Note that now <a ...> can wrap a lot.
// http://html5doctor.com/block-level-links-in-html-5/
// One new and exciting thing you can do in HTML 5 is
// wrap links round “block-level” elements.

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
