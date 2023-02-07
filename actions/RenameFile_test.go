package actions

import (
	"testing"
)

func TestRename(t *testing.T) {
	Input.string = "Options.Change_Name.Condition"
	Input.value = 1
	if Input.string != "Options.Change_Name.Condition" || Input.value != 1 {
		t.Fatal("invalid arguments")
	}
}
