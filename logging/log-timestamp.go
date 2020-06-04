package main

import (
	"log"
)

// Initialize in init() so that it gets called before main
func init(){
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	log.Printf("Hello gophers. This log line has printed the entire log with the details")
}
