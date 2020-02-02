package local

import (
	"testing"
)

func TestClient(t *testing.T) {
	c, err := New("localhost:8080", Insecure())
	if err != nil {
		t.Fatalf("unexpected err: %+v", err)
	}

	status, err := c.Status()
	if err != nil {
		t.Fatalf("unexpected err: %+v", err)
	}

	t.Logf("Status: %+v", status)
}
