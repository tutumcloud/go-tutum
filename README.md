go-tutum
========

**Installation:**

	go get github.com/tutumcloud/go-tutum
	

**Auth:**
	

Manually set in your golang code

	tutum.User = "IchabodDee"
	tutum.ApiKey = "yourApiKeyHere"

Store in a config file in ~/.tutum

	[auth]
	user = "username"
	apikey = "apikey"
	
Set the environment variables TUTUM_USER and TUTUM_APIKEY

**Example:**
	
	package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns an array of Container structs.
	    list, err := tutum.ListContainers()
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(list[0])
	}
