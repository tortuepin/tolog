package main

import "log"
import "flag"
import "time"
import "path/filepath"
import "fmt"

import "./libs"

var (
	aDir  = flag.String("d", "", "dataディレクトリの場所")
	aFile = flag.String("f", "", "開始ファイル名")
	aN    = flag.Int("n", 0, "何日前までみるか(指定なしだとぜんぶ)")
)

func main() {
	flag.Parse()
	tags := flag.Args()
	log.Println(tags)

	// 引数の検証
	if !tolog.Exists(*aDir) {
		log.Fatal("directory : '", *aDir, "' is not Exist")
	}
	// ファイル名からその日付のtimeオブジェクトをつくる
	targetDate, err := time.Parse(tolog.DateFormat+tolog.FileType, *aFile)
	if err != nil {
		log.Fatal(targetDate, " is not match filename format. [", tolog.DateFormat, "]")
	}

	// ファイル名取得
	filenames := []string{}
	if *aN <= 0 {
		filenames = tolog.GetAllFilenames(*aDir)
	} else {
		filenames = tolog.GetFilenames(*aDir, targetDate, *aN)
	}

	// 全部のItemを取得
	allItems := tolog.GetAllItemsFromFilenames(filenames)

	// 順番にlogを見ていってtagに当てはまるやつだけpuckup
	// こんな深くなるもん？？？
	// もっといい書き方できそうだけどねぇ
	// mapとか使ったらできるかも
	retLogs := map[string][]tolog.LogItem{}
	dates := []string{}
	for _, tologItems := range allItems {
		for _, log := range tologItems.Log {
			for _, tag := range log.Tag {
				for _, targettag := range tags {
					if tag == targettag {
						base := filepath.Base(tologItems.Filename)
						date, _ := time.Parse(tolog.DateFormat+tolog.FileType, base)
						_, ok := retLogs[date.Format(tolog.DateFormat)]
						if !ok {
							dates = append(dates, date.Format(tolog.DateFormat))
						}
						retLogs[date.Format(tolog.DateFormat)] = append(retLogs[tologItems.Filename], log)
					}
				}
			}
		}
	}

	// 出力する
	/*
		190226
		[12:20]
		...
		[12:30]
		...
		190227
		みたいな
	*/

	for _, d := range dates {
		fmt.Println(d)
		for _, l := range retLogs[d] {
			tolog.LogPrinter(l)
			fmt.Println()
		}
	}
}
