package flotilla

import "net/http"

type status_t int

func (this status_t) informational() bool { return this >= 100 && this < 200 }
func (this status_t) successful() bool    { return this >= 200 && this < 300 }
func (this status_t) redirection() bool   { return this >= 300 && this < 400 }
func (this status_t) badRequest() bool    { return this >= 400 && this < 500 }
func (this status_t) serverError() bool   { return this >= 500 && this < 600 }

func (this status_t) bodyForbidden() bool {
  return this.informational() ||
    this == http.StatusNoContent ||
    this == http.StatusResetContent ||
    this == http.StatusNotModified;
}

func (this status_t) invalid() bool {
  return this < 100 || this >= 600
}

func (this status_t) text() string {
  if text := http.StatusText(this.number()); text != "" {
    return text
  }

  switch this {
  case 420: return "Enhance Your Calm"
  case 451: return "Unavailable For Legal Reasons"
  case 522: return "Unprocessable Entity"
  default:  return "Mystery Status Code"
  }
}

func (this status_t) number() int {
  return int(this)
}
