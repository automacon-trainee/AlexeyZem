package main

import (
	"strconv"
	"testing"
	"time"
)

type testCase struct {
	key  string
	val  any
	want any
}

func TestHashMap(t *testing.T) {
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
		h := NewHashMap(1, WithHashCRC32())
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
		h := NewHashMap(1, WithHashCRC64())
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
}

func TestMeasureTime(t *testing.T) {
	start := time.Now()
	res := MeasureTime(func() {
		time.Sleep(time.Millisecond)
	})
	end := time.Since(start)
	if res > end {
		t.Errorf("MeasureTime() failed: measure time too large")
	}
}

func Benchmark64(b *testing.B) {
	m := NewHashMap(100, WithHashCRC64())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < b.N; i++ {
		m.Get(strconv.Itoa(i))
	}
}

func Benchmark32(b *testing.B) {
	m := NewHashMap(100, WithHashCRC32())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < b.N; i++ {
		m.Get(strconv.Itoa(i))
	}
}
