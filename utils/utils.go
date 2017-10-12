package utils

import (
"net/http"
"github.com/Unotechsoftware/Hydrav3/utils/logger"
)

// Wrapper Function to logger/logger.go
func Logger(handler http.Handler, name string) http.Handler {
	return (logger.Logger(handler, name))
}
