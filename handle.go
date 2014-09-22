package flotilla

import "net/http"

func Handle(pattern string) *Handler {
  this := new(Handler)
  this.methods = make(map[string]func(*http.Request))
  
  http.HandleFunc(pattern, func (w http.ResponseWriter, r *http.Request) {
    defer Respond(w, r)
    f, ok := this.methods[r.Method]
    Ensure(ok, http.StatusMethodNotAllowed)
    f(r)
  })

  return this
}

type Handler struct {
  methods map[string]func(r *http.Request)
};

func (this *Handler) On(method string, f func(*http.Request)) *Handler {
  _, ok := this.methods[method]
  if ok {
    panic("Handler: duplicate method handler: " + method)
  }
  this.methods[method] = f
  return this
}

func (this *Handler) Get    (f func(*http.Request)) *Handler { return this.On("GET",     f) }
func (this *Handler) Put    (f func(*http.Request)) *Handler { return this.On("PUT",     f) }
func (this *Handler) Post   (f func(*http.Request)) *Handler { return this.On("POST",    f) }
func (this *Handler) Options(f func(*http.Request)) *Handler { return this.On("OPTIONS", f) }
