package lwdx

import "time"

// LinkInfo summarizes an outbound or incoming link.
type LinkInfo struct {
	text string
	// More TBD
}

// LwDocCtx is contextual info (DB ID, import batch info, linking info)
// common to LwDITA maps, LwDITA topics, and other compatible documents.
type LwDocCtx struct {
	ID                 int
	OutboundLwDocLinks []LinkInfo
	OutboundHyperLinks []LinkInfo
	VisibleLinkTargets []string
	ImportBatchID      int
	OwnImportPath      string
}

// LwdMap is TODO.
type LwdMap struct {
	LwDocCtx
	Navtitle    string // topicmeta > navtitle: (#PCDATA|%ph;)* = all-inline
	Linktext    string // topicmeta > linktext: (#PCDATA|%ph;)* = all-inline
	HeaderDatas string // topicmeta > data: (K,V)[]
	Body        string // (topicref|keydef)[]:
	// > topicref: (topicmeta?, topicref*)
	// > keydef:   (topicmeta?, data*)
	// Props[] K,V
}

//LwdMapTableCreate creates the DB table.
const LwdMapTableCreate = "CREATE TABLE IF NOT EXISTS lwd_map" +
	"(id INTEGER PRIMARY KEY, tmdnavtitle TEXT, tmdlinktext TEXT, " +
	"tmddatas TEXT, contents TEXT, importbatchid INT, ownimportpath TEXT, " +
	"outwardlinks TEXT, inwardlinks TEXT);"

// LwdTopic is TODO.
type LwdTopic struct {
	LwDocCtx
	Title       string // common-inline
	Shortdesc   string //    all-inline
	HeaderDatas string // prolog > data: (K,V)[]
	Body        string // ((%list-blocks;)*, section*, fn*
}

//LwdTopicTableCreate creates the DB table.
const LwdTopicTableCreate = "CREATE TABLE IF NOT EXISTS lwd_topic" +
	"(id INTEGER PRIMARY KEY, title TEXT, shortdesc TEXT, " +
	"prologdatas TEXT, body TEXT, importbatchid INT, ownimportpath TEXT, " +
	"outwardlinks TEXT, inwardlinks TEXT);"

// ImportBatch is TODO.
type ImportBatch struct {
	ID       int
	Name     string
	Time     time.Time
	Path     string
	NrMaps   int
	NrTopics int
	HasDirs  bool
}

// ImportBatchTableCreate creates the DB table.
const ImportBatchTableCreate = "CREATE TABLE IF NOT EXISTS import_batch" +
	"(id INTEGER PRIMARY KEY, name TEXT, time TIMESTAMP, path TEXT, " +
	"filect INTEGER, hasmaps INTEGER, hastopics INTEGER, hasdirs INTEGER);"

	/*
	   About time:
	   https://gobyexample.com/epoch

	    Use time.Now with Unix() to get elapsed
	   time since the Unix epoch in seconds.

	       now := time.Now()
	       fmt.Println(now)
	       // 2012-10-31 16:13:58.292387 +0000 UTC
	       secs := now.Unix()
	       fmt.Println(secs)
	       // 1351700038
	       fmt.Println(time.Unix(secs, 0))
	       // 2012-10-31 16:13:58 +0000 UTC
	*/
