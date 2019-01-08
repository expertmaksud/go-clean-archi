package helpers

import (
	"encoding/json"
	"net/http"
)

//HTTPHelper ..
type HTTPHelper struct {
}

// RespondwithJSON write json response format
func (h *HTTPHelper) RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError return error message
func (h *HTTPHelper) RespondWithError(w http.ResponseWriter, code int, msg string) {
	h.RespondwithJSON(w, code, map[string]string{"message": msg})
}
