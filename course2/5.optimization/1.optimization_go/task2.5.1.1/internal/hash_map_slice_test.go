package internal

import (
	"testing"
)

type testCase struct {
	key  string
	val  any
	want any
}

func TestHashMapSlice(t *testing.T) {
	tests := []testCase{
		{key: "a", val: "val", want: "val"},
		{key: "b", val: "val", want: "val"},
		{key: "a", val: "sum", want: "sum"},
		{key: "c", val: "some", want: "some"},
		{key: "d", val: "want d", want: "want d"},
		{key: "e", val: "want e", want: "want e"},
		{key: "f", val: "want f", want: "want f"},
	}
	{
		h := NewHashMapSlice(1, WithHashCRC32())
		for _, tc := range tests {
			h.Set(tc.key, tc.val)
			v, ok := h.Get(tc.key)
			if !ok {
				t.Errorf("Get(%s) failed", tc.key)
			}
			if v != tc.want {
				t.Errorf("Get(%s) got %v, want %v", tc.key, v, tc.want)
			}
		}

		val, ok := h.Get("wrong key")
		if ok {
			t.Errorf("Get(%s) should fail", val)
		}
	}

	{
		h := NewHashMapSlice(1, WithHashCRC64())
		for _, tc := range tests {
			h.Set(tc.key, tc.val)
			v, ok := h.Get(tc.key)
			if !ok {
				t.Errorf("Get(%s) failed", tc.key)
			}
			if v != tc.want {
				t.Errorf("Get(%s) got %v, want %v", tc.key, v, tc.want)
			}
		}
		val, ok := h.Get("wrong key")
		if ok {
			t.Errorf("Get(%s) should fail", val)
		}
	}

	{
		h := NewHashMapSlice(1, WithHashCRC32())
		_, ok := h.Get("wrong key")
		if ok {
			t.Errorf("Get should fail")
		}
	}
	{
		h := NewHashMapSlice(1, WithHashCRC64())
		h.Set(tests[0].key, tests[0].val)
		_, ok := h.Get("wrong key")
		if ok {
			t.Errorf("Get should fail")
		}
	}
}
