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

Manually set in your Go code

	tutum.User = "yourUsernameHere"
	tutum.ApiKey = "yourApiKeyHere"

Store in a config file in ~/.tutum

	[auth]
	user = "username"
	apikey = "apikey"

Set the environment variables TUTUM_USER and TUTUM_APIKEY

##Examples

**Creating and deploying a NodeCluster**

```
nodecluster, err := tutum.CreateNodeCluster(`{"name": "my_cluster", "region": "/api/v1/region/digitalocean/lon1/", "node_type": "/api/v1/nodetype/digitalocean/1gb/", "disk": 60}`)

if err != nil {
  log.Println(err)
}


nodecluster.Deploy()
```

**Creating and starting a Stack**

```
stack, err := tutum.CreateStack(`{
    "name": "my-new-stack",
    "services": [
        {
            "name": "hello-word",
            "image": "tutum/hello-world",
            "target_num_containers": 2,
            "linked_to_service": [
                {
                    "to_service": "database",
                    "name": "DB"
                }
            ]
        },
        {
            "name": "database",
            "image": "tutum/mysql"
        }
    ]
}`)

if err != nil {
  log.Println(err)
}

stack.Start()
```

**Listing running containers**

```
containers, err := tutum.ListContainers()

if err != nil {
	log.Println(err)
}

log.Println(containers)
```

**Stopping a running service**

```
service, err := tutum.GetService("7eaf7fff-882c-4f3d-9a8f-a22317ac00ce")

if err != nil {
	log.Println(err)
}

service.Stop()
```

**Stream events**

```
tutum.Stream()
```

**Apply a function at each new event**

```
tutum.OnEvent(function_name)

```

**Note** 

Add error logs while applying actions (Start, Stop, Redeploy, ...) on containers, services, nodeclusters or nodes is done like the following :

```
service, err := tutum.GetService("7eaf7fff-882c-4f3d-9a8f-a22317ac00ce")

if err != nil {
	log.Println(err)
}

if err = service.Start(); err != nil {
	log.Println(err)
}

```

The complete API Documentation is available [here](https://docs.tutum.co/v2/api/) with additional examples written in Go.
