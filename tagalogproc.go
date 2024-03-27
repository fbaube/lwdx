package lwdx

import (
	"fmt"
	SU "github.com/fbaube/stringutils"
	L "github.com/fbaube/mlog"
)

var XditaToTE map[string]*TagalogEntry = make(map[string]*TagalogEntry)

var HditaToTE map[string]*TagalogEntry = make(map[string]*TagalogEntry)

var CTTypeToTE map[CTType]*TagalogEntry = make(map[CTType]*TagalogEntry)

func init() {
	var pTE *TagalogEntry
	var s string
	var ct CTType
	var ok bool

	for _, pTE = range Tagalog {
		// XDITA
		s = pTE.Xdita
		_, ok = XditaToTE[s]
		if ok {
			if s != "" && s != "?" {
				L.L.Warning("XDita dupe: " + s)
			}
		} else {
			XditaToTE[s] = pTE
		}
		// HDITA
		s = pTE.Hdita
		_, ok = HditaToTE[s]
		if ok {
			if s != "" && s != "?" {
				L.L.Warning("HDita dupe: " + s + " (no prob)")
			}
		} else {
			HditaToTE[s] = pTE
		}
		// CTType
		ct = pTE.CTType
		_, ok = CTTypeToTE[ct]
		if ok {
			if ct != "" {
				L.L.Warning("CTType dupe: " + string(ct))
			}
		} else {
			CTTypeToTE[ct] = pTE
		}
	}
	// fmt.Printf("XDITA: %+v \n", XditaToTE)
	// fmt.Printf("HDITA: %+v \n", HditaToTE)

	L.L.Info("BLOCK:")
	for _, pTE = range Tagalog {
		if pTE.IsBlock {
			L.L.Info("%s \t %s \t %s",
				pTE.CTType, pTE.Xdita, pTE.Hdita)
		}
	}
	L.L.Info("INLINE:")
	for _, pTE = range Tagalog {
		if pTE.IsInline {
			L.L.Info("%s \t %s \t %s",
				pTE.CTType, pTE.Xdita, pTE.Hdita)
		}
	}
}

// Use zero value
// var failTE TagalogEntry

func GetTEbyXdita(s string) *TagalogEntry {
	pTE, ok := XditaToTE[s]
	if ok {
		return pTE
	}
	L.L.Warning("Get TagalogEntry by XDITA: failure for: " + s)
	return nil // failTE
}

func GetTEbyHdita(s string) *TagalogEntry {
	pTE, ok := HditaToTE[s]
	if ok {
		return pTE
	}
	// HACK
	if s == "html" || s == "head" || s == "h1" || s == "h2" {
	   return nil
	   }
	L.L.Warning("Get TagalogEntry by HDITA: failure for: " + s)
	return nil // failTE
}

func GetTEbyCTType(s CTType) *TagalogEntry {
	pTE, ok := CTTypeToTE[s]
	if ok {
		return pTE
	}
	L.L.Warning("Get TagalogEntry by CTType: failure for: " + string(s))
	return nil // failTE
}

func GetTEbyTagAndMarkupType(tag string, mut SU.MarkupType) *TagalogEntry {
	switch mut {
	case SU.MU_type_XML:
		return GetTEbyXdita(tag)
	case SU.MU_type_HTML:
		return GetTEbyHdita(tag)
	case SU.MU_type_MKDN:
		fmt.Printf("tagalogproc:lacksMKDN:L113")
	case SU.MU_type_BIN:
		fmt.Printf("tagalogproc:lacksBIN:?!?!:L115")
		return nil
	default:
		fmt.Printf("tagalogproc:nil:?!?!:L118")
		return nil
	}
	return nil
}
