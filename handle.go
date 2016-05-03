package flotilla

import "net/http"

func Handle(pattern string) *Handler {
	this := new(Handler)
	this.methods = make(map[string]http.HandlerFunc)
	this.special = make(map[string]http.HandlerFunc)

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		defer Respond(w, r)

		if r.Method == "GET" && r.URL.Path == "/" {
			f, ok := this.special["INDEX"]
			if ok {
				f(w, r)
				return
			}
		}

		f, ok := this.methods[r.Method]
		if ok {
			f(w, r)
		} else if f, hasDefault := this.special["DEFAULT"]; hasDefault {
			f(w, r)
		} else {
			Status(http.StatusMethodNotAllowed)
		}
	})

	return this
}

type Handler struct {
	methods map[string]http.HandlerFunc
	special map[string]http.HandlerFunc
}

func (this *Handler) On(method string, f http.HandlerFunc) *Handler {
	_, ok := this.methods[method]
	if ok {
		panic("Handler: duplicate method handler: " + method)
	}
	this.methods[method] = f
	return this
}

func (this *Handler) Special(name string, f http.HandlerFunc) *Handler {
	_, ok := this.special[name]
	if ok {
		panic("Handler: duplicate special handler: " + name)
	}
	this.special[name] = f
	return this
}

func (this *Handler) Get(f http.HandlerFunc) *Handler     { return this.On("GET", f) }
func (this *Handler) Put(f http.HandlerFunc) *Handler     { return this.On("PUT", f) }
func (this *Handler) Post(f http.HandlerFunc) *Handler    { return this.On("POST", f) }
func (this *Handler) Options(f http.HandlerFunc) *Handler { return this.On("OPTIONS", f) }

func (this *Handler) Default(f http.HandlerFunc) *Handler { return this.Special("DEFAULT", f) }
func (this *Handler) Index(f http.HandlerFunc) *Handler   { return this.Special("INDEX", f) }
