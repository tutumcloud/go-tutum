go-tutum
========

**Installation:**

	go get github.com/tutumcloud/go-tutum

**Usage Example:**
	
	package main

	import (
		"github.com/tutumcloud/go-tutum"
		"log"
	)

	func main() {
		//Set Credentials.
		tutum.User = "IchabodDee"
		tutum.Apikey = "5e4219650a71337419d9cd49872ab12b73dd4d3b"

		//Returns an array of Container objects.
		list := tutum.ListContainers()
		log.Println(list[0])

		//Returns an array of Application Objects
		applist := tutum.ListApplications()
		log.Println(applist[0].Image_name)
	}
