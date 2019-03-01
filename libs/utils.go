package tolog

import "time"
import "os"
import "sort"
import "log"
import "bufio"

type TologItem struct {
	Todo     []TodoItem `json:"todo"`
	Log      []LogItem  `json:"log"`
	Filename string     `json:"filename"`
}

/*
	GetFilenamesはdateに指定された日からn日前までの存在するファイル名を返す関数

	dirは"path/"みたいに、最後に/つけること
*/
func GetFilenames(dir string, date time.Time, n int) []string {
	files := []string{}
	d := date
	for i := 0; i < n; i++ {
		d = date.AddDate(0, 0, -i)
		file := dir + d.Format(DateFormat) + FileType
		if Exists(file) {
			files = append(files, file)
		}
		sort.SliceStable(files, func(i, j int) bool { return files[i] < files[j] })
	}
	return files
}

/*
	Existsはファイルの存在をチェックするやつ
*/
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

/*
	GetAllItemsはtodo, log, ...,を全部とってくるやつ
*/
func GetAllItems(dir string, date time.Time, n int) []TologItem {
	files := GetFilenames(dir, date, n)
	tolog_items := []TologItem{}
	for _, file := range files {
		item := TologItem{}
		item.Filename = file

		logs := []LogItem{}
		todos := []TodoItem{}

		f, err := os.Open(file)
		if err != nil {
			// エラー時の処理
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			if scanner.Text() == TodoHeader {
				todos = TodoReader(scanner)
			}
			if scanner.Text() == LogHeader {
				logs = LogReader(scanner)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		item.Log = logs
		item.Todo = todos

		tolog_items = append(tolog_items, item)
	}
	return tolog_items
}
