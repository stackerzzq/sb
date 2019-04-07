package sb

import "testing"

func Test_insert(t *testing.T) {
	m := make(map[string]int)
	m["key1"] = 111111
	m["key2"] = 222222

	s := New()
	s.Insert(m)
}
