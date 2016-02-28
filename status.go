package flotilla

import "net/http"
import "errors"

type StatusCode int

func (this StatusCode) Informational() bool { return this >= 100 && this < 200 }
func (this StatusCode) Successful() bool    { return this >= 200 && this < 300 }
func (this StatusCode) Redirection() bool   { return this >= 300 && this < 400 }
func (this StatusCode) BadRequest() bool    { return this >= 400 && this < 500 }
func (this StatusCode) ServerError() bool   { return this >= 500 && this < 600 }

func (this StatusCode) BodyForbidden() bool {
	return this.Informational() ||
		this == http.StatusNoContent ||
		this == http.StatusResetContent ||
		this == http.StatusNotModified
}

func (this StatusCode) Invalid() bool {
	return this < 100 || this >= 600
}

func (this StatusCode) Text() string {
	if text := http.StatusText(this.Number()); text != "" {
		return text
	}

	switch this {
	case 420:
		return "Enhance Your Calm"
	case 451:
		return "Unavailable For Legal Reasons"
	case 522:
		return "Unprocessable Entity"
	default:
		return "Mystery Status Code"
	}
}

func (this StatusCode) Number() int {
	return int(this)
}

func (this StatusCode) Error() error {
	return errors.New(this.Text())
}

const (
	StatusContinue           StatusCode = 100
	StatusSwitchingProtocols StatusCode = 101

	StatusOK                   StatusCode = 200
	StatusCreated              StatusCode = 201
	StatusAccepted             StatusCode = 202
	StatusNonAuthoritativeInfo StatusCode = 203
	StatusNoContent            StatusCode = 204
	StatusResetContent         StatusCode = 205
	StatusPartialContent       StatusCode = 206

	StatusMultipleChoices   StatusCode = 300
	StatusMovedPermanently  StatusCode = 301
	StatusFound             StatusCode = 302
	StatusSeeOther          StatusCode = 303
	StatusNotModified       StatusCode = 304
	StatusUseProxy          StatusCode = 305
	StatusTemporaryRedirect StatusCode = 307

	StatusBadRequest                   StatusCode = 400
	StatusUnauthorized                 StatusCode = 401
	StatusPaymentRequired              StatusCode = 402
	StatusForbidden                    StatusCode = 403
	StatusNotFound                     StatusCode = 404
	StatusMethodNotAllowed             StatusCode = 405
	StatusNotAcceptable                StatusCode = 406
	StatusProxyAuthRequired            StatusCode = 407
	StatusRequestTimeout               StatusCode = 408
	StatusConflict                     StatusCode = 409
	StatusGone                         StatusCode = 410
	StatusLengthRequired               StatusCode = 411
	StatusPreconditionFailed           StatusCode = 412
	StatusRequestEntityTooLarge        StatusCode = 413
	StatusRequestURITooLong            StatusCode = 414
	StatusUnsupportedMediaType         StatusCode = 415
	StatusRequestedRangeNotSatisfiable StatusCode = 416
	StatusExpectationFailed            StatusCode = 417
	StatusTeapot                       StatusCode = 418
	StatusEnhanceYourCalm              StatusCode = 420
	StatusUnavailableForLegalReasons   StatusCode = 451

	StatusInternalServerError     StatusCode = 500
	StatusNotImplemented          StatusCode = 501
	StatusBadGateway              StatusCode = 502
	StatusServiceUnavailable      StatusCode = 503
	StatusGatewayTimeout          StatusCode = 504
	StatusHTTPVersionNotSupported StatusCode = 505
	StatusUnprocessableEntity     StatusCode = 522
)
