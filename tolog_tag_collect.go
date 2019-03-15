package main

import "flag"
import "log"
import "path/filepath"
import "./libs"
import "os"

var (
	aDir  = flag.String("d", "", "dataディレクトリの場所")
	aFile = flag.String("o", tolog.TagFile, "tagリストを出力する場所")
)

func main() {
	flag.Parse()
	if *aDir == "" {
		log.Fatal("-d option should be set")
	}
	d, err := filepath.Abs(*aDir)
	if err != nil {
		log.Fatal(*aDir, " can not convert to Abs")
	}

	// ディレクトリの存在確認
	if !tolog.Exists(d) {
		log.Fatal(d, " is not Exist")
	}
	tagFile := d + "/" + *aFile

	files := tolog.GetAllFilenames(d)
	items := tolog.GetAllItemsFromFilenames(files)
	tmap := map[string]bool{}
	tags := []string{}
	for _, i := range items {
		// todoのとこのtagあつめ
		for _, t := range i.Todo {
			if !tmap[t.Tag] {
				tmap[t.Tag] = true
				tags = append(tags, t.Tag)
			}
		}
		// logのとこのtagあつめ
		for _, l := range i.Log {
			for _, ltag := range l.Tag {
				if !tmap[ltag] {
					tmap[ltag] = true
					tags = append(tags, ltag)
				}
			}
		}
	}

	// ファイルに書き込み
	fp, err := os.Create(tagFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	for _, t := range tags {
		fp.WriteString(t + "\n")
	}
}
