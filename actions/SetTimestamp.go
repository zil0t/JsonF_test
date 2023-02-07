package actions

import (
	"strings"
	"time"
	JStruct "zz/json"
)

// SetTimestamp Вносим Timestamp в Json
func SetTimestamp() (string, string) {
	tn := time.Now()
	Ct := tn.Format("2006-01-02 15:04:05")
	replacer := strings.NewReplacer("2023-01-26T13:00:00+03:00", Ct)
	Json := replacer.Replace(JStruct.Json)
	return Ct, Json
}
