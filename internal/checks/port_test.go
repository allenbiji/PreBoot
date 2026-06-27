package checks_test

import (
	"net"
	"strconv"
	"strings"
	"testing"

	"github.com/allenbiji/clone-sage/internal/checks"
	"github.com/allenbiji/clone-sage/internal/model"
	"github.com/allenbiji/clone-sage/internal/registry"
)

func TestBuildPortFreeCheck(t *testing.T) {
	tests := []struct {
		name    string
		opts    map[string]string
		wantErr string
	}{
		{"no options", nil, "requires a 'port' option"},
		{"empty port", map[string]string{"port": ""}, "requires a 'port' option"},
		{"valid port", map[string]string{"port": "9000"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := registry.Build(cfg(model.TypePortFree, tt.opts))
			if (err != nil) != (tt.wantErr != "") {
				t.Fatalf("wantErr=%q got=%v", tt.wantErr, err)
			}
			if err != nil && !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error %q does not contain %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestPortFreeCheck_Execute(t *testing.T) {
	t.Run("port in use", func(t *testing.T) {
		t.Parallel()
		l, err := net.Listen("tcp", ":0")
		if err != nil {
			t.Fatal(err)
		}
		defer l.Close()
		port := strconv.Itoa(l.Addr().(*net.TCPAddr).Port)
		check := &checks.PortFreeCheck{Port: port}
		err = check.Execute()
		if err == nil {
			t.Fatal("expected error for port in use, got nil")
		}
		if !strings.Contains(err.Error(), "not free") {
			t.Errorf("error %q does not contain 'not free'", err.Error())
		}
	})

	t.Run("port free", func(t *testing.T) {
		t.Parallel()
		// Bind on :0 to get an ephemeral port, then release it.
		l, err := net.Listen("tcp", ":0")
		if err != nil {
			t.Fatal(err)
		}
		port := strconv.Itoa(l.Addr().(*net.TCPAddr).Port)
		l.Close()

		check := &checks.PortFreeCheck{Port: port}
		if err := check.Execute(); err != nil {
			t.Errorf("expected nil for free port, got: %v", err)
		}
	})
}
