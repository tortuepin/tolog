package main

import "log"
import "flag"
import "time"

import "./libs"

var (
	aDir  = flag.String("d", "", "dataディレクトリの場所")
	aFile = flag.String("f", "", "開始ファイル名")
	aN    = flag.Int("n", 0, "何日前までみるか(指定なしだとぜんぶ)")
)

func main() {
	flag.Parse()

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
	for _, i := range allItems {
		log.Println(i)
	}

	// 出力する

}
