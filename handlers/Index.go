package handlers

import (
	"net/http"
	//"encoding/json"
//	"io"
//	"io/ioutil"
//	"github.com/antigloss/go/logger"
	"fmt"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "hello bla")
        
}

