//		Copyright (c) VoteBackend - All Rights Reserved
//
//	Unauthorized copying of this file, via any medium is strictly prohibited
//	Proprietary and confidential
//	Written by Ilyes Cherfaoui <ogfris@protonmail.com>, 2019

package utils

import (
	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type FormError struct {
	Message string `json:"message"`
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteJson(w *fasthttp.Response, data interface{}) {
	w.Header.Set("Content-Type", "application/json")
	PanicError(jsoniter.NewEncoder(w.BodyWriter()).Encode(data))
}
