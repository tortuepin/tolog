package main

import "testing"
import "time"
import "os"
import "io/ioutil"
import "bufio"
import "encoding/json"
import "strings"
import "./libs"

const dir = "./test/data/"
const file = "190226.md"

var date = time.Date(2019, 2, 27, 0, 0, 0, 0, time.Local)

// Utils

func TestFileExist(t *testing.T) {
	if !tolog.Exists(dir + file) {
		t.Fatal(dir + file + " is not exist.")
	}
}

func TestGetFilenames(t *testing.T) {
	ans := []string{"./test/data/190209.md", "./test/data/190211.md", "./test/data/190212.md", "./test/data/190221.md", "./test/data/190226.md"}
	files := tolog.GetFilenames(dir, time.Date(2019, 2, 27, 0, 0, 0, 0, time.Local), 20)

	for i, _ := range ans {
		if !(ans[i] == files[i]) {
			t.Fatal("GetFilenames is not work.\n", "want '", ans[i], "'\n got '", files[i], "'")
		}
	}

}

// Todo

func TestTodoReader(t *testing.T) {
	ans := []tolog.TodoItem{{"todo1", "", true, 0}, {"todo1_1", "", true, 1}, {"todo2", "", false, 0}, {"todo3", "@tag1", false, 0}, {"todo4", "@tag1", false, 0}, {"todo5", "", false, 0}, {"todo6", "", false, 0}}

	f, err := os.Open(dir + file)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	todos := []tolog.TodoItem{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == tolog.TodoHeader {
			todos = tolog.TodoReader(scanner)
		}
	}

	for i, _ := range ans {
		if !(ans[i].Title == todos[i].Title) ||
			!(ans[i].Tag == todos[i].Tag) ||
			!(ans[i].Done == todos[i].Done) ||
			!(ans[i].Depth == todos[i].Depth) {
			t.Fatal("TodoReader is not work", ans[i].Depth)
		}
	}
}

func TestTodoGetActive(t *testing.T) {
	expect := []tolog.TodoItem{{"todo2", "", false, 0}, {"todo3", "@tag1", false, 0}, {"todo4", "@tag1", false, 0}, {"todo5", "", false, 0}, {"todo6", "", false, 0}, {"todo8", "", false, 0}}
	allitems := tolog.GetAllItems(dir, date, 20)
	alltodos := []tolog.TodoItem{}
	for _, a := range allitems {
		alltodos = append(alltodos, a.Todo...)
	}
	active := tolog.TodoGetActive(alltodos)
	for i, _ := range active {
		if active[i].Title != expect[i].Title ||
			active[i].Tag != expect[i].Tag ||
			active[i].Done != expect[i].Done ||
			active[i].Depth != expect[i].Depth {
			t.Fatal("want ", expect, "\n got ", active)
		}
	}
}

func TestGetAllItems(t *testing.T) {
	// TODO これTODO関係ないので別のとこに移動
	a := tolog.GetAllItems(dir, time.Date(2019, 2, 27, 0, 0, 0, 0, time.Local), 20)
	jsonByte, err := json.Marshal(a)

	data, err := ioutil.ReadFile(dir + "t.GetAllItems.JSON.golden")
	if err != nil {
		t.Fatal(err)
	}

	expect := strings.Trim(string(data), "\n ")
	actual := strings.Trim(string(jsonByte), "\n ")

	if expect != actual {
		t.Fatal("errr\n ", "want ", string(data)[0:10], "\ngot ", string(jsonByte)[0:10])
	}
}
