package main

import (
	_ "GoInAction/sample/matchers"
	"GoInAction/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
