package lwdx

// Equivalents is an imperfect attampt to correlate tags across
// LwDITA languages.
//
// Note that for MDITA, the "Mdita" value is from the Goldmark parser
// library, and therefore (to take one example), "Document" means that
// for a node in question, its property [goldmark.ast.NodeKind] has the
// value [goldmark.ast.KindDocument].
// .
var Equivalents = []Equivs{
	{"topic", "html", "Document"},
}
