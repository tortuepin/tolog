package main

import "math/rand"
import "time"
import "flag"
import "fmt"
import "log"

var (
	aN = flag.Int("n", 6, "目の数")
)

func main() {
	flag.Parse()

	if *aN <= 0 {
		log.Fatal("n should > 0")
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(*aN))

}
