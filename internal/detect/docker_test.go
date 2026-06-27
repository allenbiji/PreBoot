package detect

import "testing"

func TestExtractHostPort(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"host:container", "8080:80", "8080"},
		{"ip:host:container", "127.0.0.1:8080:80", "8080"},
		{"bare port no colon", "80", ""},
		{"empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := extractHostPort(tt.input)
			if got != tt.want {
				t.Errorf("extractHostPort(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
