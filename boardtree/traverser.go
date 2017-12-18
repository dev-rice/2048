package boardtree

import (
	"sort"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/stats"
)

type Traverser struct {
	GetScore scoreFunc
	MaxDepth int
}

type actionTuple struct {
	score  uint64
	action actions.Action
}
type actionTuples []actionTuple

func (s actionTuples) Len() int {
	return len(s)
}
func (s actionTuples) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s actionTuples) Less(i, j int) bool {
	return s[i].score < s[j].score
}

type scoreFunc func(b [][]int64) uint64

func (t Traverser) GetBestMove(gameBoard [][]int64) actions.Action {
	r := buildRoot(gameBoard, t.GetScore, t.MaxDepth)

	a := make([]actionTuple, 0)
	if r.up != nil {
		a = append(a, actionTuple{action: actions.MoveUp, score: r.up.score})
	}
	if r.down != nil {
		a = append(a, actionTuple{action: actions.MoveDown, score: r.down.score})
	}
	if r.left != nil {
		a = append(a, actionTuple{action: actions.MoveLeft, score: r.left.score})
	}
	if r.right != nil {
		a = append(a, actionTuple{action: actions.MoveRight, score: r.right.score})
	}

	sort.Sort(actionTuples(a))
	return a[len(a)-1].action
}

type Node struct {
	score uint64
	up    *Node
	down  *Node
	left  *Node
	right *Node
}

func buildRoot(b [][]int64, getScore scoreFunc, depth int) *Node {
	if depth == 0 {
		return &Node{
			score: getScore(b),
		}
	}

	n := &Node{}

	upBoard, err := board.MoveUp(b, stats.NewScore())
	if err != nil {
		n.up = nil
	} else {
		n.up = buildRoot(upBoard, getScore, depth-1)
	}

	downBoard, err := board.MoveDown(b, stats.NewScore())
	if err != nil {
		n.down = nil
	} else {
		n.down = buildRoot(downBoard, getScore, depth-1)
	}

	leftBoard, err := board.MoveLeft(b, stats.NewScore())
	if err != nil {
		n.left = nil
	} else {
		n.left = buildRoot(leftBoard, getScore, depth-1)
	}

	rightBoard, err := board.MoveRight(b, stats.NewScore())
	if err != nil {
		n.right = nil
	} else {
		n.right = buildRoot(rightBoard, getScore, depth-1)
	}

	n.score = bestNodeScore(n.up, n.down, n.left, n.right)
	return n
}

func bestNodeScore(nodes ...*Node) uint64 {
	best := uint64(0)
	for _, n := range nodes {
		if n != nil {
			if n.score > best {
				best = n.score
			}
		}
	}
	return best
}
