package stringzeug_test

import (
	"github.com/plombardi89/gozeug/stringzeug"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIsNotEmpty(t *testing.T) {
	assert.True(t, stringzeug.IsNotEmpty([]string{"a"}))
	assert.False(t, stringzeug.IsNotEmpty([]string{}))
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, stringzeug.IsEmpty([]string{}))
	assert.False(t, stringzeug.IsEmpty([]string{"a"}))
}

func TestRemoveAtIndex(t *testing.T) {
	table := []struct {
		in        []string
		index     int
		out       []string
		succeeded bool
	}{
		// empty slice
		{[]string{}, 5, []string{}, false},

		// negative index
		{[]string{}, -10, []string{}, false},

		// index greater than slice size
		{[]string{"a", "b", "c"}, 5, []string{"a", "b", "c"}, false},

		{[]string{"a", "b", "c"}, 1, []string{"a", "c"}, true},
	}

	for _, r := range table {
		out, res := stringzeug.RemoveAtIndex(r.in, r.index)
		assert.Equal(t, r.out, out)
		assert.Equal(t, r.succeeded, res)
	}
}

func TestRemoveIf(t *testing.T) {
	table := []struct {
		in        []string
		predicate func(string) bool
		out       []string
	}{
		{[]string{"baz", "bar", "foo"}, func(s string) bool { return strings.HasPrefix(s, "ba") }, []string{"foo"}},
		{[]string{"baz", "bar", "foo"}, func(s string) bool { return false }, []string{"baz", "bar", "foo"}},
	}

	for _, r := range table {
		out := stringzeug.RemoveIf(r.in, r.predicate)
		assert.Equal(t, r.out, out)
	}
}
