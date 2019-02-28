package main

import "testing"
import "./libs"

const dir = "./test/data/"
const file = "190226.md"

func TestFileExist(t *testing.T) {
	if !tolog.Exists(dir + file) {
		t.Fatal(dir + file + " is not exist.")
	}
}
