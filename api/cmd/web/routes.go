package main

import (
	"net/http"

	"github.com/erik-jenkins/goblog/api/cmd/web/config"
)

func routes(env *config.Env) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("OK"))
	})

	return mux
}
