package handler

import (
	"net/http"

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

func (h *IpInfoHandler) GetIpInfo(w http.ResponseWriter, r *http.Request) {
	statusCode, response := h.Service.GetIpInfo(r.Context(), r)
	utils.EncodeResponse[service.GetIpInfoResponse](w, statusCode, response)
	h.Logger.Info(r.Context(), "GetIpInfo", "responseStatusCode", statusCode)
}

func (h *IpInfoHandler) GetHeaderByName(w http.ResponseWriter, r *http.Request) {
	headerName := r.PathValue("headerName")
	statusCode, response := h.Service.GetHeaderByName(headerName, r.Context(), r)
	utils.EncodeResponse[service.GetIpInfoResponse](w, statusCode, response)
	h.Logger.Info(r.Context(), "GetHeaderByName", "responseStatusCode", statusCode)
}
