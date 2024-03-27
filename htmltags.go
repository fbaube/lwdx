package lwdx

import "fmt"

// MAIN REFERENCE:
// https://html.spec.whatwg.org/

// var's defined in this file:
// var AllHtmlBlockTags  map[string]void
// var AllHtmlInlineTags map[string]void
// var AllHtmlHiddenTags map[string]void
// var AllHtmlVoidTags   map[string]void
// var AllHtmlEmbeddedContentTags    []string
// var AllHtmlInteractiveContentTags []string

var AllHtmlTags map[string]BlockInlineOther

func init() {
	AllHtmlTags = make(map[string]BlockInlineOther)

	AllHtmlBlockTags = make(map[string]void)
	for _, s := range allHtmlBlockTags {
		AllHtmlBlockTags[s] = aVoid
		AddToAllHtmlTags(s, BIO(true, false))
	}
	AllHtmlInlineTags = make(map[string]void)
	for _, s := range allHtmlInlineTags {
		AllHtmlInlineTags[s] = aVoid
		AddToAllHtmlTags(s, BIO(false, true))
	}
	AllHtmlHiddenTags = make(map[string]void)
	for _, s := range allHtmlHiddenTags {
		AllHtmlHiddenTags[s] = aVoid
		AddToAllHtmlTags(s, BIO(false, false))
	}
}

func AddToAllHtmlTags(tag string, bio BlockInlineOther) {
	prevBio, ok := AllHtmlTags[tag]
	if !ok {
		// Not in the map yet
		AllHtmlTags[tag] = bio
	} else {
		// Is already in map
		// Is value a dupe ?
		if bio.IsBlock == prevBio.IsBlock &&
			bio.IsInline == prevBio.IsInline {
			fmt.Printf("Html dupeVal <%s>\t %t,%t \n",
				tag, bio.IsBlock, prevBio.IsBlock)
		} else {
			fmt.Printf("Html conflict <%s>\t B,%t,%t I,%t,%t \n", tag,
				bio.IsBlock, prevBio.IsBlock,
				bio.IsInline, prevBio.IsInline)
		}
	}
}

// Two general rules:
//  - An inline element cannot contain a block-level element.
//  - Most elements categorized as phrasing content [i.e. inline]
//    can only contain elements that are themselves categorized
//    as phrasing content, not any flow content.

// ==========
// Some notes
// ==========
//
// About links:
// Note that now <a ...> can wrap a lot.
// http://html5doctor.com/block-level-links-in-html-5/
// One new and exciting thing you can do in HTML 5 is
// wrap links round “block-level” elements.
//
// About CSS "display: inline-block"
// Unlike inline, "display: inline-block" lets
// you set a width and height on the element.
// Unlike block, "display: inline-block" does
// not add line-break after the element, so
// the element can sit next to other elements.
//
// About <li>:
// WhatWG: An li element's end tag can be omitted if
//  - another li element immediately follows it, or
//  - there is no more content in the parent element
//
// Tags with differing Modes:
// map: html = image map, lwdita = ToC
// body: lwdita = topic
// video:

// = = = = = = = = =

// =======
//  BLOCK
// =======

// AllHtmlBlockTags are in a few places:
//   - https://html.spec.whatwg.org/#the-page
//     html, body { display: block; }
//   - https://html.spec.whatwg.org/#sections-and-headings
//     article, aside, h1, h2, h3, h4, h5, h6, hgroup, nav, section {
//     display: block; }
//   - https://html.spec.whatwg.org/#flow-content-3
//     address, blockquote, center, dialog, div, figure,
//     figcaption, footer, form, header, hr, legend,
//     listing, main, p, plaintext, pre, search, xmp {
//     display: block; }
//   - https://html.spec.whatwg.org/#lists
//     dir, dd, dl, dt, menu, ol, ul { display: block; }
//     li { display: list-item; text-align: match-parent; }
var AllHtmlBlockTags map[string]void
var allHtmlBlockTags = []string{
	"address", "article", "aside",
	"body", "blockquote", "center",
	"dd", "dialog", "dir", "div", "dl", "dt",
	"fieldset", "figcaption", "figure", "footer", "form",
	"header", "hgroup", "hr", "html",
	"h1", "h2", "h3", "h4", "h5", "h6",
	"legend", "li", "listing",
	"main" /* "menu", */, "nav", "ol",
	"p", "plaintext", "pre",
	"search", "section",
	"table", "tfoot", "ul", "xmp",
}

// ============
//    INLINE
// ("phrasing")
// ============

// AllHtmlInlineTags are listed
// [here](https://html.spec.whatwg.org/#phrasing-content) and
// [here](https://html.spec.whatwg.org/#phrasing-content-3)
var AllHtmlInlineTags map[string]void
var allHtmlInlineTags = []string{
	"a", "abbr", "acronym", "map>area", "audio",
	"b", "bdi", "bdo", "big", "br", "button",
	"canvas", "cite", "code", "col", "command",
	"data" /* "datalist", */, "del", "dfn", "em", "embed",
	"i", "iframe", "img", "input", "ins",
	"kbd", "label", /* "link", */
	"map", "mark", "math", /* "menu", */
	"meta+itemprop", "meter",
	"nobr", "noscript", "object", "output",
	"picture", "progress", "q", "rt", "ruby",
	"s", "samp" /* "script", */, "select", "slot", "small",
	"span", "strike", "strong", "sub", "sup", "svg",
	/* "template", */ "textarea", "time", "tt",
	"u", "var", "video", "wbr",
}

// ======
// HIDDEN
// ======

// AllHtmlHiddenTags are listed
// [here](https://html.spec.whatwg.org/#hidden-elements):
// area, base, basefont, datalist, head, link, meta, noembed,
// noframes, param, rp, script, style, template, title {
// display: none; }
var AllHtmlHiddenTags map[string]void
var allHtmlHiddenTags = []string{
	"area", "base", "basefont", "datalist", "head",
	"link", "menu", "meta", "noembed", "noframes", "param",
	"rp", "script", "style", "template", "title",
}

// AllHtmlVoidTags are start tags that do not take end tags:
// "endless", "slashless". All are
// [here](https://html.spec.whatwg.org/multipage/syntax.html#elements-2):
// area, base, br, col, embed, hr, img, input, link, meta, source, track, wbr
//
// (Quoting) Void elements only have a start tag;
// end tags must not be specified for void elements.
// .
var AllHtmlVoidTags map[string]void
var allHtmlVoidTags = []string{
        "area", "base", "br", "col", "command", "embed",
        "hr", "img", "input", "keygen", "link", "meta",
	"param", "source", "track", "wbr"}

func init() {
	AllHtmlVoidTags = make(map[string]void)
	for _, s := range allHtmlVoidTags {
		AllHtmlVoidTags[s] = aVoid
	}
}

// = = = = = = = = =
/*

4.4 Grouping content

https://html.spec.whatwg.org/#flow-content-2

3.2.5.2.2 Flow content

Most elements that are used in the body of documents and applications are categorized as flow content.

aabbraddressarea (if it is a descendant of a map element)articleasideaudiobbdibdoblockquotebrbuttoncanvascitecodedatadatalistdeldetailsdfndialogdivdlemembedfieldsetfigurefooterformh1h2h3h4h5h6headerhgrouphriiframeimginputinskbdlabellink (if it is allowed in the body)main (if it is a hierarchically correct main element)mapmarkMathML mathmenumeta (if the itemprop attribute is present)meternavnoscriptobjectoloutputppicturepreprogressqrubyssampscriptsearchsectionselectslotsmallspanstrongsubsupSVG svgtabletemplatetextareatimeuulvarvideowbrautonomous custom elementstext

4.5 Text-level semantics

3.2.5.2.2 Flow content [AS ABOVE]

https://html.spec.whatwg.org/#phrasing-content

3.2.5.2.5 Phrasing content

Phrasing content is the text of the document, as well as elements that mark up
that text at the intra-paragraph level. Runs of phrasing content form paragraphs.

a abbr map>area audio b bdi bdo br button canvas cite code data datalist del dfn
em embed i iframe img input ins kbd label link map mark math meta+itemprop meter
noscript object output picture progress q ruby s samp script select slot small
span strong sub sup svg template textarea time u var video wbr text

Most elements that are categorized as phrasing content can only contain elements
that are themselves categorized as phrasing content, not any flow content.

https://html.spec.whatwg.org/#embedded-content-2

3.2.5.2.6 Embedded content

Embedded content is content that imports another resource into the document,
or content from another vocabulary that is inserted into the document.
*/
var AllHtmlEmbeddedContentTags = []string{
	"audio", "canvas", "embed", "iframe", "img",
	"math", "object", "picture", "svg", "video",
}

/*
https://html.spec.whatwg.org/#interactive-content

Interactive content is content that is specifically intended for user interaction.
*/
var AllHtmlInteractiveContentTags = []string{
	"a+href", "audio+controls", "button", "details",
	"embed", "iframe", "input", "label", "select",
	"textarea", "video+controls",
}

// = = = = = = = = =
/*

https://html.spec.whatwg.org/multipage/syntax.html#syntax-tag-omission

Certain tags can be omitted. Omitting an element's start tag
in the situations described below does not mean the element
is not present; it is implied, but it is still there. For
example, an HTML document always has a root html element,
even if the string <html> doesn't appear anywhere in the markup.

*/
