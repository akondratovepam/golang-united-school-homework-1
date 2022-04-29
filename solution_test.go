package solution

import (
	"testing"

	"github.com/kyokomi/emoji"
)

func TestGetMessage(t *testing.T) {
	expected := emoji.Sprint("Hello, :world_map:!")
	actual := GetMessage()

	if actual != expected {
		t.Fatalf(`GetMessage() = %q, want %q`, actual, expected)
	}
}
