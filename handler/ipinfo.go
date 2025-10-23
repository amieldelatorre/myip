package handler

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/amieldelatorre/myip/service"
	"github.com/amieldelatorre/myip/utils"
)

type IpInfoHandler struct {
	Logger  utils.CustomJsonLogger
	Service service.IpInfoService
}

func NewIpInfoHandler(logger utils.CustomJsonLogger, svc service.IpInfoService) IpInfoHandler {
	return IpInfoHandler{Logger: logger, Service: svc}
}

func (h *IpInfoHandler) Index(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	if userAgent == "" || strings.HasPrefix(userAgent, "curl") {
		h.GetIpInfo(w, r)
		return
	}

	h.GetIpInfoHtml(w, r)
}

func (h *IpInfoHandler) GetIpInfo(w http.ResponseWriter, r *http.Request) {
	statusCode, response := h.Service.GetIpInfo(r.Context(), r)
	utils.EncodeResponse[service.GetIpInfoResponse](w, statusCode, response)
	h.Logger.Info(r.Context(), "GetIpInfo", "responseStatusCode", statusCode)
}

func (h *IpInfoHandler) GetIpInfoHtml(w http.ResponseWriter, r *http.Request) {
	_, response := h.Service.GetIpInfo(r.Context(), r)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		h.Logger.Error(r.Context(), "GetIpInfoHtml", "error", err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, response)
	if err != nil {
		h.Logger.Error(r.Context(), "GetIpInfoHtml", "error", err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
}

func (h *IpInfoHandler) GetHeaderByName(w http.ResponseWriter, r *http.Request) {
	headerName := r.PathValue("headerName")
	statusCode, response := h.Service.GetHeaderByName(headerName, r.Context(), r)
	utils.EncodeResponse[service.GetIpInfoResponse](w, statusCode, response)
	h.Logger.Info(r.Context(), "GetHeaderByName", "responseStatusCode", statusCode)
}
