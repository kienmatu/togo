package handlers

import (
	"kienmatu/go-todos/utils"
	"net/http"
)

type HomeHandler struct {
}

func (h *HomeHandler) Index(rw http.ResponseWriter, r *http.Request){
	utils.Response(rw, "Service is running!")
}

func NewHome() *HomeHandler{
	return &HomeHandler{}
}