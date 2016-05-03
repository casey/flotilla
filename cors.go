package flotilla

import "net/http"

func UsePermissiveCORS(w http.ResponseWriter, r *http.Request) {
	if value := r.Header.Get("Access-Control-Request-Headers"); value != "" {
		w.Header().Set("Access-Control-Allow-Headers", value)
	}

	if r.Header.Get("Access-Control-Request-Method") != "" {
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, CONNECT")
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "1728000")
}
