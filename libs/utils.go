package tolog

import "time"
import "os"
import "sort"
import "log"
import "bufio"
import "strings"

/*
	GetFilenamesはdateに指定された日からn日前までの存在するファイル名を返す関数

	dirは"path/"みたいに、最後に/つけること
*/
func GetFilenames(dir string, date time.Time, n int) []string {
	files := []string{}
	d := date
	for i := 0; i < n; i++ {
		d = date.AddDate(0, 0, -i)
		file := dir + "/" + d.Format(DateFormat) + FileType
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

// HeaderSearcherはheaderに指定されたsectionの範囲を探すやつ
func HeaderSearcher(filename string, header string) (int, int) {
	ret_start := -1
	ret_end := -1
	f, err := os.Open(filename)
	if err != nil {
		// エラー時の処理
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 1

	// 始まりを見つける
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \n")
		if line == header {
			ret_start = i
			i = i + 1
			break
		}
		i = i + 1
	}
	// 終わりを見つける
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \n")
		if ret_start != -1 && strings.HasPrefix(line, HeaderPrefix) {
			break
		}
		i = i + 1
		ret_end = i
	}

	return ret_start, ret_end
}

// ReadLinesはファイルをstringの配列として読み込むやつ
func ReadLines(filename string) []string {

	lines := []string{}
	f, err := os.Open(filename)
	if err != nil {
		// エラー時の処理
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// 始まりを見つける
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// SliceReplacerはmainの[]stringのstartからendまでをsubでいれかえるやつ
func SliceReplacer(main []string, sub []string, start int, end int) []string {
	ret := main[0:start]
	for _, l := range sub {
		ret = append(ret, l)
	}
	ret = append(ret, main[end:]...)
	return ret
}
