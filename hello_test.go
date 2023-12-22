package pkg

import "testing"

func TestReadNumber(t *testing.T) {
	t.Run("", func(t *testing.T) {
		if got := Hello("word"); got != "hello word !" {
			t.Errorf("error")
		}
	})

}
