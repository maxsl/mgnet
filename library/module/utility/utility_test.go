package utility

import (
	"testing"
)

func Test_All(t *testing.T) {
	id := GenSessionId(10, 20, "hello")

	if len(id) != 40 {
		t.Errorf("ERROR: GenSessionId len error")
	}
	
	
	t.Logf("SessionId : %s", id)
}