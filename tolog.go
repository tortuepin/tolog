// vim:foldmethod=marker:
package main

import "fmt"
import "flag"
import "log"
import "os"

//import "bufio"
import "time"
import "encoding/json"

import "./libs"

func main() {
	flag.Parse()
	args := flag.Args()

	if flag.NArg() != 1 {
		log.Fatal("Usage : com filename")
	}

	log_file := args[0]
	f, err := os.Open(log_file)
	if err != nil {
		// エラー時の処理
		log.Fatal(err)
	}
	defer f.Close()

	//scanner := bufio.NewScanner(f)

	//logs := []tolog.LogItem{}
	//todos := []tolog.TodoItem{}
	//for scanner.Scan() {
	//	if scanner.Text() == tolog.TodoHeader {
	//		todos = tolog.TodoReader(scanner)
	//	}
	//	if scanner.Text() == tolog.LogHeader {
	//		logs = tolog.LogReader(scanner)
	//	}
	//}
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(tolog.TodoGetActive(todos))
	//fmt.Println(tolog.GetFilenames("test/data/", time.Now(), 20))
	//fmt.Println(todos)
	//fmt.Println(logs)
	//a := tolog.GetAllItems("test/data/", time.Now(), 20)
	a := tolog.GetAllItems("./test/data/", time.Date(2019, 2, 27, 0, 0, 0, 0, time.Local), 20)
	//for _, i := range a {
	//	fmt.Println(i.Filename, "filename")
	//	fmt.Println(i.Todo)
	//	fmt.Println(i.Log)
	//}
	jsonByte, err := json.Marshal(a)
	fmt.Println(string(jsonByte))
}
