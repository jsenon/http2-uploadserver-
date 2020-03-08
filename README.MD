# Simple HTTP Uploader Server

A Simple Http server for testing upload with opentracing and prometheus implementation.

## Usage

```sh
Upload Server

Usage:
  http2-uploadserver [flags]
  http2-uploadserver [command]

Available Commands:
  help        Help about any command
  server      Start server

Flags:
      --config string      config file (default is $HOME/.http2-uploadserver.yaml)
      --debug              Set log level to Debug
      --disabletrace       Disable the trace
  -h, --help               help for http2-uploadserver
      --jaegerurl string   Set jaegger agent endpoint (without port, without http://)
  -t, --toggle             Help message for toggle
      --version            version
```

For example in order to launch the server without opentracing, and store img in /tmp

```sh
http2-uploadserver server --outputdir /tmp --disabletrace 
```


## Deployment


```sh
kubectl apply -f deployment/deployment.yml
kubectl apply -f deployment/svc.yml
```

## Build

Have a look on the `Makefile`, you can generate docker, binary, or artifact for other plateform.

### Build Docker 

```sh
make build
```

### Build locally

```sh
make build_local
```
