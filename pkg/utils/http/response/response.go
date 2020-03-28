package response

import (
	"github.com/siller174/meetingHelper/pkg/logger"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, message []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeBytes, err := w.Write(message)
	if err != nil {
		logger.Error("Could not write json message in response")
		return err
	}
	if writeBytes == 0 {
		logger.Error("Wrote empty message in response")
	}
	return nil
}

func Empty(w http.ResponseWriter)  {
	w.WriteHeader(http.StatusNoContent)
}
