// NOTION (and also OTHER) block types

// NOTES on syntaxes

// DTD excerpts

// https://html.spec.whatwg.org/#non-replaced-elements and forward from there.
// Block  elements always start from a new line.
// Inline elements  never start from a new line.
// Block  elements cover space L-to-R as far as it can go.
// Inline elements cover space only as bounded by the tags in the HTML element.

package lwdx

func BIO(B, I bool) BlockInlineOther {
	return BlockInlineOther{B, I}
}

// REFERENCE: latest official as of 2023.03:
// https://www.oasis-open.org/committees/download.php/70571/ (PDF)
// https://www.oasis-open.org/committees/download.php/70836/lwDITA_SC_minutes_06_March_2023.txt
// Lightweight DITA Version 1.0
// Working Draft 06
// 28 November 2022

// BUT WD8? 18 January 2023
// ?? https://dita-lang.org/lwdita/resources/oasis-cover.html

// re B I EM STRONG :
// https://www.ditawriter.com/a-more-semantic-way-of-describing-emphasized-inline-content-in-dita-2-0/

// CT_TBS indicates a CTT that is not defined by Pandoc
// but might end up being defined here as an addition.
var CT_TBS CTType = ""

// Tagalog is the tag catalog, including info for all three
// flavors of LwDITA, and a 3-state flag Block/Inline/Neither.
//
// Note that for MDITA,
//   - the entry "" indicates no such element, certainly not an element
//     found in Goldmark.
//   - the entry "?" indicates a value that maybe can be filled in after
//     Goldmark is properly analysed, and if not, then element requires
//     syntax that is not expressible -- and might be replaced by "".
//   - the entry "(hdita)" indicates that the spec says to use HDITA
//     in the MDITA
//
// in the format of this table.
// .
var Tagalog TagalogListing = []*TagalogEntry{
	/* REF
	type TagalogEntry struct {
	     CTType
	     Xdita, Hdita, Mdita string
	     IsBlock, IsInline bool } */

	// T_Blk_Para, "", "", "", BIO(false, false)}, // 3.3.15
	{CT_Blk_CdBlk, "", "", "", BIO(false, false)},
	{CT_Blk_Quote, "", "", "", BIO(false, false)},
	// _Blk_OList, "", "", "", BIO(false, false)}, // 3.3.14
	// _Blk_UList, "", "", "", BIO(false, false)}, // 3.3.19
	{CT_Blk_Hedr, "", "", "", BIO(false, false)},
	{CT_Blk_HRule, "", "", "", BIO(false, false)},
	{CT_Blk_Table, "", "", "", BIO(false, false)},
	// _Blk_Figure, "", "", "", BIO(false, false)}, // 3.3.9
	// CTT_Blk_Div, "", "", "", BIO(false, false)}, // 3.3.4

	{CT_Inl_Text, "", "", "", BIO(false, false)},
	// T_Inl_Emph, "", "", "", BIO(false, false)}, // 3.5.1
	// _Inl_Undln, "", "", "", BIO(false, false)}, // 3.4.5
	// _Inl_Strng, "", "", "", BIO(false, false)}, // 3.5.2
	{CT_Inl_Strike, "", "", "", BIO(false, false)},
	// _Inl_Super, "", "", "", BIO(false, false)}, // 3.4.4
	// CT_Inl_Sub, "", "", "", BIO(false, false)}, // 3.4.3
	{CT_Inl_Quote, "", "", "", BIO(false, false)},
	{CT_Inl_Citatn, "", "", "", BIO(false, false)},
	{CT_Inl_Code, "", "", "", BIO(false, false)},
	{CT_Inl_LnBrk, "", "", "", BIO(false, false)},
	// _Inl_Link, "", "", "", BIO(false, false)}, // 3.3.20
	// Inl_Image, "", "", "", BIO(false, false)}, // 3.3.11
	// _Inl_Span, "", "", "", BIO(false, false)}, // 3.3.16

	// <linktext> is missing from the spec WD6,
	// so make an entry that is like <alt>
	//  CT_TBS, "alt", "alt", "", BIO(false, false)},
	{CT_TBS, "linktext", "?", "", BIO(false, false)},

	// ==========================
	// 3.2 Basic topic components
	// ==========================
	// <!ELEMENT topic (title, shortdesc?, prolog?, body?)  >
	// ==========================
	// 3.2.4 Topic
	{CT_TBS, "topic", "article", "", BIO(false, false)},
	// 3.2.3 Title
	// can be w topics, maps, sections, examples, figures, tables, etc.
	// HDITA, MDITA(core+ext.): Messy!
	{CT_TBS, "title", "title", "", BIO(true, false)},
	// 3.2.2 Short description
	// HDITA: First elm in article after title, if is para
	// MDITA(core+ext.): 1st block after title, if is para
	// "Processors SHOULD render the content of the <shortdesc>
	//      element as the initial paragraph of the topic."
	{CT_TBS, "shortdesc", "", "", BIO(true, false)},
	// 3.2.1 Body
	{CT_TBS, "body", "body", "", BIO(false, false)},
	// ===================
	// 3.3 Body components
	// ===================
	// 3.3.1 Alternate text
	// XDITA: <alt> inside <image>
	// HDITA: @alt attribute on <img>
	// MDITA(core+ext.): Text in brackets ([]) in an Image
	{CT_TBS, "alt", "alt", "", BIO(false, false)},
	// 3.3.2 Definition description (a definition's RH side)
	{CT_TBS, "dd", "dd", "", BIO(false, true)},
	// 3.3.3 Description
	// XDITA: <desc> inside <audio>, <f1;95;0cig>, <video>
	// HDITA: @title in <audio>, <video>
	// "When used in conjunction with figures, processors SHOULD
	// consider the content of description components to be part
	// of the content flow.
	// When used in conjunction with cross references, processors
	// often choose to render the content of description components
	// as hover help or other forms of link preview."
	{CT_TBS, "desc", "@title", "?", BIO(false, true)},
	// 3.3.4 Division
	{CT_Blk_Div, "", "", "", BIO(false, false)},
	// 3.3.5 Definition list (the outermost wrapper for the definition list)
	{CT_TBS, "dl", "dl", "?", BIO(true, false)},
	// 3.3.6 Definition list entry (wraps <dt> & <dd>, in XDITA only)
	{CT_TBS, "dlentry", "?", "?", BIO(true, false)},
	// 3.3.7 Definition term (a definition's LH side)
	{CT_TBS, "dt", "dt", "", BIO(false, true)},
	// 3.3.8 Example
	{CT_TBS, "example", "<div data-class=\"example\">",
		"h{.example}", BIO(true, false)},
	// 3.3.9 Figure
	{CT_Blk_Figure, "fig", "figure", "", BIO(true, false)},
	// 3.3.10 Footnote
	{CT_TBS, "fn", "<span data-class=\"fn\">", "?", BIO(true, false)},
	// 3.3.11 Image
	{CT_Inl_Image, "image", "img", "?", BIO(false, true)},
	// 3.3.12 List item
	{CT_TBS, "li", "li", "?", BIO(true, false)},
	// 3.3.13 Note
	{CT_TBS, "note", "<div data-class=\"note\">", "", BIO(true, false)},
	// 3.3.14 Ordered list
	{CT_Blk_OList, "ol", "ol", "?", BIO(true, false)},
	// 3.3.15
	{CT_Blk_Para, "p", "p", "?", BIO(true, false)},
	// 3.3.16 Paragraph
	// "The phrase component often is used to enclose a phrase for reuse
	// or conditional processing.
	// The phrase component frequently is used as a specialization base,
	// to create phrase-level markup that can provide additional semantic
	// meaning or trigger specific processing or formatting. For example,
	// all highlighting domain elements are specializations of phrase."
	{CT_Inl_Span, "ph", "span", "", BIO(false, true)},
	// 3.3.17 Preformatted text
	{CT_TBS, "pre", "pre", "?", BIO(true, false)},
	// 3.3.18 Section
	// "Multiple sections within a single topic do not represent a hierarchy,
	// but rather peer divisions of that topic. Sections cannot be nested.
	// Sections can have titles."
	{CT_TBS, "section", "section", "?", BIO(true, false)},
	// 3.3.19 Unordered list
	{CT_Blk_UList, "ul", "ul", "", BIO(true, false)},
	// 3.3.20 Cross reference
	// A cross reference is an inline link. A cross reference can link to
	// * a different location within the current topic
	// * another topic
	// * a specific location in another topic
	// * an external resource such as a PDF or web page.
	// MDITA: core: the equivalent of the XDITA @keyref attribute is supported.
	// extended: attributes can be specified by using the HDITA representation.
	{CT_Inl_Link, "xref", "a@href", "", BIO(false, true)},
	// ===========================
	// 3.4 Highlighting components
	// ===========================
	// Use these to highlight text with styles (such as bold, and italic).
	// Never use these when a semantically specific element is available.
	// These elements are not intended for use by specializers, and are
	// intended solely for use by authors when no semantically appropriate
	// element is available and a formatting effect is required.
	// ===========================
	// 3.4.1 Bold text
	// Bold text is used to draw a reader's attention to a
	// phrase without otherwise adding meaning to the content.
	{CT_Inl_Bold, "b", "b", "?", BIO(false, true)},
	// 3.4.2 Italic text
	// Italic text is used to emphasize the key points in printed text,
	// or when quoting a speaker, to show which words the speaker stressed.
	{CT_Inl_Italic, "i", "i", "?", BIO(false, true)},
	// 3.4.3 Subscript
	{CT_Inl_Sub, "sub", "sub", "", BIO(false, true)},
	// 3.4.4 Superscript
	{CT_Inl_Super, "sup", "sup", "", BIO(false, true)},
	// 3.4.5 Underline
	{CT_Inl_Undln, "u", "u", "", BIO(false, true)},
	// =======================
	// 3.5 Emphasis components
	// =======================
	// The emphasis elements are used to emphasize
	// text that is important or serious.
	// =======================
	// 3.5.1 Emphasized text
	// This indicates special meaning or particular importance.
	{CT_Inl_Emph, "em", "em", "*,_", BIO(false, true)},
	// 3.5.2 Strong text
	// This of greater importance than the surrounding text.
	{CT_Inl_Strng, "strong", "strong", "**,__", BIO(false, true)},
	// ==================
	// 3.6 Map components
	// ==================
	// Map components include the core components
	// of LwDITA maps, such as <topicref> and
	// <reltable> (needs to be updated for LwDITA).
	// ====================
	// 3.6.1 Key definition
	// A key definition is a simple way to define a key without
	// making the definition itself a part of rendered content.
	// It is a convenience component. It is equivalent to a
	// topic reference component that defines a key while also
	// setting @processing-role to "resource-only". Attributes
	// defaulted on the key definition component ensure that
	// key definitions do not appear in the TOC, do not add
	// extra links, and are not rendered as topics.
	// XDITA+HDITA: The following considerations apply:
	// • @keys is required.
	// • @href might be omitted when the key def is used for variable text.
	// • @processing-role has a default value of "resource-only".
	{CT_TBS, "keydef", "<div data-class=\"keydef\">", "(hdita)", BIO(false, true)}, // OR BLK
	// 3.6.2 Key text
	// A variable or link text that is used when resolving key references.
	// It also specifies alternate text for images that are referenced by keys.
	{CT_TBS, "keytext", "<div data-class=\"keytext\">", "(hdita)", BIO(false, true)},
	// 3.6.3 Map
	// DITA map is the way to aggregate topic references and define a context for
	// those references. It contains references to topics, maps, and other resources.
	// In a map referenced by another map, the title might be discarded as topics
	// from the submap are aggregated into a larger publication.
	// MDITA(core?ext'd): A Markdown file with title and UL of titles for topics and
	// their associated file names.
	// Rendering: One might make use of the relationships defined in the map to create
	// a TOC, aggregate topics into a PDF, or create links among topics in the output.
	{CT_TBS, "map", "nav", "TBS", BIO(true, false)},
	// 3.6.4 Navigation title
	// A navigation title is an alternative title for a resource. Use where the topic
	// title is unsuitable for use in a table of contents or navigation pane.
	{CT_TBS, "navtitle", "", "", BIO(false, false)},
	// 3.6.5 Topic metadata
	// This is metadata that applies to a topic based on its context in a map.
	// The metadata in a topic metadata component is specific to a given context
	// within a map. If a reference to a single resource appears more than once
	// in a map or set of maps, unique metadata can be specified in each instance.
	// For example, when the parent topic reference results in a link, components
	// within the topic metadata component can be used to provide context-specific
	// information about the link, such as link text or a navigation title.
	{CT_TBS, "topicmeta", "", "", BIO(false, false)},
	// 3.6.6 Topic reference
	// This is a way to reference a topic (or another resource) from a DITA map.
	// It can nest, which enables the expression of navigation and ToC hierarchies,
	// as well as containment hierarchies and parent-child relationships.
	{CT_TBS, "topicref", "li>a@href", "li>[link](/URI \"title\")", BIO(false, true)},
	// =======================
	// 3.7 Metadata components
	// These include information in <topicmeta> (in maps) or <prolog> (in topics),
	// as well as indexing elements that can be placed in additional locations
	// within topic content. (Needs to be updated for LwDITA)
	// =======================
	// 3.7.1 Data
	// This is a generic component for metadata within a topic or map.
	// Complex metadata can use nested data structures.
	// By default, <data> contents are not rendered; do not
	// use it to embed text as part of the content flow. A processor MAY
	// use a particular data component to trigger specialized rendering.
	{CT_TBS, "data", "head>meta", "(yaml-header)", BIO(false, false)},
	// 3.7.2 Prolog
	// This has metadata about the topic; e.g. author information or subject category.
	{CT_TBS, "prolog", "head>meta", "yaml-header)", BIO(false, false)},
	// =========================
	// 3.8 Multimedia components
	// These reference audio or video content.
	// =========================
	// 3.8.1 Audio
	// This is modeled on the HTML5 <audio> element.
	// An audio resource can be referenced by @href, @keyref, and nested
	// media-source components. Playback behaviors such as auto-playing,
	// looping, and muting are determined by attributes. When not speci-
	// fied, default behavior is determined by the user agent that is
	// used to present the media.
	// When an audio resource cannot be rendered in a meaningful way,
	// processors SHOULD render the <fallback> element, if it is present.
	{CT_TBS, "audio", "audio", "(hdita)", BIO(false, false)},
	// 3.8.2 Fallback
	// This is content to present when multimedia objects cannot be rendered.
	{CT_TBS, "fallback", "?", "(hdita)", BIO(false, false)},
	// 3.8.3 Source
	// This specifies the location of an audio or video resource.
	// This is modeled on the HTML5 media <source>. If multiple <media-source>
	// are present, the user agent evaluates them in document order and selects
	// the first playable resource.
	{CT_TBS, "audio|video>media-source", "audio|video>source", "(hdita)", BIO(false, false)},
}
