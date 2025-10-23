package main

import (
	"net/http"

	"github.com/amieldelatorre/myip/handler"
)

func RegisterRoutes(mux *http.ServeMux, m handler.Middlware, ipInfoHandler handler.IpInfoHandler) {
	getIndex := m.RecoverPanic(m.AddRequestId(http.HandlerFunc(ipInfoHandler.Index)))
	getHeaderByNameHandlerFunc := m.RecoverPanic(m.AddRequestId(http.HandlerFunc(ipInfoHandler.GetHeaderByName)))

	mux.Handle("GET /", getIndex)
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	mux.Handle("GET /{headerName}", getHeaderByNameHandlerFunc)
}
