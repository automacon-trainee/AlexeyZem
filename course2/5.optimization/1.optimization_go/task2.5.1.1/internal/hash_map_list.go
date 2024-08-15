package internal

type HashMapList struct {
	list   List
	length int
}

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	data Data
	next *Node
}

func (h *HashMapList) Set(key string, value any) {
	node := &Node{data: Data{key: key, value: value}, next: nil}
	if h.length == 0 {
		h.list.head = node
		h.list.tail = h.list.head
	} else {
		cur := h.list.head
		for cur.next != nil {
			if cur.data.key == key {
				cur.data.value = value
				return
			}
			cur = cur.next
		}
		h.list.tail.next = node
		h.list.tail = node
	}
	h.length++
}

func (h *HashMapList) Get(key string) (any, bool) {
	if h.list.head == nil {
		return nil, false
	}
	cur := h.list.head
	for i := 0; i < h.length; i++ {
		if cur.data.key == key {
			return cur.data.value, true
		}
		cur = cur.next
	}
	return nil, false
}

func NewHashMapList() *HashMapList {
	return &HashMapList{list: List{nil, nil}, length: 0}

}
