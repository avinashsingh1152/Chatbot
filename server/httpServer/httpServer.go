package httpServer

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/core"
	"server/utils"
)

type HttpServer struct {
	Core   core.Core
	Logger *log.Logger
	Router *mux.Router
}

func (h *HttpServer) Init() {
	h.registerRoutes()
}

func (h *HttpServer) registerRoutes() {
	h.Router.HandleFunc("/", h.HandlePing()).Methods("GET")
	h.Router.HandleFunc("/checkGrpcConnection", h.HandleGrpcConnection()).Methods("GET")
	h.Router.HandleFunc("/uploadImage", h.HandleUploadImage()).Methods("POST")
	h.Router.HandleFunc("/getImages", h.HandleGetImage()).Methods("GET")
	h.Router.HandleFunc("/conversation", h.HandleConversation()).Methods("POST")
}

func (h HttpServer) HandlePing() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			utils.HttpSuccessWith2XX("/ping", http.StatusOK, w, r)
			return
		default:
			utils.HttpSuccessWith4XX("bad request", http.StatusBadRequest, w, r)
		}
	}
}

func (h HttpServer) HandleGrpcConnection() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			resp, err := h.Core.GetGrpcPing()
			if err != nil {
				utils.HttpSuccessWith4XX(resp, http.StatusBadRequest, w, r)
				return
			}
			utils.HttpSuccessWith2XX(resp, http.StatusOK, w, r)
			return
		default:
			utils.HttpSuccessWith4XX("bad request", http.StatusBadRequest, w, r)
		}
	}
}

func (h HttpServer) HandleUploadImage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.UploadImage(w, r)
			return
		default:
			utils.HttpSuccessWith4XX("bad request", http.StatusBadRequest, w, r)
		}
	}
}

func (h HttpServer) HandleGetImage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetImages(w, r)
			return
		default:
			utils.HttpSuccessWith4XX("bad request", http.StatusBadRequest, w, r)
		}
	}
}

func (h HttpServer) HandleConversation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.Conversation(w, r)
			return
		default:
			utils.HttpSuccessWith4XX("bad request", http.StatusBadRequest, w, r)
		}
	}
}
