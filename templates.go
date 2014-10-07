package flotilla

import "appengine"
import "sync"
import "html/template"
import "io/ioutil"
import "path/filepath"

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

  bytes, e := ioutil.ReadFile(path)
  
  if e != nil {
    panic(e)
  }

  t := template.Must(template.New(filepath.Base(path)).Funcs(funcmap).Parse(string(bytes)))

  if !appengine.IsDevAppServer() {
    templatecache[path] = t
  }

  return t
}
