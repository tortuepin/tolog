package tolog

const FileType = ".md"
const TagFile = ".tags.tolog"
const DateFormat = "060102"
const TabSetting = 2
const HeaderPrefix = "## "
const TodoHeader = "## todo"
const LogHeader = "## log"

// 先頭に来れるタグ
var TagAtPrifix = []string{"@"}

// タグ全部
var TagPrifix = []string{"#", "@"}

type TologItem struct {
	Todo     []TodoItem `json:"todo"`
	Log      []LogItem  `json:"log"`
	Filename string     `json:"filename"`
}
type LogItem struct {
	Name     string   `json:"name"`
	Tag      []string `json:"tag"`
	Contents []string `json:"contents"`
}
type TodoItem struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	Done  bool   `json:"done"`
	Depth int    `json:"depth"`
}
