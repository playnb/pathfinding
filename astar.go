package pathfinding

type Map interface {
	GetWidth() int
	GetHeight() int
	GetNeighbors(x int, y int) [][3]int
}

type Evaluation func(from *Node, to *Node) int

//TODO 这里new了一堆Node,其实可以做成NodePool,每次寻路获一个pool出来
func AStar(mapData Map, startX, startY, stopX, stopY int, eval Evaluation) []*Node {
	closedSet := newNodeList(mapData.GetWidth(), mapData.GetHeight())
	openSet := newNodeList(mapData.GetWidth(), mapData.GetHeight())
	pq := make(PriorityQueue, 0, mapData.GetWidth()*mapData.GetHeight()) // heap, used to find minF

	freeSet := newNodeList(mapData.GetWidth(), mapData.GetHeight())
	start := freeSet.createNode(startX, startY)
	stop := freeSet.createNode(stopX, stopY)
	openSet.addNode(start)
	pq.PushNode(start)

	for len(openSet.nodes) != 0 {
		current := pq.PopNode()
		openSet.removeNode(current)
		closedSet.addNode(current)

		//检查是否找到目标点
		if current.X == stop.X && current.Y == stop.Y {
			return retracePath(current)
		}

		for _, pos := range mapData.GetNeighbors(current.X, current.Y) {
			x := pos[0]
			y := pos[1]
			cost := current.cost + pos[2]

			if (x < 0) || (x >= mapData.GetWidth()) || (y < 0) || (y >= mapData.GetHeight()) {
				//地图外的点不做处理
				continue
			}
			if closedSet.getNode(x, y) != nil {
				//close表中的节点
				continue
			}

			neighbor := freeSet.createNode(x, y)
			if !openSet.hasNode(neighbor) {
				//插入Open表中
				neighbor.parent = current
				neighbor.cost = cost
				neighbor.evalCost = eval(neighbor, stop)
				openSet.addNode(neighbor)
				pq.PushNode(neighbor)
			} else if cost < neighbor.cost {
				neighbor.parent = current
				neighbor.cost = cost
				pq.UpdateNode(neighbor)
			}
		}
	}
	return nil
}

func retracePath(currentNode *Node) []*Node {
	var path []*Node
	path = append(path, currentNode)
	for currentNode.parent != nil {
		path = append(path, currentNode.parent)
		currentNode = currentNode.parent
	}
	//Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
