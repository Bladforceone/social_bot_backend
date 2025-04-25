package midleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовки для ВСЕХ запросов (включая OPTIONS)
		headers := w.Header()

		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Methods", "*")
		headers.Set("Access-Control-Allow-Headers", "*")
		headers.Set("Access-Control-Allow-Credentials", "true")
		headers.Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
