package mime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMimetype(t *testing.T) {
	t.Run("to-mime", func(t *testing.T) {
		t.Run("mime-json", func(t *testing.T) {
			mt := ToMimetype(MimeApplicationJSON)
			assert.Equal(t, mt, mimeApplicationJSON)
		})
	})
}
