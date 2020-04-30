package pathfinding

import (
	"container/heap"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	c.Convey("测试优先级队列", t, func() {
		items := map[string]int{
			"V3": 3, "V2": 2, "V4": 4,
		}
		pq := make(PriorityQueue, len(items))
		i := 0
		for _, priority := range items {
			pq[i] = &Node{
				evalCost:  priority,
				heapIndex: i,
			}
			i++
		}
		heap.Init(&pq)
		item := &Node{
			evalCost: 1,
		}
		heap.Push(&pq, item)

		c.So(heap.Pop(&pq).(*Node).evalCost, c.ShouldEqual, 1)
		c.So(heap.Pop(&pq).(*Node).evalCost, c.ShouldEqual, 2)
		c.So(heap.Pop(&pq).(*Node).evalCost, c.ShouldEqual, 3)
		c.So(heap.Pop(&pq).(*Node).evalCost, c.ShouldEqual, 4)
	})

}
