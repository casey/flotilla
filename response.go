package flotilla

import "fmt"
import "appengine"
import "net/http"
import "runtime/debug"

type response_t struct {
  why      string
  status   StatusCode
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

func Status(status StatusCode) {
  response_t{"Status", status, "", ""}.finish()
}

func Body(status StatusCode, body, mimetype string) {
  response_t{"Body", status, body, mimetype}.finish()
}

func Text(status StatusCode, body string) {
  response_t{"Text", status, body, "text/plain; charset=utf-8"}.finish()
}

func HTML(status StatusCode, body string) {
  response_t{"HTML", status, body, "text/html; charset=utf-8"}.finish()
}

func Ensure(condition bool, status StatusCode) {
  if !condition {
    response_t{"Ensure", status, "", ""}.finish()
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
    if this.status.BodyForbidden() {
      this.body = "\n"
    } else {
      this.body = fmt.Sprintf("%v %v\n", this.status.Number(), this.status.Text())
    }
  }

  if this.mimetype == "" {
    this.mimetype = http.DetectContentType([]byte(this.body))
  }
}
