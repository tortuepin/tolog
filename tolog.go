// vim:foldmethod=marker:
package main

import "fmt"
import "flag"
import "log"
import "os"
import "bufio"
import "time"

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

	scanner := bufio.NewScanner(f)

	logs := []tolog.LogItem{}
	todos := []tolog.TodoItem{}
	for scanner.Scan() {
		if scanner.Text() == tolog.TodoHeader {
			todos = tolog.TodoReader(scanner)
		}
		if scanner.Text() == tolog.LogHeader {
			logs = tolog.LogReader(scanner)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(tolog.GetFilenames("test/data/", time.Now(), 20))
	fmt.Println(todos)
	fmt.Println(logs)
}
