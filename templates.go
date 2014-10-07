package flotilla

import "appengine"
import "sync"
import "html/template"

var templatecache = make(map[string]*template.Template)
var templatemutex sync.Mutex
var funcmap = make(template.FuncMap)

func init() {
  funcmap["FileURL"] = FileURL
  funcmap["FileURLs"] = FileURLs
}

func ParseHTMLTemplate(path string) *template.Template {
  templatemutex.Lock()
  defer templatemutex.Unlock()
  if template, ok := templatecache[path]; ok {
    return template
  }

  t := template.Must(template.New("").Funcs(funcmap).ParseFiles(path))

  if !appengine.IsDevAppServer() {
    templatecache[path] = t
  }

  return t
}
