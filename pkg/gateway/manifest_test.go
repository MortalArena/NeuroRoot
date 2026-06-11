package gateway

import (
	"testing"
)

func TestNormalizeSitePath(t *testing.T) {
	tests := []struct {
		in      string
		want    string
		wantErr bool
	}{
		{"/", "index.html", false},
		{"style.css", "style.css", false},
		{"../etc/passwd", "", true},
		{"/foo/bar", "foo/bar", false},
	}
	for _, tt := range tests {
		got, err := NormalizeSitePath(tt.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("NormalizeSitePath(%q) err=%v", tt.in, err)
			continue
		}
		if !tt.wantErr && got != tt.want {
			t.Errorf("NormalizeSitePath(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestParseManifestRejectsTraversal(t *testing.T) {
	data := []byte(`{"version":1,"title":"x","files":{"../evil.html":"abc"}}`)
	if _, err := ParseManifest(data); err == nil {
		t.Fatal("should reject traversal path in manifest")
	}
}
