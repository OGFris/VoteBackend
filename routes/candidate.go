//		Copyright (c) VoteBackend - All Rights Reserved
//
//	Unauthorized copying of this file, via any medium is strictly prohibited
//	Proprietary and confidential
//	Written by Ilyes Cherfaoui <ogfris@protonmail.com>, 2019

package routes

import (
	"github.com/OGFris/VoteBackend/database"
	"github.com/valyala/fasthttp"
	"net/http"
)

func GetCandidate(ctx *fasthttp.RequestCtx) {
	// TODO
}

func PostCandidate(ctx *fasthttp.RequestCtx) {
	name := string(ctx.PostArgs().Peek("name"))
	party := string(ctx.PostArgs().Peek("party"))
	description := string(ctx.PostArgs().Peek("description"))

	if name == "" || party == "" || description == "" {
		ctx.Error("please provide all the information required correctly to do this", http.StatusBadRequest)

		return
	}

	if database.Exist(name) {
		ctx.Error("name already exist", http.StatusBadRequest)

		return
	}

	err := database.Create(name, party, description)
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(http.StatusOK)
}
