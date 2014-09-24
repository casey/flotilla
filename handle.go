package flotilla

import "net/http"

func Handle(pattern string) *Handler {
  this := new(Handler)
  this.methods = make(map[string]func(*http.Request))
  this.special = make(map[string]func(*http.Request))
  
  http.HandleFunc(pattern, func (w http.ResponseWriter, r *http.Request) {
    defer Respond(w, r)

    if r.Method == "GET" && r.URL.Path == "/" {
      f, ok := this.special["INDEX"]
      if ok {
        f(r)
        return
      }
    }

    f, ok := this.methods[r.Method]
    if ok {
      f(r)
    } else if f, hasDefault := this.special["DEFAULT"]; hasDefault {
      f(r)
    } else {
      Status(http.StatusMethodNotAllowed)
    }
  })

  return this
}

type Handler struct {
  methods map[string]func(r *http.Request)
  special map[string]func(r *http.Request)
};

func (this *Handler) On(method string, f func(*http.Request)) *Handler {
  _, ok := this.methods[method]
  if ok {
    panic("Handler: duplicate method handler: " + method)
  }
  this.methods[method] = f
  return this
}

func (this *Handler) Special(name string, f func(*http.Request)) *Handler {
  _, ok := this.special[name]
  if ok {
    panic("Handler: duplicate special handler: " + name)
  }
  this.special[name] = f
  return this
}

func (this *Handler) Get    (f func(*http.Request)) *Handler { return this.On("GET",     f) }
func (this *Handler) Put    (f func(*http.Request)) *Handler { return this.On("PUT",     f) }
func (this *Handler) Post   (f func(*http.Request)) *Handler { return this.On("POST",    f) }
func (this *Handler) Options(f func(*http.Request)) *Handler { return this.On("OPTIONS", f) }

func (this *Handler) Default(f func(*http.Request)) *Handler { return this.Special("DEFAULT", f) }
func (this *Handler) Index  (f func(*http.Request)) *Handler { return this.Special("INDEX",   f) }
