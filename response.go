package flotilla

import "fmt"
import "appengine"
import "net/http"
import "runtime/debug"

type response_t struct {
  why      string
  status   status_t
  body     string
  mimetype string
}

var debugEnabled = false

func Debug(enabled bool) {
  debugEnabled = enabled
}

func (this response_t) finish() {
  if appengine.IsDevAppServer() && debugEnabled {
    debug.PrintStack()
  }
  panic(this)
}

func Status(status status_t) {
  response_t{"Status", status, "", ""}.finish()
}

func Body(status status_t, body, mimetype string) {
  response_t{"Body", status, body, mimetype}.finish()
}

func Ensure(condition bool, status status_t) {
  if !condition {
    response_t{"Ensure", status, "", ""}.finish()
  }
}

func Check(e error) {
  if e != nil {
    response_t{"Check: " + e.Error(), http.StatusInternalServerError, "", ""}.finish()
  }
}

func (this *response_t) repair() {
  if this.why == "" {
    this.why = "Unknown"
  }

  if this.status == 0 {
    this.status = http.StatusInternalServerError
  }

  if this.body == "" && this.mimetype == "" {
    if this.status.bodyForbidden() {
      this.body = "\n"
    } else {
      this.body = fmt.Sprintf("%v %v\n", this.status.number(), this.status.text())
    }
  }

  if this.mimetype == "" {
    this.mimetype = http.DetectContentType([]byte(this.body))
  }
}
