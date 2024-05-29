package service

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/amieldelatorre/myip/utils"
)

type GetIpInfoResponse struct {
	RequestHeaders map[string][]string `json:"requestHeaders,omitempty"`
	SourceIp       string              `json:"sourceIp,omitempty"`
	Errors         map[string]string   `json:"errors,omitempty"`
}

type GetHeaderByNameResponse struct {
	Header string            `json:"header,omitempty"`
	Value  string            `json:"value,omitempty"`
	Errors map[string]string `json:"errors,omitempty"`
}

type IpInfoService struct {
	Logger utils.CustomJsonLogger
}

func NewIpInfoService(logger utils.CustomJsonLogger) IpInfoService {
	return IpInfoService{Logger: logger}
}

func (i *IpInfoService) GetIpInfo(ctx context.Context, r *http.Request) (int, GetIpInfoResponse) {
	response := GetIpInfoResponse{
		Errors:         map[string]string{},
		RequestHeaders: map[string][]string{},
	}

	sourceIp, err := getRequestIpAddress(r)
	if err != nil {
		i.Logger.Error(ctx, "Could not get request ip", "error", err)
		response.Errors["Server"] = "Server could not handle your request, please try again later."
		return http.StatusInternalServerError, response
	}

	for key, value := range r.Header {
		response.RequestHeaders[key] = value
	}

	response.SourceIp = sourceIp

	return http.StatusOK, response
}

func (i *IpInfoService) GetHeaderByName(headerName string, ctx context.Context, r *http.Request) (int, GetHeaderByNameResponse) {
	headerValue := r.Header.Get(headerName)

	responseStatus := http.StatusOK
	response := GetHeaderByNameResponse{
		Header: headerName,
		Value:  headerValue,
		Errors: map[string]string{},
	}

	if headerValue == "" {
		response.Errors[headerName] = "Could not be found"
		responseStatus = http.StatusNotFound
	}

	return responseStatus, response
}

func getRequestIpAddress(r *http.Request) (string, error) {
	sourceIp := r.RemoteAddr

	forwardedIps := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
	if len(forwardedIps) > 0 {
		ip := forwardedIps[0]
		if ip != "" {
			sourceIp = ip
		}
	}

	ip, _, err := net.SplitHostPort(sourceIp)
	if err != nil {
		_, isNetAddrError := err.(*net.AddrError)
		if isNetAddrError {
			return sourceIp, nil
		} else {
			return "", err
		}
	}

	return ip, nil
}
