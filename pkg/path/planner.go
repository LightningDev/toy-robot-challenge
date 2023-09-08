package path

import (
	"container/heap"

	"github.com/LightningDev/toy-robot-challenge/pkg/position"
	"github.com/LightningDev/toy-robot-challenge/pkg/table"
)

// Store calculation of your node
// can point to previous one
type Node struct {
	position  position.Position
	cost      int
	heuristic int
	totalCost int
	prevNode  *Node
	index     int
}

type NodeQueue []*Node

func (nq NodeQueue) Len() int { return len(nq) }

func (nq NodeQueue) Less(i, j int) bool {
	return nq[i].totalCost < nq[j].totalCost
}

func (nq NodeQueue) Swap(i, j int) {
	nq[i], nq[j] = nq[j], nq[i]
	nq[i].index = i
	nq[j].index = j
}

func (nq *NodeQueue) Push(x interface{}) {
	n := len(*nq)
	node := x.(*Node)
	node.index = n
	*nq = append(*nq, node)
}

func (nq *NodeQueue) Pop() interface{} {
	old := *nq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*nq = old[0 : n-1]
	return node
}

func heuristic(pos1, pos2 position.Position) int {
	return abs(pos1.X-pos2.X) + abs(pos1.Y-pos2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// A* Algorigthm
func Find(start position.Position, goal position.Position, table table.Table) []position.Position {
	openList := make(NodeQueue, 0)                 // list of node we haven't visited
	closedList := make(map[position.Position]bool) // list of node we have visited

	startNode := &Node{position: start, cost: 0}               // start at robot position
	startNode.heuristic = heuristic(start, goal)               // calculate shortest score to target
	startNode.totalCost = startNode.cost + startNode.heuristic // total cost at the beginning
	heap.Push(&openList, startNode)

	// UP, DOWN, LEFT, RIGHT
	directions := [4]position.Position{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}

	for len(openList) > 0 {
		current := heap.Pop(&openList).(*Node)
		if current.position == goal {
			var path []position.Position
			for current != nil {
				path = append([]position.Position{current.position}, path...)
				current = current.prevNode
			}
			return path
		}

		closedList[current.position] = true
		for _, dir := range directions {
			neighborPos := position.Position{X: current.position.X + dir.X, Y: current.position.Y + dir.Y}

			if neighborPos.X < 0 || neighborPos.X >= table.Width || neighborPos.Y < 0 || neighborPos.Y >= table.Height {
				continue
			}

			// If this is obstacle, skip
			if !table.IsValidPosition(neighborPos.X, neighborPos.Y) {
				continue
			}

			// If already visited, skip
			if closedList[neighborPos] {
				continue
			}

			tentativeCost := current.cost + 1

			found := false
			var neighborNode *Node
			for _, node := range openList {
				if node.position == neighborPos {
					found = true
					neighborNode = node
					break
				}
			}

			if !found || tentativeCost < neighborNode.cost {
				if !found {
					neighborNode = &Node{position: neighborPos}
					heap.Push(&openList, neighborNode)
				} else {
					heap.Remove(&openList, neighborNode.index)
				}

				neighborNode.prevNode = current
				neighborNode.cost = tentativeCost
				neighborNode.heuristic = heuristic(neighborPos, goal)
				neighborNode.totalCost = neighborNode.cost + neighborNode.heuristic
				heap.Push(&openList, neighborNode)
			}
		}
	}
	return nil
}
