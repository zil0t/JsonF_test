package actions

import (
	"github.com/tidwall/gjson"
	"testing"
	"zz/values"
)

// Тест для примера
func TestCheckTime(t *testing.T) {
	ct, _ := SetTimestamp()
	g := gjson.Get(values.JsonAdr, "Options.Timestamp")
	if g.Str >= ct {
		t.Fatal("Test false")

	}
}
