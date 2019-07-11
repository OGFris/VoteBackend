package main

import (
	"github.com/OGFris/VoteBackend/utils"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := fasthttprouter.New()
	s := &fasthttp.Server{
		Handler:          router.Handler,
		DisableKeepalive: true,
	}

	utils.PanicError(s.ListenAndServe(":" + port))
}
