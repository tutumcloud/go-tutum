go-tutum
========

Go library for Tutum's API. Full documentation available at [https://docs.tutum.co/v2/api/?golang](https://docs.tutum.co/v2/api/?golang
)

##Set up

**Installation:**

In order to install the Tutum Go library, you can use :

	go get github.com/tutumcloud/go-tutum


**Auth:**

In order to be able to make requests to the API, you should first obtain an ApiKey for your account. For this, log into Tutum, click on the menu on the upper right corner of the screen and select **Get Api Key**. 

You can use your ApiKey with the Go library in any of the following ways:

Manually set in your golang code

	tutum.User = "yourUsernameHere"
	tutum.ApiKey = "yourApiKeyHere"

Store in a config file in ~/.tutum

	[auth]
	user = "username"
	apikey = "apikey"

Set the environment variables TUTUM_USER and TUTUM_APIKEY


##Function and examples

**Services**

- ListServices() : returns all the services in a JSON object

######Example
```
	    list, err := tutum.ListServices()
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(list)
```

- GetService(uuid string) : returns the details of the service in a JSON object

######Example

```
	    service, err := tutum.GetService("0a6e3c0d-cefd-4347-930b-d112d991ba52")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(service)
```


- CreateService(newService []byte) : returns the newly created service in a JSON object

######Example
```
	    newservice, err := tutum.CreateService([]byte(`{"image": "tutum/hello-world", "name": "my-new-app", "target_num_containers": 2}`))
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(newservice)
```

- GetServiceLogs(uuid string) : returns the logs of the service

######Example
```
	    logs, err := tutum.GetServiceLogs("37062446-aaef-46d6-b2d0-0cb4c1bab8cf")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(logs)
```

- UpdateService(uuid string, updatedService []byte) : returns the updated service in a JSON object

######Example
```
	    updatedService, err := tutum.UpdateService("0a6e3c0d-cefd-4347-930b-d112d991ba52", []byte(`{"container_envvars": [{"key": "NAME", "value": "New Service"}]}`))
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(updatedService)
```


- StartService(uuid string) : returns the newly started service in a JSON object

######Example
```
	    start, err := tutum.StartService("0a6e3c0d-cefd-4347-930b-d112d991ba52")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(start)
```

- StopService(uuid string) : returns the newly stopped service in a JSON object

- RedeployService(uuid string) : returns the newly redeployed service in a JSON object

- TerminateService(uuid string) : returns the newly terminated service in a JSON object



**Containers**

- ListContainers() : returns all the containers in a JSON object

######Example

```
	    list, err := tutum.ListContainers()
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(list)
```

- GetContainer(uuid string) : returns the details of the container in a JSON object

######Example

```
	    container, err := tutum.GetContainer("fcf37b7f-2df5-4a45-9acb-fd11dca3d562")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(container)
```

- GetContainerLogs(uuid string) : returns the logs of the container

######Example

```
	    logs, err := tutum.GetContainerLogs("fcf37b7f-2df5-4a45-9acb-fd11dca3d562")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(logs)
```


- StartContainer(uuid string) : returns the newly started container in a JSON object


######Example

```
	    start, err := tutum.StartContainer("fcf37b7f-2df5-4a45-9acb-fd11dca3d562")
	    if err != nil {
	        log.Fatal(err)
	    }
	    log.Println(start)
```


- StopContainer(uuid string) : returns the newly stopped container in a JSON object


- RedeployContainer(uuid string) : returns the newly redeployed container in a JSON object

- TerminateContainer(uuid string) : returns the newly terminated container in a JSON object

**NodeClusters**

- ListNodeClusters() : returns all the nodeclusters in a JSON object

- GetNodeCluster(uuid string) : returns the details of a specific container in a JSON object

- CreateNodeCluster(newCluster []byte) : returns the newly created nodecluster in a JSON object

- DeployNodeCluster(uuid string) : returns the newly deployed nodecluster in a JSON object

- UpdateNodeCluster(uuid string, updatedNode []byte) : returns the newly updated nodecluster in a JSON object

- TerminateNodeCluster(uuid string) : returns the newly terminated nodecluster in a JSON object


The complete API Documentation is available [here](https://docs.tutum.co/v2/api/) with additionnal examples written in Go.
