package utility

import (
	"testing"
)

func Test_All(t *testing.T) {
	id := GenSessionId()

	if len(id) != 40 {
		t.Errorf("ERROR: GenSessionId len error")
	}
	
	t.Logf("%s", id)
}