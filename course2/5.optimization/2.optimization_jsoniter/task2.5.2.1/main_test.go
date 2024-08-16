package main

import (
	"testing"
)

func BenchmarkStandart(b *testing.B) {
	st := &StandartJson{}
	users := GenerateUser(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := st.Marshal(&users[i])
		if err != nil {
			b.Fatal(err)
		}
		_, _ = st.Unmarshal(data)
	}
}

func BenchmarkEasyJson(b *testing.B) {
	st := &EasyJson{}
	users := GenerateUser(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := st.Marshal(&users[i])
		if err != nil {
			b.Fatal(err)
		}
		_, _ = st.Unmarshal(data)
	}
}

func BenchmarkJsointer(b *testing.B) {
	st := &Jsointer{}
	users := GenerateUser(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := st.Marshal(&users[i])
		if err != nil {
			b.Fatal(err)
		}
		_, _ = st.Unmarshal(data)
	}
}
