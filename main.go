//		Copyright (c) VoteBackend - All Rights Reserved
//
//	Unauthorized copying of this file, via any medium is strictly prohibited
//	Proprietary and confidential
//	Written by Ilyes Cherfaoui <ogfris@protonmail.com>, 2019

package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/OGFris/VoteBackend/routes"
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
	router.GET("/candidates", routes.Candidates)
	router.POST("/candidate", routes.PostCandidate)

	s := &fasthttp.Server{
		Handler:          router.Handler,
		DisableKeepalive: true,
	}

	utils.PanicError(s.ListenAndServe(":" + port))
}
