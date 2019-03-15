package main

import "flag"
import "os"
import "time"
import "log"

//import "bufio"
import "./libs"

const DIR_ENV = "TOLOG_DIR_ENV"

var (
	aDir  = flag.String("d", "", "dataディレクトリの場所")
	aFile = flag.String("f", "", "ファイル名")
	aN    = flag.Int("n", 30, "何日分のtodoを集めるか(デフォルト30)")
)

func main() {
	flag.Parse()
	// aDir, aFileが""だった場合環境変数/今日の日付から読み込み
	// TODO これ環境変数から読み込まないほうがいいかもしれない
	// 環境変数が設定されていなかったときdに""が入るよ
	// 必ずoptionから受け取るようにするほうが安全
	d := *aDir
	f := *aFile
	n := *aN
	if d == "" {
		d = os.Getenv(DIR_ENV)
	}
	if f == "" {
		f = time.Now().Format(tolog.DateFormat) + tolog.FileType
	}

	targetDir := d
	targetFile := d + "/" + f
	// ファイル名からその日付のtimeオブジェクトをつくる
	targetDate, err := time.Parse(tolog.DateFormat+tolog.FileType, f)
	if err != nil {
		log.Fatal("can not parse '", f, "' to time")
	}

	// 引数の検証
	// ディレクトリ, ファイルの存在確認
	if !tolog.Exists(targetFile) {
		log.Fatal(targetFile, " is not Exist")
	}

	// ファイルを読み込む
	lines := tolog.ReadLines(targetFile)
	// Todoの場所をみつける(配列は0始まりなので-1)
	tStart, tEnd := tolog.HeaderSearcher(targetFile, tolog.TodoHeader)
	tStart = tStart - 1
	tEnd = tEnd - 1

	// 新しいtodoリストを作る
	allItems := tolog.GetAllItems(targetDir, targetDate, n)
	todos := []tolog.TodoItem{}
	for _, i := range allItems {
		todos = append(todos, i.Todo...)
	}
	newTodos := tolog.TodoGetActive(todos)
	key, todoMap := tolog.TodoGetTagMap(newTodos)
	newTodoLines := []string{}
	newTodoLines = append(newTodoLines, tolog.TodoHeader)
	newTodoLines = append(newTodoLines, tolog.TodoMap2Strings(todoMap, key)...)

	// 差し替える
	ret := tolog.SliceReplacer(lines, newTodoLines, tStart, tEnd)

	// ファイルを書き換える
	fp, err := os.Create(targetFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	for _, r := range ret {
		fp.WriteString(r + "\n")
	}

	log.Println("end")
}
