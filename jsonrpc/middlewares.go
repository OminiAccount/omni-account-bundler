package jsonrpc

import (
	"bytes"
	"github.com/OAB/utils/log"
	"io"
	"net/http"
)

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if r.Method == http.MethodOptions {
		// 	return
		// }
		// if r.Method != http.MethodPost {
		// 	http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		// 	return
		// }
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// w.Header().Set("Content-type", "application/json")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Only allow POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
			return
		}

		// Set content type for POST requests
		w.Header().Set("Content-type", "application/json")

		body := io.LimitReader(r.Body, 1024*1024*5)
		data, err := io.ReadAll(body)
		if err != nil {
			http.Error(w, "read data error", http.StatusBadRequest)
			return
		}
		/*var req types.Request
		if err := json.Unmarshal(data, &req); err != nil {
			http.Error(w, "invalid json object request body", http.StatusBadRequest)
			return
		}*/
		log.Infof("%s %s", r.Method, string(data))
		r.Body = io.NopCloser(bytes.NewBuffer(data))
		// Serve the request
		h.ServeHTTP(w, r)

	})
}
