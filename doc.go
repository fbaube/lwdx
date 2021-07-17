// Package lwdx does Lightweight Document Transformations.
// Documents are presumed to be Lightweight DITA ("LwDITA") but in
// practice, deviations and alternatives will probably be handled.
//
// This package tries to kill several birds with one quite abstract stone:
// 
//  - Provide round-trippable transformations among all the "official"
//    representations of LwDITA (XDITA,HDITA,MDITA) (and maybe throw in
//    JSON too, for compatibility with pandoc)
//  - Provide a nice one-size-fits-all representation of "documents"
//  - Correctly handle XML mixed content, which means not relying too
//    much on Golang's encoder/xml, which is highly record-oriented; 
//    note that Golang's backquote-annotations for mapping XML to 
//    structure fields could be used instead but it would require a
//    lot of nested "innerxml" attributes, and who the heck needs that ?
//  - Also, Golang's encoder/xml package does not have representations
//    for parse trees, ASTs, or AST nodes, so ...
//  - [Planned] Provide maximum interoperability by implementing (partially,
//    at least) key interfaces of "popular" Golang markup-related packages
//  - Finally, make it unnecessary for the user to understand XML in any
//    "deep", psychosis-inducing, crazy-complex way.
//
// NOTE:100 DEFINITIONS:
//  - A "tag" is everything btwn "<" and ">", and does include the attributes.
//  - An "elm" is the start tag, end tag, and everything in-btwn.
//  - "element content" is everything btwn the matching start and end tags.
//  - A "doc" is a root elm (and its subtree) plus metadata about the file.
//  - Our own hierarchy is Doc => Elm => Tag => NS+Name+[]Atr
//
// The goal of LwDITA is not to re-create all the unwelcome complexity of
// DITA, but rather, to cut thru it all and reduce the cognitive burden of
// learning and using DITA.
//
// Similarly, the goal of XML processing in this package is not to re-create
// all the unwelcome complexity of XML. The processing here is more liberal,
// less constrained; those requiring strict compliance and validation should
// use other tools *somewhere* else in the toolchain.
//
// We do seek maximum interop with all other Go code, tho, so it is critical
// (and lazy) that we rely on Go's XML package (and other Go packages) to do
// the heavy lifting of encoding identification and XML tokenization. Beyond
// that, we rely on our own wits, and we DO try to parse and make sense of XML
// declarations (the file-initial "<?xml ..."), and we do a cursory check of
// DOCTYPE declarations.
//
// When processing XML we have the basic choice of building a tree or checking
// parsing events. This choice is normally described as "DOM-oriented" versus
// "stream parsing". Here in package lwdx we choose to build a tree, for
// simplicity and completeness. There are valid reasons tho to avoid building
// a tree; as documentation for the SAX parser says,
//
//  Tree-based APIs are useful for a wide range of applications, 
//  but they normally put a great strain on system resources, 
//  especially if the document is large ... It is inefficient 
//  to build a tree of parse nodes, only to map it onto a new
//  data structure and then discard the original.
//
// Well, one of the reasons we all dig the Go language is its efficiency, 
// so let's not worry too much about littering system memory with remnants 
// of parsing, at least not until we have some good becnhmarks for this app.
// Also, we're dealing with mixed content that is being used for purposes
// of producing documentation; we most decidedly are NOT dealing with huge
// record-orient XML files that act as surrogates for database tables, so
// let's not worry too awfully much about straining system resources. If
// you find yourself writing 20MB files in DITA or DocBook, you should
// probably have your head examined.
//
// Like the Go language itself, this package is opinionated about formatting.
// An important benefit of this is that file diff's, generated when checking
// in file changes, will be more reliable and more consistently formatted.
// Therefore, given the slightest chance, XML will be mercilessly reformatted,
// altho not necessarily in-place so as to avoid harming those with tender
// code format sensibilities.
//
// The issue of performing text replacements (using XML general entities)
// is open. If we can do it using someone else's code, that's fine, but if
// not, we declare their unavailability to be a limitation of this package,
// because anyways text substitution is sposta be done "the DITA way" with
// conref's, NOT by relying on gnarly XML functionalities.
//
// Note that when we generate a parse tree, we try to simplify data structures 
// and code, at the cost of complicating the *interpretation* of the data
// structures. Our parse trees contain nodes, and we use the basic node
// structure to represent non-tag markup, tags, elements, and text content.
// For non-tag markup (all the "<?" and "<!" stuff), we set the node field
// "StringValue", and this is a signal to check and interpret the node's NS
// (namespace) and Name (i.e. XML tag) fields as something other than markup
// and mixed content. The system looks like this:
//
// For tags, NS holds any XML namespaces (with a trailing colon !), but NS can
// also be "<?", "<!" (without that trailing colon !) and then we also use
// StringValue. This convention is also used for for CharData !
//
//   <? = PI, sort-of including opening "<xml" (mnemonic: "?" = "IS it XML_?_")
//   <! = Comments, CDATA, DOCTYPE
//      (& in DTDs [only]: ELEMENT, ATTLIST, ENTITY, NOTATION, Cond'l section)
//
// In each case, the node's field StringValue contains the text of the non-tag
// markup. It is stripped of the text that is in the field Name. For example:
//
//  PI:      NS "<?" Name The-PI's-target
//  DOCTYPE: NS "<!" Name "DOCTYPE"
//  Comment: NS "<!" Name "--"
//  PCDATA:  NS "  " Name "" (i.e. the content text in mixed content)
//  CDATA:   NS "<!" Name "CDATA" (i.e. <![[ CDATA ...)
//          (but maybe the parser handles these, so if we want 
//           to write them back out, we need to use a flag ?
//           So maybe we could just use an inline tag ? 
//
// "One of the ideas behind LwDITA is to trying
// to reduce the cognitive load for new users."
//
// golang encoding/xml:
// "Mapping between XML elements and data structures is inherently flawed: 
// an XML element is an order-dependent collection of anonymous values, 
// while a data structure is an order-independent collection of named 
// values. See package json for a textual representation more suitable 
// to data structures. "
package lwdx
