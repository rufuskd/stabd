package handlers

import(
  "net/http"
  "stabd/pkg/handlers/filesystem"
)

func RegisterHandlers() {
  http.HandleFunc("/filesystem",filesystem.Summary)
}
