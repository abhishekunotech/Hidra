package utils

import (
"net/http"
"github.com/Unotechsoftware/Hydra/utils/logger"

)
func Logger(handler http.Handler, name string) http.Handler {
	return (logger.Logger(handler, name))
}
