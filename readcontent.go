package flotilla

import "net/http"
import "io"

func ReadContent(r *http.Request) ([]byte, error) {
  buffer := make([]byte, r.ContentLength)
  _, e := io.ReadFull(r.Body, buffer)
  r.Body.Close()
  return buffer, e
}
