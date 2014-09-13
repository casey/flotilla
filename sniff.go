package flotilla

import "mime"
import "net/http"

func Sniff(extension string, data []byte) string {
  if extension == ".sniff" {
    return http.DetectContentType(data)
  } else if extension != "" { 
    if guess := mime.TypeByExtension(extension); guess != "" {
      return guess
    }
  }
  return "application/octet-stream"
}
