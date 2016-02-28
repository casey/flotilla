package flotilla

import "bytes"
import "fmt"
import "errors"
import "appengine"
import "net/http"
import "runtime/debug"

type response_t struct {
	why      string
	status   StatusCode
	body     string
	mimetype string
	to       string
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

func statusResponse(why string, status StatusCode) {
	var response response_t
	response.why = why
	response.status = status
	response.finish()
}

func redirectResponse(why string, status StatusCode, to string) {
	var response response_t
	response.why = why
	response.status = status
	response.to = to
	response.finish()
}

func bodyResponse(why string, status StatusCode, body, mimetype string) {
	var response response_t
	response.why = why
	response.status = status
	response.body = body
	response.mimetype = mimetype
	response.finish()
}

func Die(text string) bool {
	panic(errors.New(text))
	return true
}

func Redirect(status StatusCode, to string) {
	if !status.Redirection() {
		Die(fmt.Sprintf("Redirect received non-redirect status code: %v %v",
			status, status.Text()))
	}
	redirectResponse("Redirect", status, to)
}

func Status(status StatusCode) {
	statusResponse("Status", status)
}

func Body(status StatusCode, body, mimetype string) {
	bodyResponse("Body", status, body, mimetype)
}

func Text(status StatusCode, body string) {
	bodyResponse("Text", status, body, "text/plain; charset=utf-8")
}

func HTML(status StatusCode, body string) {
	bodyResponse("HTML", status, body, "text/html; charset=utf-8")
}

func Ensure(condition bool, status StatusCode) {
	if !condition {
		statusResponse("Ensure", status)
	}
}

func Template(status StatusCode, path string, data interface{}) {
	template := ParseHTMLTemplate(path)

	var b bytes.Buffer
	e := template.Execute(&b, data)

	if e != nil {
		Text(StatusInternalServerError, e.Error())
	}

	HTML(status, b.String())
}

func OK(r *http.Request) {
	statusResponse("OK", http.StatusOK)
}

func Check(e error) {
	if e != nil {
		statusResponse("Check: "+e.Error(), http.StatusInternalServerError)
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
