package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Thank you for pinging the Index!\n Get the API List at /apiList")

}
