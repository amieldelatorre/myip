package main

import (
	"net/http"

	"github.com/amieldelatorre/myip/handler"
)

func RegisterRoutes(mux *http.ServeMux, m handler.Middlware, ipInfoHandler handler.IpInfoHandler) {
	getIpInfoHandlerFunc := m.RecoverPanic(m.AddRequestId(http.HandlerFunc(ipInfoHandler.GetIpInfo)))
	getHeaderByNameHandlerFunc := m.RecoverPanic(m.AddRequestId(http.HandlerFunc(ipInfoHandler.GetHeaderByName)))

	mux.Handle("GET /", getIpInfoHandlerFunc)
	mux.Handle("GET /{headerName}", getHeaderByNameHandlerFunc)
}
