package tolog

import "time"
import "os"

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
