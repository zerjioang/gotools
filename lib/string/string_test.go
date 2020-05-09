package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("create-empty", func(t *testing.T) {
		empty := New()
		assert.Equal(t, empty, String{})
	})
	t.Run("create-with-data", func(t *testing.T) {
		empty := NewWith([]byte("a"))
		assert.NotNil(t, empty)
	})
	t.Run("char-at", func(t *testing.T) {
		empty := NewWith([]byte("foo-bar"))
		c := empty.CharAt(0)
		assert.Equal(t, c, []byte("f")[0])
	})
	t.Run("last-index-of", func(t *testing.T) {
		example := []byte("foo-bar-hello-world")
		search := []byte("h")
		exampleString := NewWith(example)
		location := strings.LastIndex("foo-bar-hello-world", "h")
		position := exampleString.LastIndex(search)
		assert.Equal(t, position, location)
	})
	t.Run("is-empty", func(t *testing.T) {
		example := []byte("foo-bar-hello-world")
		exampleString := NewWith(example)
		eval := exampleString.IsEmpty()
		assert.Equal(t, eval, false)
	})
	t.Run("to-lower-case", func(t *testing.T) {
		example := []byte("Foo-bar-HELLO-world")
		exampleString := NewWith(example)
		exampleString.LowerCase()
		assert.Equal(t, exampleString.String(), "foo-bar-hello-world")
	})
	t.Run("to-capitalize-case", func(t *testing.T) {
		example := []byte("foo-bar")
		exampleString := NewWith(example)
		exampleString.Capitalize()
		assert.Equal(t, exampleString.String(), "Foo-bar")
	})
	t.Run("to-upper-case", func(t *testing.T) {
		example := []byte("Foo-bar-HELLO-world")
		exampleString := NewWith(example)
		exampleString.UpperCase()
		assert.Equal(t, exampleString.String(), "FOO-BAR-HELLO-WORLD")
	})
	t.Run("to-title-case", func(t *testing.T) {
		example := []byte("FOO-BAR-HELLO-WORLD")
		exampleString := NewWith(example)
		exampleString.TitleCase()
		assert.Equal(t, exampleString.String(), "Foo-Bar-Hello-World")
	})
	t.Run("reverse", func(t *testing.T) {
		example := []byte("Foo-bar-HELLO-world")
		exampleString := NewWith(example)
		exampleString.Reverse()
		assert.Equal(t, exampleString.String(), "dlrow-OLLEH-rab-ooF")
	})
	t.Run("count-match-byte", func(t *testing.T) {
		example := []byte("Foo-bar-HELLO-world")
		exampleString := NewWith(example)
		matches := exampleString.CountByte([]byte("o")[0])
		assert.Equal(t, matches, 3)
	})
	t.Run("contains", func(t *testing.T) {
		example := []byte("Foo-bar-HELLO-world")
		exampleString := NewWith(example)
		result := exampleString.Contains([]byte("world"))
		assert.True(t, result)
	})
	t.Run("compare", func(t *testing.T) {
		example := []byte("Foo-bar-HELLO-world")
		exampleString := NewWith(example)
		result := exampleString.Compare(exampleString)
		assert.Equal(t, result, 0)
	})
	t.Run("has-suffix", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			example := []byte("Foo-bar-HELLO-world")
			exampleString := NewWith(example)
			result := exampleString.HasSuffix("world")
			assert.True(t, result)
		})
		t.Run("fail", func(t *testing.T) {
			example := []byte("Foo-bar-HELLO-world")
			exampleString := NewWith(example)
			result := exampleString.HasSuffix("blablabla")
			assert.False(t, result)
		})
	})
	t.Run("has-prefix", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			example := []byte("Foo-bar-HELLO-world")
			exampleString := NewWith(example)
			result := exampleString.HasPrefix([]byte("Foo"))
			assert.True(t, result)
		})
		t.Run("fail", func(t *testing.T) {
			example := []byte("Foo-bar-HELLO-world")
			exampleString := NewWith(example)
			result := exampleString.HasPrefix([]byte("blablabla"))
			assert.False(t, result)
		})
	})
	t.Run("is-numeric", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			example := []byte("123456")
			exampleString := NewWith(example)
			result := exampleString.IsNumeric()
			assert.True(t, result)
		})
		t.Run("fail", func(t *testing.T) {
			example := []byte("123456a")
			exampleString := NewWith(example)
			result := exampleString.IsNumeric()
			assert.False(t, result)
		})
	})
	t.Run("is-hexadecimal", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			example := []byte("d46d1326aed64ac499cc02a128339b99")
			exampleString := NewWith(example)
			result := exampleString.IsHexadecimal()
			assert.True(t, result)
		})
		t.Run("fail", func(t *testing.T) {
			example := []byte("foo-bar")
			exampleString := NewWith(example)
			result := exampleString.IsHexadecimal()
			assert.False(t, result)
		})
	})
}
