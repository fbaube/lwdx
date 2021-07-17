package lwdx

import (
	"fmt"
	"os"
	"github.com/fbaube/gtoken"
	// "github.com/fbaube/gtree"
)

// SetLwdTagType sets IsBlock, IsInline, IsUnseen, or none.
func SetLwdTagType(N *gtoken.GToken) bool {
// func SetLwdTagType(N *gtree.GTag) bool {
	var i, ilim int
	N.IsBlock = false
	N.IsInline = false
	ilim = len(BlockTags)
	for i = 0; i < ilim; i++ {
		if N.GName.Local == BlockTags[i] {
			N.IsBlock = true
			return true
		}
	}
	ilim = len(InlineTags)
	for i = 0; i < ilim; i++ {
		if N.GName.Local == InlineTags[i] {
			N.IsInline = true
			return true
		}
	}
	fmt.Fprintf(os.Stderr, "Tag <%s>: block/inline not found \n",
		N.GName.Local)
	return false
}
