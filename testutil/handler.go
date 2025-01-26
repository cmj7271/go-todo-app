package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jg any
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("cannot unmarshal want %q: %v", want, err)
	}
	if err := json.Unmarshal(got, &jg); err != nil {
		t.Fatalf("cannot unmarshal got %q: %v", string(want), err)
	}
	if diff := cmp.Diff(jw, jg); diff != "" {
		t.Fatalf("want and got diff (-want +got):\n%s", diff)
	}
}

func AssertResponse(t *testing.T, want, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() { _ = want.Body.Close() })

	gb, err := io.ReadAll(want.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got.StatusCode != status {
		t.Errorf("want status %d, but got %d, body: %q", status, got.StatusCode, gb)
	}

	if len(gb) == 0 && len(body) == 0 {
		return
	}
	AssertJSON(t, body, gb)
}

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()
	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read %q: %v", path, err)
	}
	return bt
}
