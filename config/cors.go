package config

import "net/http"

//EnableCors function
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Origin, x-api-key,Content-Type, X-Auth-Token, Authorization")
	(*w).Header().Set("Content-type", "application/json")
}

//EnableCorsOptions function
func EnableCorsOptions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Origin, x-api-key,Content-Type, X-Auth-Token, Authorization")
			w.Header().Set("Content-type", "application/json")
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Origin, x-api-key,Content-Type, X-Auth-Token, Authorization")
		w.Header().Set("Content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}
