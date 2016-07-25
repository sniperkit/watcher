package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vsouza/watcher/config"
	"github.com/vsouza/watcher/db"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)
	c := config.GetConfig()
	fmt.Println(c.GetString("db.host"))
	db.Init()
	d := db.GetDB()
	fmt.Println(d)
}
