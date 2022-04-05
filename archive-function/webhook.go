package main

import (
	"fmt"
	"net/http"
)

type WebhookHandler struct {
}

func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handleTFCWebhookEvent(w, r)
	default:
		fmt.Fprintf(w, "GET is the only supported method on this path")
	}
}

func handleTFCWebhookEvent(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	return
}