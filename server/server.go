package server

import (
	"errors"
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

func handleCORS(ctx *fasthttp.RequestCtx) {
	// Add headers CORS
	ctx.Response.Header.Add(
		"Access-Control-Allow-Origin",
		"*",
	)
	ctx.Response.Header.Add(
		"Access-Control-Allow-Credentials",
		"true",
	)
	ctx.Response.Header.Add(
		"Access-Control-Allow-Headers",
		"*",
	)
	ctx.Response.Header.Add(
		"Access-Control-Max-Age",
		"86400",
	)
	ctx.Response.Header.Add(
		"Access-Control-Allow-Methods",
		"POST, GET, OPTIONS, DELETE, PUT",
	)
	// Response
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

func Run() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Method()) != fasthttp.MethodOptions {
			ctx.Error("Method not allowed", fasthttp.StatusMethodNotAllowed)
		}
		handleCORS(ctx)
	}

	err := fasthttp.ListenAndServe(":3333", requestHandler)
	if errors.Is(err, fasthttp.ErrConnectionClosed) {
		fmt.Println("CORS Server closed")
	} else if err != nil {
		fmt.Printf("CORS Server Err: %s", err.Error())
		os.Exit(1)
	}
}
