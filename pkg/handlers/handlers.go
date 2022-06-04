package handlers

import(
  "net/http"
  "stabd/pkg/handlers/filesystem"
  "stabd/pkg/handlers/mem"
)

func RegisterHandlers() {
  http.HandleFunc("/filesystem",filesystem.Summary)
  http.HandleFunc("/mem",mem.Summary)
}
