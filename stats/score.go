package stats

type Score struct {
	score int64
}

// Possible stats to add
//	1. number of moves
//	1. game time
//	1. time for each move
//  1.

func NewScore() *Score {
	return &Score{
		score: 0,
	}
}

func (s *Score) Add(n int64) {
	s.score += n
}

func (s Score) Get() int64 {
	return s.score
}
