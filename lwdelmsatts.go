package lwdx

// THIS FILE is based mostly on the series of PDF dpcuments
//

// LwTag is a generic "tag" for a structure in a LwDoc. -n-
// It has a "master" ComponentName and variants for (XHM)DITA.
type LwTag struct {
	ComponentName string
	Xdita         string
	Hdita         string
	Mdita         string
}

// BlockTags is all the tags that either (a) are considered to be block tags
// by HTML & CSS, or (b) should in any case be output starting on a new line.
// The values here should be based on references, and not just pulled out of
// a butt.
var BlockTags = []string{
	"topic",
	"title",
	"shortdesc",
	"body",
	"section",
	"ul",
	"ol",
	"li",
	"taskbody",
	"context",
	"prereq",
	"steps",
	"step",
	"table",
	"hr",
	"h1",
	"h2",
	"h3",
	"h4",
	"h5",
	"author",
	"keydef",
	"keyword",
	"keywords",
	"map",
	"topicmeta",
	"topicref",
	"navtitle",
}

// InlineTags is all the tags that should be output on
// "the same line". The values here should be based on
// references, and not just pulled out of a butt.
var InlineTags = []string{
	"b",
	"i",
	"u",
	"em",
	"strong",
	"ph",
	"span",
}

// OldTags is Appendix A.1 DITA 1.3 elements in LwDITA -n-
// "This topic lists the DITA 1.3 elements that are available in LwDITA.
// It also lists how to represent them in XDITA, HDITA, and MDITA."
var OldTags = []LwTag{
	{"Component", "XDITA", "HDITA", "MDITA"},
	{"Alt text", "<alt>", "Attr on <img>", "[text]"},
	{"Body", "<body>", "<body>", "N/A"},
	{"Bold", "<b>", "<strong>", "** or __"},
	{"Cross ref", "<xref>", "<href>", "[link](/URI 'title')"},
	{"Data", "<data>", "<meta>", "YAML.hdr vars"},
	{"Def desc", "<dd>", "<dd>", "XDITA.<dd>"},
	{"Def list entry", "<dlentry>", "N/A", "N/A"},
	{"Def term", "<dt>", "<dt>", "XDITA.<dt>"},
	{"Def list", "<dl>", "<dl>", "XDITA.<dl>"},
	{"Description", "<desc>", "<caption>",
		"in <table>; <figcaption> in <figure>; N/A in links"},
	{"Figure", "<fig>", "<figure>", "N/A"},
	{"Footnote", "<fn>", "@data-hd-class=\"fn\"", "XDITA.<fn>"},
	{"Image", "<image>", "<img>",
		"![alt text for an image](images/image_name.jpg)"},
	{"Italics", "<i>", "<em>", "* or _"},
	{"List item", "<li>", "<li>", "' -,+,* for ul; 0-9,.,) for ol"},
	{"Map", "<map>", "<nav>", "N/A"},
	{"Note", "<note>", "<p data-hd-class=\"note\">", "XDITA.<note>"},
	{"Ordered list", "<ol>", "<ol>", "See list item"},
	{"Paragraph", "<p>", "<p>", "Two CR's"},
	{"Nav title", "<navtitle>", "N/A", "N/A"},
	{"Phrase", "<ph>", "<span>", "XDITA.<ph>"},
	{"Pre text", "<pre>", "<pre>", "```text```"},
	{"Prolog", "<prolog>", "<meta>", "YAML.hdr"},
	{"Section", "<section>", "<section>", "##"},
	{"Short description", "<shortdesc>", "<p data-hd-class=\"shortdesc\">",
		"YAML.hdr.var shortdesc (yes|no) says if first para is a short desc"},
	{"Table", "<simpletable>", "<table>", "GH-MD syntax, using | and -"},
	{"S.tbl entry", "<stentry>", "<td>", "See Table"},
	{"S.tbl header", "<sthead>", "<th>", "See Table"},
	{"S.tbl row", "<strow>", "<tr>", "See Table"},
	{"Subscript", "<sub>", "<sub>", "XDITA.<sub>"},
	{"Superscript", "<sup>", "<sup>", "XDITA.<sup>"},
	{"Title", "<title>", "<h1> topic, <h2> section", "# topic, ## section"},
	{"Topic", "<topic>", "<article>", "N/A"},
	{"Topic metadata", "<topicmeta>", "N/A", "N/A"},
	{"Topicreference", "<topicref>", "<href> inside a <li>",
		"[link](/URI 'title') inside a list item"},
	{"Underline", "<u>", "N/A", "N/A"},
	{"Unordered list", "<ul>", "<ul>", "See list item"},
}

// NewTags is Appendix A.2 New elements -n-
// "This topic lists the new XML elements that are part of LwDITA and how
// to represent them in XDITA and HDITA. These new elements are not avail-
// able in the MDITA core profile and, if needed, can be represented with
// their raw HDITA equivalents as part of the MDITA extended profile."
var NewTags = []LwTag{
	{"Component", "XDITA", "HDITA", "MDITA"},
	{"Audio", "<audio>", "<audio>", "XDITA.<audio>"},
	{"Controls", "<controls>", "<audio|video>.@controls", "XDITA.@controls"},
	{"Fallback", "<fallback>", "<source>", "XDITA.<fallback>"},
	{"Footnote reference", "<fnref>", "?", "?"},
	{"Poster", "<poster>", "<video>.@poster", "XDITA.<video>.@poster"},
	{"Source", "<source>", "<source>", "XDITA.<source>"},
	{"Track", "<track>", "<audio|video>.@track", "XDITA.<audio|video>.@track"},
	{"Video", "<video>", "<video>", "XDITA.<video>"},
}

// OldAtts is Appendix A.3 DITA 1.3 attributes in LwDITA -n-
// "This topic lists the DITA 1.3 attributes available in LwDITA,
// and how to represent them in XDITA and HDITA."
var OldAtts = []LwTag{
	{"Component", "XDITA", "HDITA", "MDITA"},
	{"Content reference", "@conref", "@data-hd-conref", "XDITA.@conref"},
	{"Direction", "@dir", "@dir", "XDITA.@dir"},
	{"Expanse", "@expanse", "N/A", "N/A"},
	{"Frame", "@frame", "N/A", "N/A"},
	{"Identifier", "@id", "@id", "topic: YAML.hdr.id; else XDITA.@id"},
	{"Importance", "@importance", "@data-hd-importance", "XDITA.@importance"},
	{"Key reference", "@keyref", "@data-hd-keyref", "XDITA.@keyref"},
	{"Language", "@xml:lang", "@lang",
		"topic: YAML.hdr.lang; else XDITA.@xml:lang"},
	{"Output class", "@outputclass", "@class", "XDITA.@outputclass"},
	{"Props", "@props", "@data-hd-props", "XDITA.@props"},
	{"Scale", "@scale", "N/A", "N/A"},
	{"Translate", "@translate", "@translate", "XDITA.@translate"},
	{"Type", "@type", "@data-hd-type", "XDITA.@type"},
}
