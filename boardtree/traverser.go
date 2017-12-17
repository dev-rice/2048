package boardtree

import (
	"sort"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
	"github.com/donutmonger/2048/stats"
)

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

func GetBestMove(gameBoard [][]int64, getScore scoreFunc) actions.Action {
	r := buildRoot(gameBoard, getScore, 10)

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

func buildRoot(b [][]int64, score scoreFunc, maxdepth int) *Node {
	r := &Node{
		score: 0,
	}
	upBoard, err := board.MoveUp(b, stats.NewScore())
	if err != nil {
		r.up = nil
	} else {
		r.up = buildTree(upBoard, score, maxdepth-1)
	}

	downBoard, err := board.MoveDown(b, stats.NewScore())
	if err != nil {
		r.down = nil
	} else {
		r.down = buildTree(downBoard, score, maxdepth-1)
	}

	leftBoard, err := board.MoveLeft(b, stats.NewScore())
	if err != nil {
		r.left = nil
	} else {
		r.left = buildTree(leftBoard, score, maxdepth-1)
	}

	rightBoard, err := board.MoveRight(b, stats.NewScore())
	if err != nil {
		r.right = nil
	} else {
		r.right = buildTree(rightBoard, score, maxdepth-1)
	}

	return r
}

func buildTree(b [][]int64, score scoreFunc, depth int) *Node {
	if depth == 0 {
		return &Node{
			score: score(b),
		}
	}

	n := &Node{}

	upBoard, err := board.MoveUp(b, stats.NewScore())
	if err != nil {
		n.up = nil
	} else {
		n.up = buildTree(upBoard, score, depth-1)
	}

	downBoard, err := board.MoveDown(b, stats.NewScore())
	if err != nil {
		n.down = nil
	} else {
		n.down = buildTree(downBoard, score, depth-1)
	}

	leftBoard, err := board.MoveLeft(b, stats.NewScore())
	if err != nil {
		n.left = nil
	} else {
		n.left = buildTree(leftBoard, score, depth-1)
	}

	rightBoard, err := board.MoveRight(b, stats.NewScore())
	if err != nil {
		n.right = nil
	} else {
		n.right = buildTree(rightBoard, score, depth-1)
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
