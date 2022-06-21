package service

type State int

const (
	Idle State = iota
	EditsName
	InGame
)

type StatePool interface {
	Set(key string, state State)
	Get(key string) State
}

type statePool struct {
	pool map[string]State
}

func NewStatePool() *statePool {
	return &statePool{
		pool: make(map[string]State),
	}
}

func (s statePool) Set(key string, state State) {
	s.pool[key] = state
}

func (s statePool) Get(key string) State {
	if s, ok := s.pool[key]; ok {
		return s
	}
	return Idle
}
