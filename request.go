package flotilla

import "appengine"
import "appengine/urlfetch"
import "strings"
import "io/ioutil"
import "net/http"

func Request(c appengine.Context, method, url, body string) (StatusCode, string, error) {
  client := urlfetch.Client(c)
  request, e := http.NewRequest(method, url, strings.NewReader(body))
  if e != nil {
    return 0, "", e
  } 
  response, e := client.Do(request)
  if e != nil {
    return 0, "", e
  }
  payload, e := ioutil.ReadAll(response.Body)
  response.Body.Close()
  if e != nil {
    return 0, "", e
  }

  status := StatusCode(response.StatusCode)
  result := string(payload)
  if status.Successful() {
    return status, result, nil
  } else {
    return status, result, status.Error()
  }
}

func Get(c appengine.Context, url string) (StatusCode, string, error) {
  return Request(c, "GET",  url, "")
}

func Put(c appengine.Context, url string, body string) (StatusCode, string, error) {
  return Request(c, "PUT",  url, body)
}

func Post(c appengine.Context, url string, body string) (StatusCode, string, error) {
  return Request(c, "POST", url, body)
}
