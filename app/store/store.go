package store

type Store struct {
	values map[string]string
}

func NewStore() Store {
	return Store{
		values: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	val, exists := s.values[key]
	return val, exists
}

func (s *Store) Set(key string, val string) {
	s.values[key] = val
}
