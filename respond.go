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

  if value := r.Header.Get("Access-Control-Request-Headers"); value != "" {
    w.Header().Set("Access-Control-Allow-Headers", value)
  }

  if r.Header.Get("Access-Control-Request-Method") != "" {
    w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, CONNECT")
  }

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("License", "Anyone may do anything with this.")
  w.Header().Set("Warranty", `"AS IS" WITH NO WARRANTY OF ANY KIND EXPRESS OR IMPLIED.`)
  w.Header().Set("Access-Control-Max-Age", "1728000")
  w.Header().Set("Content-Type", response.mimetype)
  w.WriteHeader(response.status.Number())
  w.Write([]byte(response.body))
}
