package game

type Score struct {
	score int64
}

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
