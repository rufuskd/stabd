package filesystem

import(
  "os"
  "bufio"
  "unicode"
  "strings"
  "net/http"
  "syscall"
  "encoding/json"
  "log"
)

type filesystem_stats struct {
  Name    string
  Type    int64
	Bsize   int64
	Blocks  uint64
	Bfree   uint64
	Bavail  uint64
	Files   uint64
	Ffree   uint64
  Fsid    syscall.Fsid
	Namelen int64
	Frsize  int64
	Flags   int64
	Spare   [4]int64
}

type filesystem_message struct {
  FS  []filesystem_stats
}

//Return a list of all filesystems mounted
func fslist() ([]string, error) {
  var retval []string
  mounts, err := os.Open("/proc/mounts")
  if err != nil {
    log.Print("Unable to read /proc/mounts, stabd might not be running with" +
      "sufficient permissions even though it doesn't need root")
    return nil, err
  } else {
    defer mounts.Close()
    scanner := bufio.NewScanner(mounts)
    for scanner.Scan() {
      currentMount := strings.FieldsFunc(scanner.Text(),unicode.IsSpace)[0]
    retval = append(retval, currentMount)
    }
    return retval, nil
  }
}

//Make a statfs call, marshall relevant info into json and return it
func Summary(w http.ResponseWriter, r *http.Request) {
  mounts, err := fslist()
  if err != nil {
    log.Print("Unable to read /proc/mounts, stabd might not be running with" +
      "sufficient permissions even though it doesn't need root")
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Unable to read /proc/mounts"))
  } else {
    var retVal filesystem_message
    var buf syscall.Statfs_t

    for _, currentMount := range mounts {
      err := syscall.Statfs(currentMount, &buf)
      if err != nil {
        log.Print("Unable to read filesystem stats for " + currentMount + " moving on")
      } else {
        fsinfo := filesystem_stats {
          currentMount,
          buf.Type,
          buf.Bsize,
        	buf.Blocks,
        	buf.Bfree,
        	buf.Bavail,
        	buf.Files,
        	buf.Ffree,
          buf.Fsid,
        	buf.Namelen,
        	buf.Frsize,
        	buf.Flags,
        	buf.Spare}
        retVal.FS = append(retVal.FS,fsinfo)
      }
    }

    resp, _ := json.Marshal(retVal)
    w.Write(resp)
  }
}
