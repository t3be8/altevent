package comment

import (
	"altevent/delivery/middlewares"
	"testing"
)

var (
	token string
)

func TestUseTokenizer(t *testing.T) {
	t.Run("Set Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(99, "Nobar EUFA Champions", "Nobar final liga champions")
	})
}
