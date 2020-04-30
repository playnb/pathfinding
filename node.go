package pathfinding

type Node struct {
	X         int
	Y         int
	parent    *Node
	evalCost  int //起点计算的代价+到终点的估价
	cost      int //从起点计算的代价
	heapIndex int // only used and maintained by pqueue
}

func (n *Node) priority() int {
	return n.cost + n.evalCost
}

func newNode(x int, y int) *Node {
	return &Node{
		X:      x,
		Y:      y,
		parent: nil,
	}
}

type nodeList struct {
	nodes         map[int]*Node
	width, height int
}

func newNodeList(width, height int) *nodeList {
	return &nodeList{
		nodes:  make(map[int]*Node, width*height),
		width:  width,
		height: height,
	}
}

func (n *nodeList) addNode(node *Node) {
	n.nodes[node.X+node.Y*n.width] = node
}

func (n *nodeList) getNode(x, y int) *Node {
	return n.nodes[x+y*n.width]
}

func (n *nodeList) removeNode(node *Node) {
	delete(n.nodes, node.X+node.Y*n.width)
}

func (n *nodeList) hasNode(node *Node) bool {
	if n.getNode(node.X, node.Y) != nil {
		return true
	}
	return false
}

func (n *nodeList) createNode(x, y int) *Node {
	node := n.getNode(x, y)
	if node == nil {
		node = newNode(x, y)
	}
	return node
}
