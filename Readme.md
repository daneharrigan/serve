# serve

A simple HTTP server for local development of static content. Serve binds to
port 8080, but if 8080 is already in use it will look for an available port
by incrementing the port by 1. Serve will try up to 5 ports by default.

All content is served with the `Cache-Control: max-age=0` header and
`Last-Modified` set to the current time.

## Install

```bash
$ go get github.com/daneharrigan/serve
$ go install github.com/daneharrigan/serve
```

## Usage

Serve my current directory over port 8080:

```bash
$ serve
fn=start port=8080 directory=foo
```

Serve my current directory over port 5000:

```bash
$ serve -p 5000
fn=start port=5000 directory=foo
```

Serve a diffrent directory:

```bash
$ serve -d ../foo
fn=start port=8080 directory=foo
```

Try binding to a port up to 10 times:

```bash
$ serve -r 10
fn=start port=8080 directory=foo
fn=start error="port already in use"
fn=start port=8081 directory=foo
fn=start error="port already in use"
fn=start port=8082 directory=foo
```
