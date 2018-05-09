# Session middleware for [Gear](https://github.com/teambition/gear)

[![Build][Build-Status-Image]][Build-Status-Url] [![Coverage][Coverage-Image]][Coverage-Url] [![ReportCard][reportcard-image]][reportcard-url] [![GoDoc][godoc-image]][godoc-url] [![License][license-image]][license-url]

## Quick Start

### Download and install

```bash
$ go get -u -v github.com/go-session/gear-session
```

### Create file `server.go`

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/go-session/gear-session"
	"github.com/teambition/gear"
	"gopkg.in/session.v2"
)

func main() {
	app := gear.New()

	app.Use(gearsession.New(
		session.SetCookieName("session_id"),
		session.SetSign([]byte("sign")),
	))

	router := gear.NewRouter()

	router.Get("/", func(ctx *gear.Context) error {
		store := gearsession.FromContext(ctx)
		store.Set("foo", "bar")
		err := store.Save()
		if err != nil {
			return gear.ErrInternalServerError.From(err)
		}
		return ctx.Redirect("/foo")
	})

	router.Get("/foo", func(ctx *gear.Context) error {
		store := gearsession.FromContext(ctx)
		foo, ok := store.Get("foo")
		if !ok {
			return gear.ErrNotFound
		}
		return ctx.End(http.StatusOK, []byte(fmt.Sprintf("foo:%s", foo)))
	})
	app.UseHandler(router)

	app.Listen(":8080")
}
```

### Build and run

```bash
$ go build server.go
$ ./server
```

### Open in your web browser

<http://localhost:8080>

    foo:bar


## MIT License

    Copyright (c) 2018 Lyric

[Build-Status-Url]: https://travis-ci.org/go-session/gear-session
[Build-Status-Image]: https://travis-ci.org/go-session/gear-session.svg?branch=master
[Coverage-Url]: https://coveralls.io/github/go-session/gear-session?branch=master
[Coverage-Image]: https://coveralls.io/repos/github/go-session/gear-session/badge.svg?branch=master
[reportcard-url]: https://goreportcard.com/report/github.com/go-session/gear-session
[reportcard-image]: https://goreportcard.com/badge/github.com/go-session/gear-session
[godoc-url]: https://godoc.org/github.com/go-session/gear-session
[godoc-image]: https://godoc.org/github.com/go-session/gear-session?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg