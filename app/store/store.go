package store

import "time"

type Expiry struct {
	Enabled  bool
	ExpireAt time.Time
}

type StoreValue struct {
	String string
	Expiry Expiry
}

type Store struct {
	values map[string]StoreValue
}

func NewStore() Store {
	return Store{
		values: make(map[string]StoreValue),
	}
}

func (s *Store) Get(key string) (string, bool) {
	val, exists := s.values[key]

	// check if has expired
	now := time.Now()

	if val.Expiry.Enabled && val.Expiry.ExpireAt.Before(now) {
		return "", false
	}

	return val.String, exists
}

func (s *Store) Set(key string, val string) {
	s.values[key] = StoreValue{
		String: val,
		Expiry: Expiry{},
	}
}

func (s *Store) SetWithExpiry(key string, val string, ttl time.Duration) {
	expireAt := time.Now().Add(ttl)

	storeValue := StoreValue{
		String: val,
		Expiry: Expiry{
			Enabled:  true,
			ExpireAt: expireAt,
		},
	}

	s.values[key] = storeValue
}
