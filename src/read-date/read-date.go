package main

import "fmt"
import "time"

func main() {
	t := time.Now()
  fmt.Println(t.Format("20060102150405"));
}