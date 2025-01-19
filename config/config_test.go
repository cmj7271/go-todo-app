package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create new config: %v", err)
	}
	if got.Port != wantPort {
		t.Fatalf("got port %d, want %d", got.Port, wantPort)
	}

	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Fatalf("got env %s, want %s", got.Env, wantEnv)
	}
}
