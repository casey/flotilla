package flotilla

import "appengine"
import "crypto/sha256"
import "encoding/hex"
import "io/ioutil"
import "path/filepath"

var hashcache = make(map[string]string)

func hash(path string) string {
  b, e := ioutil.ReadFile(path)
  
  if e != nil {
    panic(e)
  }

  sha := sha256.New()
  sha.Write(b)
  sum := sha.Sum(nil)
  hash := hex.EncodeToString(sum)

  if !appengine.IsDevAppServer() {
    hashcache[path] = hash
  }

  return hash
}

func Link(path string) string {
  return filepath.Join(hash(path)[:16], path)
}

func Links(paths ...string) []string {
  links := make([]string, len(paths))

  for i, path := range paths {
    links[i] = Link(path)
  }
  
  return links
}