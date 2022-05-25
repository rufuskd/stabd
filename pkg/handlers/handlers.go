package handlers

import(
  "net/http"
  "stabd/pkg/handlers/filesystem"
)

func RegisterHandlers() {
  http.HandleFunc("/filesystem",filesystem.Summary)
  http.HandleFunc("/filesystem/inodes",filesystem.Inode)
  
}
