package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKEY(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey abc123")
	key, err := GetAPIKey(header)
	if err != nil {
		t.Fatal("err")
	}
	if key != "abc123" {
		t.Fatal("wrong key")
	}
}
func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
