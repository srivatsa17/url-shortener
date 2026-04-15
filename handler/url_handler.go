package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/srivatsa17/url-shortener/model"
	"github.com/srivatsa17/url-shortener/service"
	"github.com/srivatsa17/url-shortener/utils"
)

const (
	HeaderContentType = "Content-Type"
	JsonContentType   = "application/json"
)

type URLHandler struct {
	svc service.URLService
}

func NewURLHandler(s service.URLService) *URLHandler {
	return &URLHandler{
		svc: s,
	}
}

func (u *URLHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(HeaderContentType, JsonContentType)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(
		map[string]string{
			"status": "healthy",
		}); err != nil {
		log.Printf("json encoding failed: %v", err)
	}
}

func (u *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	// Disallow requests other than POST method
	if r.Method != http.MethodPost {
		w.Header().Set(HeaderContentType, JsonContentType)
		w.WriteHeader(http.StatusMethodNotAllowed)

		if err := json.NewEncoder(w).Encode(
			model.ErrorResponse{
				Error: "Method not allowed",
				Code:  http.StatusMethodNotAllowed,
			}); err != nil {
			log.Printf("json encoding failed: %v", err)
		}

		return
	}

	// Decode the request body
	var request model.URLShortenRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.URL == "" {
		w.Header().Set(HeaderContentType, JsonContentType)
		w.WriteHeader(http.StatusBadRequest)

		if err := json.NewEncoder(w).Encode(
			model.ErrorResponse{
				Error: "Invalid request body: URL is required",
				Code:  http.StatusBadRequest,
			}); err != nil {
			log.Printf("json encoding failed: %v", err)
		}
		return
	}

	// Validate URL format
	if !utils.IsValidURL(request.URL) {
		w.Header().Set(HeaderContentType, JsonContentType)
		w.WriteHeader(http.StatusBadRequest)

		if err := json.NewEncoder(w).Encode(
			model.ErrorResponse{
				Error: "Invalid URL format",
				Code:  http.StatusBadRequest,
			}); err != nil {
			log.Printf("json encoding failed: %v", err)
		}

		return
	}

	// Get the response
	response, err := u.svc.ShortenURL(request.URL)
	if err != nil {
		log.Printf("Error shortening URL: %v", err)
		w.Header().Set(HeaderContentType, JsonContentType)
		w.WriteHeader(http.StatusInternalServerError)

		if err := json.NewEncoder(w).Encode(
			model.ErrorResponse{
				Error: "Failed to shorten URL",
				Code:  http.StatusInternalServerError,
			}); err != nil {
			log.Printf("json encoding failed: %v", err)
		}

		return
	}

	w.Header().Set(HeaderContentType, JsonContentType)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("json encoding failed: %v", err)
	}
}

func (u *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Query().Get("code")
	if shortCode == "" {
		w.Header().Set(HeaderContentType, JsonContentType)
		w.WriteHeader(http.StatusBadRequest)

		if err := json.NewEncoder(w).Encode(
			model.ErrorResponse{
				Error: "Missing short code",
				Code:  http.StatusBadRequest,
			}); err != nil {
			log.Printf("json encoding failed: %v", err)
		}

		return
	}

	longURL, err := u.svc.GetURL(shortCode)
	if err != nil {
		w.Header().Set(HeaderContentType, JsonContentType)
		w.WriteHeader(http.StatusNotFound)

		if err := json.NewEncoder(w).Encode(
			model.ErrorResponse{
				Error: "Short code not found",
				Code:  http.StatusNotFound,
			}); err != nil {
			log.Printf("json encoding failed: %v", err)
		}

		return
	}

	// Redirect to long URL
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}
