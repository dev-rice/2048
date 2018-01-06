package ai

import (
	"sort"

	"github.com/donutmonger/2048/actions"
	"github.com/donutmonger/2048/board"
)

type Traverser struct {
	GetRating ratingFunc
	MaxDepth  int
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

type ratingFunc func(b [][]int64) uint64

func (t Traverser) GetBestMove(gameBoard [][]int64) actions.Action {
	r := buildRoot(gameBoard, t.GetRating, t.MaxDepth)

	a := make([]actionTuple, 0)
	if r.up != nil {
		a = append(a, actionTuple{action: actions.MoveUp, score: r.up.rating})
	}
	if r.down != nil {
		a = append(a, actionTuple{action: actions.MoveDown, score: r.down.rating})
	}
	if r.left != nil {
		a = append(a, actionTuple{action: actions.MoveLeft, score: r.left.rating})
	}
	if r.right != nil {
		a = append(a, actionTuple{action: actions.MoveRight, score: r.right.rating})
	}

	sort.Sort(actionTuples(a))
	return a[len(a)-1].action
}

type Node struct {
	rating uint64
	up     *Node
	down   *Node
	left   *Node
	right  *Node
}

func buildRoot(b [][]int64, getRating ratingFunc, depth int) *Node {
	if depth == 0 {
		return &Node{
			rating: getRating(b),
		}
	}

	n := &Node{}

	upBoard, _, err := board.MoveUp(b)
	if err != nil {
		n.up = nil
	} else {
		n.up = buildRoot(upBoard, getRating, depth-1)
	}

	downBoard, _, err := board.MoveDown(b)
	if err != nil {
		n.down = nil
	} else {
		n.down = buildRoot(downBoard, getRating, depth-1)
	}

	leftBoard, _, err := board.MoveLeft(b)
	if err != nil {
		n.left = nil
	} else {
		n.left = buildRoot(leftBoard, getRating, depth-1)
	}

	rightBoardCompressed, _, err := board.MoveRight(board.CompressBoardGrid(b))
	if err != nil {
		n.right = nil
	} else {
		n.right = buildRoot(board.UncompressBoard(rightBoardCompressed), getRating, depth-1)
	}

	n.rating = bestNodeRating(n.up, n.down, n.left, n.right)
	return n
}

func bestNodeRating(nodes ...*Node) uint64 {
	best := uint64(0)
	for _, n := range nodes {
		if n != nil {
			if n.rating > best {
				best = n.rating
			}
		}
	}
	return best
}
