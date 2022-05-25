package filesystem

import(
  "fmt"
  "os"
  "bufio"
  "unicode"
  "strings"
  "net/http"
  "syscall"
  "encoding/json"
  "log"
)

type filesystem_message struct {
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

//Make a statfs call, marshall relevant info into json and return it
func Summary(w http.ResponseWriter, r *http.Request) {
  mounts, err := os.Open("/proc/mounts")
  if err != nil {
    log.Print("Unable to read /proc/mounts, stabd might not be running with" +
      "sufficient permissions even though it doesn't need root")
  } else {
    defer mounts.Close()
    scanner := bufio.NewScanner(mounts)
    for scanner.Scan() {
      fmt.Println(strings.FieldsFunc(scanner.Text(),unicode.IsSpace)[0] + "kupo")
    }
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
      buf.Fsid,
    	buf.Namelen,
    	buf.Frsize,
    	buf.Flags,
    	buf.Spare}
    resp, _ := json.Marshal(fsinfo)

    w.Write(resp)
  }
}

func Inode(w http.ResponseWriter, r *http.Request) {
  var buf syscall.Statfs_t
  syscall.Statfs("/dev/sda3", &buf)
  inodeMap := map[string]uint64{"InodesAvail" : buf.Ffree, "InodesUsed" : buf.Files}
  resp, _ := json.Marshal(inodeMap)

  w.Write(resp)
}
