package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())
		require.Equal(t, 30, l.Back().Value)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Front().Next.Next.Value)

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())
		require.Equal(t, 30, l.Front().Next.Value)

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
	t.Run("with move from middle", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)          // [10]
		l.PushBack(20)           // [10, 20]
		l.PushBack(30)           // [10, 20, 30]
		middle := l.Front().Next // 20
		l.MoveToFront(middle)
		require.Equal(t, 20, l.Front().Value)
	})
	t.Run("very complex", func(t *testing.T) {
		l := NewList()
		var middle *ListItem
		for i := 10; i < 101; i += 10 {
			a := l.PushBack(i)
			if i == 50 {
				middle = a
			}
		} //[10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, elems)
		require.Equal(t, 50, middle.Value)
		l.MoveToFront(middle) //[50, 10, 20, 30, 40, 60, 70, 80, 90, 100]
		l.Remove(middle)      //[10, 20, 30, 40, 60, 70, 80, 90, 100]
		require.Equal(t, 60, l.Front().Next.Next.Next.Next.Value)
		require.Equal(t, 40, l.Front().Next.Next.Next.Value)
		l.Remove(l.Front()) //[20, 30, 40, 60, 70, 80, 90, 100]
		l.Remove(l.Back())  //[20, 30, 40, 60, 70, 80, 90]
		require.Equal(t, 70, l.Front().Next.Next.Next.Next.Value)
		l.MoveToFront(l.Back())  //[90, 20, 30, 40, 60, 70, 80]
		l.MoveToFront(l.Front()) //[90, 20, 30, 40, 60, 70, 80]
		require.Equal(t, 80, l.Back().Value)
		elems = make([]int, 0, l.Len())
		for i := l.Back(); i != nil; i = i.Prev {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{80, 70, 60, 40, 30, 20, 90}, elems)
	})
}
