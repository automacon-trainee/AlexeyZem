package main

import (
	"fmt"
)

type BrowseHistory struct {
	stack []string
}

func (h *BrowseHistory) Visit(url string) {
	h.stack = append(h.stack, url)
	fmt.Println("Посещение", url)
}

func (h *BrowseHistory) Back() {
	if len(h.stack) > 0 {
		h.stack = h.stack[:len(h.stack)-1]
		fmt.Println("Возврат к", h.stack[len(h.stack)-1])
	} else {
		fmt.Println("Нет больше истории для возврата")
	}
}

func (h *BrowseHistory) PrintHistory() {
	fmt.Println("История браузера:")
	for i := len(h.stack) - 1; i >= 0; i-- {
		fmt.Println(h.stack[i])
	}
}

func main() {
	browseHistory := &BrowseHistory{}
	browseHistory.Visit("http://www.baidu.com")
	browseHistory.Visit("http://google.com")
	browseHistory.Visit("https://github.com")
	browseHistory.Back()
	browseHistory.PrintHistory()
}
