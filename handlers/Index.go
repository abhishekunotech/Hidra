package handlers

import (
	"fmt"
	"net/http"
)


// Function exposes a simple landing page
//
// No Business Logic
//
// No Output
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Thank you for pinging the Index!\n Get the API List at /apiList")

}
