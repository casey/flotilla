package flotilla

import "appengine"
import "net/http"

func Respond(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	relic := recover()
	response, ok := relic.(response_t)

	if !ok {
		c.Errorf("handler: relic was not response: %v", relic)
	}

	response.repair()

	if r.Method == "HEAD" && response.status == 200 {
		response.body = ""
	}

	c.Infof("handler: why: %v", response.why)

	w.Header().Set("Content-Type", response.mimetype)

	if response.to != "" {
		http.Redirect(w, r, response.to, response.status.Number())
	} else {
		w.WriteHeader(response.status.Number())
		w.Write([]byte(response.body))
	}
}
