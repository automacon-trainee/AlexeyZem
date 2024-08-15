package main

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		if d.length != 0 || d.curr != nil && d.head != nil || d.tail != nil {
			t.Errorf("Init fail")
		}
	}

	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		cur := d.head
		for _, c := range commits {
			if *cur.data != c {
				t.Errorf("Init fail")
			}
			cur = cur.next
		}
	}
}

func TestLen(t *testing.T) {
	d := DoubleLinkedList{}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		if d.Len() != 4 {
			t.Errorf("Len fail")
		}
	}
	{
		commits := []Commit{}
		d.Init(commits)
		if d.Len() != 0 {
			t.Errorf("Len fail")
		}
	}
}

func TestCurrent(t *testing.T) {
	d := DoubleLinkedList{}
	{
		commits := []Commit{}
		d.Init(commits)
		if d.Current() != nil {
			t.Errorf("Current fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		cur := d.head.next.next
		d.curr = cur
		if d.Current() != cur {
			t.Errorf("Current fail")
		}
	}
}

func TestSetCurrent(t *testing.T) {
	d := DoubleLinkedList{}
	{
		commits := []Commit{}
		d.Init(commits)
		err := d.SetCurrent(10)
		if err == nil {
			t.Errorf("SetCurrent fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		cur := d.head.next.next
		err := d.SetCurrent(2)
		if err != nil {
			t.Errorf("SetCurrent fail")
		}
		if d.Current() != cur {
			t.Errorf("Set Current fail want %v, got %v", *cur.data, *d.Current().data)
		}
	}
}

func TestNext(t *testing.T) {
	d := DoubleLinkedList{}
	{
		node := d.Next()
		if node != nil {
			t.Errorf("Next fail want nil, got %v", node)
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		_ = d.SetCurrent(1)
		if d.Next() != d.Current().next {
			t.Errorf("Next fail want %v, got %v", d.Current().next, d.Next())
		}
	}
}

func TestPrev(t *testing.T) {
	d := DoubleLinkedList{}
	{
		node := d.Prev()
		if node != nil {
			t.Errorf("Prev fail want nil, got %v", node)
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		_ = d.SetCurrent(1)
		if d.Prev() != d.Current().prev {
			t.Errorf("Prev fail want %v, got %v", d.Current().prev, d.Prev())
		}
	}
}

func TestInsert(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		err := d.Insert(10, Commit{UUID: "123", Message: "first commit", Date: "2 november"})
		if err == nil {
			t.Errorf("Insert fail")
		}
	}
	{
		d.Init([]Commit{})
		err := d.Insert(0, Commit{UUID: "123", Message: "first commit", Date: "2 november"})
		if err != nil {
			t.Errorf("Insert fail")
		}
		if !reflect.DeepEqual(*d.head.data, Commit{UUID: "123", Message: "first commit", Date: "2 november"}) {
			t.Errorf("Insert first commit in empty fail")
		}
	}
	{
		d.Init([]Commit{
			{UUID: "3", Message: "first commit", Date: "3 november"},
		})
		err := d.Insert(0, Commit{UUID: "123", Message: "first commit", Date: "2 november"})
		if err != nil {
			t.Errorf("Insert fail")
		}
		if !reflect.DeepEqual(*d.head.data, Commit{UUID: "123", Message: "first commit", Date: "2 november"}) {
			t.Errorf("Insert first commit in non empty fail ")
		}
	}
	{
		d.Init([]Commit{
			{UUID: "3", Message: "first commit", Date: "3 november"},
		})
		err := d.Insert(1, Commit{UUID: "123", Message: "first commit", Date: "2 november"})
		if err != nil {
			t.Errorf("Insert fail")
		}
		if !reflect.DeepEqual(*d.tail.data, Commit{UUID: "123", Message: "first commit", Date: "2 november"}) {
			t.Errorf("Insert last commit fail")
		}
	}

	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		err := d.Insert(2, Commit{UUID: "123456", Message: "some commit", Date: "20 november"})
		if err != nil {
			t.Errorf("Insert fail")
		}
		if !reflect.DeepEqual(*d.head.next.next.data, Commit{UUID: "123456", Message: "some commit", Date: "20 november"}) {
			t.Errorf("Insert commit fail. want %v, got %v", *d.head.next.next.data, Commit{UUID: "123456", Message: "some commit", Date: "20 november"})
		}
	}
}

func TestPush(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		d.Push(Commit{UUID: "123456", Message: "some commit", Date: "20 november"})
		if !reflect.DeepEqual(*d.head.data, Commit{UUID: "123456", Message: "some commit", Date: "20 november"}) {
			t.Errorf("Push fail")
		}
	}
	{
		d.Init([]Commit{
			{UUID: "123456", Message: "some commit", Date: "20 november"},
		})
		d.Push(Commit{UUID: "123", Message: "new commit", Date: "2 november"})
		if !reflect.DeepEqual(*d.tail.data, Commit{UUID: "123", Message: "new commit", Date: "2 november"}) {
			t.Errorf("Push fail")
		}
	}
}

func TestDelete(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		err := d.Delete(10)
		if err == nil {
			t.Errorf("Delete fail")
		}
	}
	{
		d.Init([]Commit{
			{UUID: "123", Message: "new commit", Date: "2 november"},
		})
		err := d.Delete(0)
		if err != nil {
			t.Errorf("Delete fail(error: %v)", err)
		}
		if d.head != nil {
			t.Errorf("Delete fail(wrong new head)")
		}
		if d.length != 0 {
			t.Errorf("Delete fail(wrong len)")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		err := d.Delete(0)
		if err != nil {
			t.Errorf("Delete fail(error: %v)", err)
		}
		if !reflect.DeepEqual(*d.head.data, Commit{UUID: "456", Message: "second commit", Date: "3 november"}) {
			t.Errorf("Delete fail")
		}
		if d.length != 3 {
			t.Errorf("Delete fail(wrong len)")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		err := d.Delete(3)
		if err != nil {
			t.Errorf("Delete fail(error: %v)", err)
		}
		if !reflect.DeepEqual(*d.tail.data, Commit{UUID: "789", Message: "third commit", Date: "4 november"}) {
			t.Errorf("Delete fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		_ = d.SetCurrent(2)
		err := d.Delete(2)
		if err != nil {
			t.Errorf("Delete fail(error: %v)", err)
		}
		if !reflect.DeepEqual(*d.Current().data, Commit{UUID: "456", Message: "second commit", Date: "3 november"}) ||
			!reflect.DeepEqual(*d.Current().next.data, Commit{UUID: "987", Message: "fourth commit", Date: "5 november"}) {
			t.Errorf("Delete fail")
		}
	}
}

func TestDeleteCurrent(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		err := d.DeleteCurrent()
		if err == nil {
			t.Errorf("DeleteCurrent fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		err := d.DeleteCurrent()
		if err != nil {
			t.Errorf("DeleteCurrent fail(error: %v)", err)
		}
		if !reflect.DeepEqual(*d.Current().data, Commit{UUID: "456", Message: "second commit", Date: "3 november"}) {
			t.Errorf("DeleteCurrent fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		d.curr = d.tail
		err := d.DeleteCurrent()
		if err != nil {
			t.Errorf("DeleteCurrent fail(error: %v)", err)
		}
		if !reflect.DeepEqual(*d.Current().data, Commit{UUID: "789", Message: "third commit", Date: "4 november"}) {
			t.Errorf("DeleteCurrent fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		d.curr = d.head.next
		err := d.DeleteCurrent()
		if err != nil {
			t.Errorf("DeleteCurrent fail(error: %v)", err)
		}
		if !reflect.DeepEqual(*d.Current().data, Commit{UUID: "123", Message: "first commit", Date: "2 november"}) {
			t.Errorf("DeleteCurrent fail")
		}
	}
}

func TestIndex(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		_, err := d.Index()
		if err == nil {
			t.Errorf("Index fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		_ = d.SetCurrent(2)
		n, err := d.Index()
		if err != nil {
			t.Errorf("Index fail(error: %v)", err)
		}
		if n != 2 {
			t.Errorf("Index fail")
		}
	}
}

func TestGetByIndex(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		_, err := d.GetByIndex(10)
		if err == nil {
			t.Errorf("GetByIndex fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		node, err := d.GetByIndex(2)
		if err != nil {
			t.Errorf("GetByIndex fail(error: %v)", err)
		}
		if !reflect.DeepEqual(*node.data, Commit{UUID: "789", Message: "third commit", Date: "4 november"}) {
			t.Errorf("GetByIndex fail")
		}
	}
}

func TestPop(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		node := d.Pop()
		if node != nil {
			t.Errorf("Pop fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		node := d.Pop()
		if !reflect.DeepEqual(*node.data, Commit{UUID: "987", Message: "fourth commit", Date: "5 november"}) {
			t.Errorf("Pop fail")
		}
	}
}

func TestShift(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		node := d.Shift()
		if node != nil {
			t.Errorf("Shift fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		node := d.Shift()
		if !reflect.DeepEqual(*node.data, Commit{UUID: "123", Message: "first commit", Date: "2 november"}) {
			t.Errorf("Shift fail")
		}
	}
}

func TestSearchUUID(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		node := d.SearchUUID("123")
		if node != nil {
			t.Errorf("SearchUUID fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		node := d.SearchUUID("789")
		if !reflect.DeepEqual(*node.data, Commit{UUID: "789", Message: "third commit", Date: "4 november"}) {
			t.Errorf("SearchUUID fail")
		}
	}
}

func TestSearch(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		node := d.Search("some")
		if node != nil {
			t.Errorf("Search fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "5 november"}}
		d.Init(commits)
		node := d.Search("third commit")
		if !reflect.DeepEqual(*node.data, Commit{UUID: "789", Message: "third commit", Date: "4 november"}) {
			t.Errorf("Search fail")
		}
	}
}

func TestReverse(t *testing.T) {
	d := DoubleLinkedList{}
	{
		d.Init([]Commit{})
		newD := d.Reverse()
		if newD.head != nil || newD.tail != nil {
			t.Errorf("Reverse fail")
		}
	}
	{
		commits := []Commit{
			{UUID: "123", Message: "first commit", Date: "2 november"},
			{UUID: "456", Message: "second commit", Date: "3 november"},
			{UUID: "789", Message: "third commit", Date: "4 november"},
			{UUID: "987", Message: "fourth commit", Date: "1 november"}}
		d.Init(commits)
		newD := d.Reverse()
		curN := newD.head
		cur := d.tail
		for curN != nil {
			if !reflect.DeepEqual(*curN.data, *cur.data) {
				t.Errorf("Reverse fail")
			}
			cur = cur.prev
			curN = curN.next
		}
	}
}

func TestLoadData(t *testing.T) {
	d := DoubleLinkedList{}
	{
		err := d.LoadData("wringPath")
		if err == nil {
			t.Errorf("LoadData dont error, but should")
		}
	}
	{
		err := d.LoadData("wrongJson")
		if err == nil {
			t.Errorf("LoadData dont error, but should")
		}
	}
	{
		err := d.LoadData("goodData")
		if err != nil {
			t.Errorf("LoadData unexpexted error: %v", err)
		}
	}
}

func TestQuickSort(t *testing.T) {
	commits := []Commit{
		{UUID: "123", Message: "first commit", Date: "2 november"},
		{UUID: "456", Message: "second commit", Date: "3 november"},
		{UUID: "789", Message: "third commit", Date: "4 november"},
		{UUID: "987", Message: "fourth commit", Date: "1 november"},
	}
	want := []Commit{
		{UUID: "987", Message: "fourth commit", Date: "1 november"},
		{UUID: "123", Message: "first commit", Date: "2 november"},
		{UUID: "456", Message: "second commit", Date: "3 november"},
		{UUID: "789", Message: "third commit", Date: "4 november"},
	}
	QuickSort(commits)
	if !reflect.DeepEqual(commits, want) {
		t.Errorf("QuickSort fail want %v, got %v", want, commits)
	}
}
