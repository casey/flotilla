package flotilla

import "appengine"
import "crypto/sha256"
import "encoding/hex"
import "io/ioutil"
import "path/filepath"
import "sync"
import "html/template"

var hashcache = make(map[string]string)
var hashmutex sync.Mutex

func hash(path string) string {
  hashmutex.Lock()
  defer hashmutex.Unlock()

  if !appengine.IsDevAppServer() {
    if cached, ok := hashcache[path]; ok {
      return cached
    }
  }

  b, e := ioutil.ReadFile(path)
  
  if e != nil {
    panic(e)
  }

  sha := sha256.New()
  sha.Write(b)
  sum := sha.Sum(nil)
  hash := hex.EncodeToString(sum)

  hashcache[path] = hash

  return hash
}

func FileURL(path string) template.URL {
  return template.URL(filepath.Join(hash(path)[:16], path))
}

func FileURLs(globs ...string) []template.URL {
  links := make([]template.URL, 0)

  for _, glob := range globs {
    paths, _ := filepath.Glob(glob)
    for _, path := range(paths) {
      links = append(links, FileURL(path))
    }
  }
  
  return links
}
