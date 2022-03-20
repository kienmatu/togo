package handlers

import (
	"kienmatu/go-todos/utils"
	"net/http"
)

type AuthHandler struct {
}

func (h *AuthHandler) Index(rw http.ResponseWriter, r *http.Request){
	utils.Response(rw, "AuthHandler is running!")
}

func NewAuthHandler() *AuthHandler{
	return &AuthHandler{}
}