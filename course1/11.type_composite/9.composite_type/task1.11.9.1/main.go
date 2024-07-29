package main

import (
	"fmt"
)

type TV interface {
	switchOff() bool
	GetStatus() bool
	GetModel() string
	switchOn() bool
}

type LgTv struct {
	status bool
	model  string
}

func (l *LgTv) switchOn() bool {
	l.status = true
	return true
}

func (l *LgTv) switchOff() bool {
	l.status = false
	return true
}

func (l *LgTv) GetStatus() bool {
	return l.status
}
func (l *LgTv) GetModel() string {
	return l.model
}

func (l *LgTv) LgHub() string {
	return "LG Hub"
}

type SamsungTv struct {
	status bool
	model  string
}

func (s *SamsungTv) switchOn() bool {
	s.status = true
	return true
}

func (s *SamsungTv) switchOff() bool {
	s.status = false
	return true
}

func (s *SamsungTv) GetStatus() bool {
	return s.status
}
func (s *SamsungTv) GetModel() string {
	return s.model
}

func (s *SamsungTv) SamsungHub() string {
	return "Samsung Hub"
}

func main() {
	samsung := SamsungTv{status: false, model: "Samsung"}
	Lg := LgTv{status: false, model: "Lg"}
	fmt.Println("Samsung model:")
	fmt.Println(samsung.GetStatus())
	fmt.Println(samsung.GetModel())
	fmt.Println(samsung.SamsungHub())
	fmt.Println(samsung.GetStatus())
	fmt.Println(samsung.switchOn())
	fmt.Println(samsung.GetStatus())
	fmt.Println(samsung.switchOff())
	fmt.Println(samsung.GetStatus())
	fmt.Println("\nLg model:")
	fmt.Println(Lg.GetStatus())
	fmt.Println(Lg.GetModel())
	fmt.Println(Lg.LgHub())
	fmt.Println(Lg.GetStatus())
	fmt.Println(Lg.switchOn())
	fmt.Println(Lg.GetStatus())
	fmt.Println(Lg.switchOff())
	fmt.Println(Lg.GetStatus())
}
