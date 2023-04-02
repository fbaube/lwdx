package lwdx

// TODO: Use Go EMBED directive for DTDs and LwDITA samples

// TagalogEntry is one record in a Tagalog - a Tag catalog
// .
type TagalogEntry struct {
	CTType
	Xdita, Hdita, Mdita string
	IsBlock, IsInline   bool
	// IsSelfClsg bool // self-closing, like <br/>
	// IsHidden   bool
}

func (p *TagalogEntry) String() string {
	return "(a TagalogEntry)"
}

// TagalogListing is a Tag Catalog for the singleton package variable [Tagalog].
// .
type TagalogListing []*TagalogEntry

// = = =

// TagTypeMapper is a type created for var gxml.TagTypes
//.
/// type TagTypeMapper map[string]TagSummary

// TagTypes is a singleton for quick characterization
// of all LwDITA tags and common HTML5 tags (not the
// exotic ones, whose handling will be ignored for now).
// .
// var TagInfo TagTypeMapper

// Tags with differing Modes:
// map: html = image map, lwdita = ToC
// body: lwdita = topic
// video:

// An inline element cannot contain a block-level element!

// Note that now <a ...> can wrap a lot.
// http://html5doctor.com/block-level-links-in-html-5/
// One new and exciting thing you can do in HTML 5 is
// wrap links round “block-level” elements.

// = = =

// Equivalents is an imperfect attampt to correlate tags across
// LwDITA languages.
// .
/* OBS
type Equivs struct {
	Xdita, Hdita, Mdita string
} */

// = = =

// LwD_XMH_gtag is a generic "tag" used to identify markup structures
// (both elements and attributes) that are common to all three flavors
// of LwDITA.
//
// It has a "master" ComponentName and the variants for (XHM)DITA.
// .
/* OBS
type LwD_XMH_gtag struct {
	ComponentName string
	Xdita         string
	Hdita         string
	Mdita         string
} */

// = = =

// type TagMode string

// = = =

// TagSummary is a set of booleans that quickly characterizes a tag, no
// matter what kind of "common" XML file it is found in - HTML5, LwDITA,
// maybe also DITA. There's a bit of mix & match going on between LwDITA
// and HTML5, and we want to be liberal about accepting near-misses (like
// B & I v EMPH & STRONG), so this approach kinda makes sense.
// .
/* OBS
type TagSummary struct {
	Html5Mode  TagMode
	LwditaMode TagMode
	// IsSelfClsg bool // self-closing, like <br/>
	// IsHidden   bool
} */
