package utils

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
)

func HttpSuccessWith2XX(data interface{}, statusCode int, w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}{
		Success: true,
		Data:    data,
	})

	if err != nil {
		log.Printf("error while marshing data from HttpSuccessWith2XX")
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func HttpSuccessWith2XXWithPagination(data interface{}, pageInfo interface{}, statusCode int, w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(struct {
		Success  bool        `json:"success"`
		Data     interface{} `json:"data"`
		PageInfo interface{} `json:"page_info"`
	}{
		Success:  true,
		Data:     data,
		PageInfo: pageInfo,
	})

	if err != nil {
		log.Printf("error while marshing data from HttpSuccessWith2XX")
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func HttpSuccessWith4XX(data interface{}, statusCode int, w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(struct {
		Success bool        `json:"failed"`
		Data    interface{} `json:"data"`
	}{
		Success: true,
		Data:    data,
	})

	if err != nil {
		log.Printf("error while marshing data from HttpSuccessWith2XX")
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func GetIPAddress(r *http.Request) string {

	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(ips[0])
	}

	// Fallback to the remote address in the mapper, which might be the proxy's address
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}
