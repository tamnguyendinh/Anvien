package version

import "testing"

func TestVersionMetadataIsPresent(t *testing.T) {
	if CommandName != "anvien" {
		t.Fatalf("CommandName = %q", CommandName)
	}
	if Version == "" {
		t.Fatal("Version must not be empty")
	}
}
