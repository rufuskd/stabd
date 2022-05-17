package main

import (
  "syscall"
  "encoding/json"
  "net/http"
)

//Start a webserver that can be used to get JSON of resource usage
func main() {
  registerHandlers()
  http.ListenAndServe(":8080", nil)
}

func registerHandlers() {
  http.HandleFunc("/filesystem",filesystem)
}

type filesystem_message struct {
  Type    int64
	Bsize   int64
	Blocks  uint64
	Bfree   uint64
	Bavail  uint64
	Files   uint64
	Ffree   uint64
	Namelen int64
	Frsize  int64
	Flags   int64
	Spare   [4]int64
}

//Make a statfs call, marshall relevant info into json and return it
func filesystem(w http.ResponseWriter, r *http.Request) {
  var buf syscall.Statfs_t
  syscall.Statfs("/dev/sda3", &buf)
  fsinfo := filesystem_message {
    buf.Type,
    buf.Bsize,
  	buf.Blocks,
  	buf.Bfree,
  	buf.Bavail,
  	buf.Files,
  	buf.Ffree,
  	buf.Namelen,
  	buf.Frsize,
  	buf.Flags,
  	buf.Spare}
  resp, _ := json.Marshal(fsinfo)

  w.Write(resp)
}
