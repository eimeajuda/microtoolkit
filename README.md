```

Alternatively we can use multicast DNS with the built in MDNS registry for a zero dependency configuration. Just pass `--registry=mdns` to the below commands e.g. `server --registry=mdns` or `micro --registry=mdns list services`.

Run the greeter service
```shell
go run examples/greeter/server/main.go
```

List services
```shell
$ micro list services
consul
go.micro.srv.greeter
```

Get Service
```shell
$ micro get service go.micro.srv.greeter
service  go.micro.srv.greeter

version 1.0.0

Id	Address	Port	Metadata
go.micro.srv.greeter-34c55534-368b-11e6-b732-68a86d0d36b6	192.168.1.66	62525	server=rpc,registry=consul,transport=http,broker=http

Endpoint: Say.Hello
Metadata: stream=false

Request: {
	name string
}

Response: {
	msg string
}
```

Query service
```shell
$ micro query go.micro.srv.greeter Say.Hello '{"name": "John"}'
{
	"msg": "Hello John"
}
```

Read the [docs](https://github.com/micro/micro/tree/master/doc) to learn more.

## Community Contributions

Project		|	Description
-----		|	------
[Micro Dashboard](https://github.com/Margatroid/micro-dashboard)	|	Dashboard for microservices toolchain micro

## Sponsors

Open source development of Micro is sponsored by Sixt

<a href="https://blog.micro.mu/2016/04/25/announcing-sixt-sponsorship.html"><img src="https://micro.mu/sixt_logo.png" width=150px height="auto" /></a>