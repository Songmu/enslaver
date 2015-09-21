package enslaver

import "testing"

func TestLabor(t *testing.T) {
	slv := Command("dummy")

	exit := slv.labor()
	if exit != exitUnknownErr {
		t.Errorf("exit code should be exitUnknownErr, but: %d", exit)
	}
}
