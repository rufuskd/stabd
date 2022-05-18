package main

import (
  "net/http"
  "stabd/pkg/handlers"
)

//Start a webserver that can be used to get JSON of resource usage
func main() {
  handlers.RegisterHandlers()
  http.ListenAndServe(":8080", nil)
}
