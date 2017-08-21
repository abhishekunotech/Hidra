package utils

import (
"net/http"
"github.com/UnotechSoftware/felicitymiddleware/v1/utils/logger"

)
func Logger(handler http.Handler, name string) http.Handler {
	return (logger.Logger(handler, name))
}