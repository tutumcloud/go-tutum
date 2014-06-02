go-tutum
========

**Installation:**

	go get github.com/tutumcloud/go-tutum
	

**Auth:**
	

Manually set in your golang code

	tutum.User = "IchabodDee"
	tutum.Apikey = "yourApiKeyHere"

Store in a config file in ~/.tutum

	[auth]
	user = "username"
	apikey = "apikey"
	
Set the environment variables TUTUM_USER and TUTUM_APIKEY

**Example:**
	
	package main

	import (
		"github.com/tutumcloud/go-tutum"
		"log"
	)

	func main() {
		//Set Credentials.
		tutum.User = "username"
		tutum.Apikey = "yourApiKey"
	
		//Returns an array of Container structs.
		list, err := tutum.ListContainers()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(list[0])
	
		//Returns an array of Application structs.
		applist, err := tutum.ListApplications()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(applist[0].Image_name)
	
	}
