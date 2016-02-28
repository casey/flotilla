package flotilla

import "sync"
import "io/ioutil"
import "net/http"

type Result struct {
	Response *http.Response
	Status   StatusCode
	Bytes    []byte
	OK       bool
	error    error
	string   string
}

func (this *Result) String() string {
	if len(this.Bytes) == 0 {
		return ""
	}

	if this.string == "" {
		this.string = string(this.Bytes)
	}

	return this.string
}

func (this *Result) Error() error {
	if this.error != nil {
		return this.error
	}

	if !this.Status.Successful() {
		return this.Status.Error()
	}

	return nil
}

type requestGroup_t struct {
	WaitGroup sync.WaitGroup
	Mutex     sync.Mutex
	Client    *http.Client
	callbacks []func()
}

func RequestGroup(client *http.Client) (this requestGroup_t) {
	this.Client = client
	return
}

func (this *requestGroup_t) Do(request *http.Request, callback func(Result)) *requestGroup_t {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()

	this.WaitGroup.Add(1)
	var result Result
	this.callbacks = append(this.callbacks, func() { callback(result) })

	go func() {
		defer this.WaitGroup.Done()
		response, e := this.Client.Do(request)
		defer response.Body.Close()
		result.Response = response
		result.error = e

		if e != nil {
			return
		}

		result.Status = StatusCode(response.StatusCode)
		bytes, e := ioutil.ReadAll(response.Body)
		result.Bytes = bytes
		result.error = e

		result.OK = result.error == nil && result.Status.Successful()
	}()

	return this
}

func (this *requestGroup_t) Wait() *requestGroup_t {
	this.WaitGroup.Wait()
	for _, callback := range this.callbacks {
		callback()
	}
	this.callbacks = nil
	return this
}

/*
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
*/
