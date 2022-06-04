package mem

import(
  "os"
  "bufio"
  "unicode"
  "strings"
  "net/http"
  "encoding/json"
  "log"
)

//Read /proc/meminfo stash it in a map and return it
func Summary(w http.ResponseWriter, r *http.Request) {
  memInfoMap := make(map[string][]string)

  memInfo, err := os.Open("/proc/meminfo")
  if(err != nil) {
    log.Print("Unable to open /proc/meminfo")
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Unable to read /proc/mounts"))
  } else {
    defer  memInfo.Close()
    scanner := bufio.NewScanner(memInfo)
    for scanner.Scan() {
      currentField := strings.FieldsFunc(scanner.Text(),unicode.IsSpace)
      currentField[0] = strings.Trim(currentField[0],":")
      memInfoMap[currentField[0]] = currentField[1:]
    }
  }

  resp, _ := json.Marshal(memInfoMap)
  w.Write(resp)
}
