package entity

type Snake struct {
	Head int
	Tail int
}

type Ladder struct {
	Bottom int
	Top    int
}

type Board struct {
	Size    int
	Snakes  map[int]int // head -> tail
	Ladders map[int]int // bottom -> top
}

func NewBoard(size int, snakes []Snake, ladders []Ladder) *Board {
	snakeMap := make(map[int]int)
	for _, s := range snakes {
		snakeMap[s.Head] = s.Tail
	}

	ladderMap := make(map[int]int)
	for _, l := range ladders {
		ladderMap[l.Bottom] = l.Top
	}

	return &Board{
		Size:    size,
		Snakes:  snakeMap,
		Ladders: ladderMap,
	}
}

// Get next position after snakes/ladders adjustment
func (b *Board) GetNextPosition(pos int) int {
	if tail, ok := b.Snakes[pos]; ok {
		return tail
	}
	if top, ok := b.Ladders[pos]; ok {
		return top
	}
	return pos
}
