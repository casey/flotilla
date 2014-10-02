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

func Status(status int) {
  response_t{"Status", status_t(status), "", ""}.finish()
}

func Body(status int, body, mimetype string) {
  response_t{"Body", status_t(status), body, mimetype}.finish()
}

func Text(status int, body string) {
  response_t{"Text", status_t(status), body, "text/plain; charset=utf-8"}.finish()
}

func HTML(status int, body string) {
  response_t{"HTML", status_t(status), body, "text/html; charset=utf-8"}.finish()
}

func Ensure(condition bool, status int) {
  if !condition {
    response_t{"Ensure", status_t(status), "", ""}.finish()
  }
}

func OK(r *http.Request) {
  response_t{"OK", http.StatusOK, "", ""}.finish()
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
