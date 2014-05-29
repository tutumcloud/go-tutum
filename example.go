package main

import (
	"github.com/tutumcloud/go-tutum"
	"log"
)

func main() {

	gotutum.User = "IchabodDee"
	gotutum.Apikey = "5e4219650a71337419d9cd49872ab12b73dd4d3b"
	inter := gotutum.ListContainers()
	log.Println("Printing What came back")
	log.Println(inter)
	// gotutum.ListContainers()
}
