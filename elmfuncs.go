package lwdx

import (
	"fmt"
	"github.com/fbaube/gtoken"
	"os"
	// "github.com/fbaube/gtree"
)

// SetLwdTagType sets IsBlock, IsInline, IsUnseen, or none.
func SetLwdTagType(N *gtoken.GToken) bool {
	// func SetLwdTagType(N *gtree.GTag) bool {
	var i, ilim int
	N.IsBlock = false
	N.IsInline = false
	ilim = len(BLK_tags)
	for i = 0; i < ilim; i++ {
		if N.GName.Local == BLK_tags[i] {
			N.IsBlock = true
			return true
		}
	}
	ilim = len(INL_tags)
	for i = 0; i < ilim; i++ {
		if N.GName.Local == INL_tags[i] {
			N.IsInline = true
			return true
		}
	}
	fmt.Fprintf(os.Stderr, "Tag <%s>: block/inline not found \n",
		N.GName.Local)
	return false
}
