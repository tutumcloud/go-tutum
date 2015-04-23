go-tutum
========

##Set up

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


##Endpoints / Methods and examples

**Services**

- ListServices() : returns all the services in a JSON object

######Example 
```
package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns an array of Service structs.
	    list, err := tutum.ListServices()
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(list)
	}
```

- GetService(uuid string) : returns the details of the service in a JSON object

######Example 

```
package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns a JSON of Service structs.
	    service, err := tutum.GetService("0a6e3c0d-cefd-4347-930b-d112d991ba52")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(service)
	}
```


- CreateService(newService []byte) : returns the newly created service in a JSON object

######Example 
```
package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns a JSON of Service structs.
	    newservice, err := tutum.CreateService([]byte(`{"image": "tutum/hello-world", "name": "my-new-app", "target_num_containers": 2}`))
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(newservice)
	}
```

- GetServiceLogs(uuid string) : returns the logs of the service

######Example 
```
package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns a string.
	    logs, err := tutum.GetServiceLogs("37062446-aaef-46d6-b2d0-0cb4c1bab8cf")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(logs)
	}
```

- UpdateService(uuid string, updatedService []byte) : returns the updated service in a JSON object

######Example
```
package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns a JSON of Service structs.
	    updatedService, err := tutum.UpdateService("0a6e3c0d-cefd-4347-930b-d112d991ba52", []byte(`{"container_envvars": [{"key": "NAME", "value": "New Service"}]}`))
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(updatedService)
	}
``` 


- StartService(uuid string) : returns the newly started service in a JSON object

######Example 
```
package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns a JSON of Service structs.
	    start, err := tutum.StartService("0a6e3c0d-cefd-4347-930b-d112d991ba52")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(start)
	}
```

- StopService(uuid string) : returns the newly stopped service in a JSON object 

- RedeployService(uuid string) : returns the newly redeployed service in a JSON object

- TerminateService(uuid string) : returns the newly terminated service in a JSON object



**Containers**

- ListContainers() : returns all the containers in a JSON object

######Example 

```
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
	    log.Println(list)
	}
```

- GetContainer(uuid string) : returns the details of the container in a JSON object

######Example 

```
package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns a JSON of Container structs.
	    container, err := tutum.GetContainer("fcf37b7f-2df5-4a45-9acb-fd11dca3d562")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(container)
	}
```

- GetContainerLogs(uuid string) : returns the logs of the container

######Example

```
package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns a string.
	    logs, err := tutum.GetContainerLogs("fcf37b7f-2df5-4a45-9acb-fd11dca3d562")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(logs)
	}
```


- StartContainer(uuid string) : returns the newly started container in a JSON object


######Example 

```
package main

	import (
	    "log"
	    "github.com/tutumcloud/go-tutum/tutum"
	)

	func main() {

	    //Set Credentials.
	    tutum.User = "tifayuki"
	    tutum.ApiKey = "d24c522c3651ad2516c848c268aecb264be34597"

	    //Returns a JSON of Container structs.
	    start, err := tutum.StartContainer("fcf37b7f-2df5-4a45-9acb-fd11dca3d562")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(start)
	}
```


- StopContainer(uuid string) : returns the newly stopped container in a JSON object


- RedeployContainer(uuid string) : returns the newly redeployed container in a JSON object

- TerminateContainer(uuid string) : returns the newly terminated container in a JSON object

