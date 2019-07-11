//		Copyright (c) VoteBackend - All Rights Reserved
//
//	Unauthorized copying of this file, via any medium is strictly prohibited
//	Proprietary and confidential
//	Written by Ilyes Cherfaoui <ogfris@protonmail.com>, 2019

package routes

import (
	"github.com/OGFris/VoteBackend/database"
	"github.com/OGFris/VoteBackend/utils"
	"github.com/valyala/fasthttp"
)

func Candidates(ctx *fasthttp.RequestCtx) {

	utils.WriteJson(&ctx.Response, database.AllCandidates())
}
