package httpServer

import (
	"encoding/json"
	"net/http"
	"server/Global"
	"server/mapper"
	"server/utils"
	"strconv"
)

func (h HttpServer) UploadImage(w http.ResponseWriter, r *http.Request) {

	ipAddress := utils.GetIPAddress(r)
	err := r.ParseMultipartForm(Global.MAX_FILE_SIZE)
	if err != nil {
		h.Logger.Println("Error while ParseMultipartForm:", err)
		panic(err)
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		h.Logger.Println("Error Retrieving the File: ", err)
		utils.HttpSuccessWith4XX("Error retrieving the file", http.StatusBadRequest, w, r)
		return
	}
	defer file.Close()

	// Check the file size using handler.Size
	if handler.Size > Global.MAX_FILE_SIZE {
		utils.HttpSuccessWith4XX("The uploaded file is too large. Please upload a file of size 15MB or less.", http.StatusBadRequest, w, r)
		return
	}

	err = h.Core.UploadFile(ipAddress, handler.Filename, file)
	if err != nil {
		h.Logger.Println("Error uploading the File: ", err)
		http.Error(w, "Error uploading the file", http.StatusInternalServerError)
		return
	}

	utils.HttpSuccessWith2XX("file uploaded successfully", http.StatusOK, w, r)
}

func (h HttpServer) GetImages(w http.ResponseWriter, r *http.Request) {

	ipAddress := utils.GetIPAddress(r)
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		pageSize = 10
	}

	cursor, err := strconv.Atoi(r.URL.Query().Get("cursor"))
	if err != nil {
		cursor = 0
	}

	resp, total, err := h.Core.GetImages(pageSize, cursor, ipAddress)
	if err != nil {
		h.Logger.Println("Error uploading the File: ", err)
		http.Error(w, "Error uploading the file", http.StatusInternalServerError)
		return
	}

	page := map[string]interface{}{
		"hasMore": int64(cursor+pageSize) < total,
		"total":   total,
	}
	utils.HttpSuccessWith2XXWithPagination(resp, page, http.StatusOK, w, r)
}

func (h HttpServer) Conversation(w http.ResponseWriter, r *http.Request) {
	ipAddress := utils.GetIPAddress(r)

	var qa mapper.QuestionAnswer
	if err := json.NewDecoder(r.Body).Decode(&qa); err != nil {
		h.Logger.Println("invalid payload: ", err)
		utils.HttpSuccessWith4XX("invalid payload", http.StatusBadRequest, w, r)
		return
	}

	resp, err := h.Core.Conversation(ipAddress, qa.Question)
	if err != nil {
		h.Logger.Println("Error getting the Conversation: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.HttpSuccessWith2XX(resp, http.StatusOK, w, r)
}
