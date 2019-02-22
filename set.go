package main

/**
	set 操作
 */
type SSet struct {
	value map[string]bool
}

func NewSet() *SSet {
	return &SSet{make(map[string]bool)}
}

func (s *SSet) Add(key string) bool {
	if s.Has(key) {
		return false
	}
	s.value[key] = true
	return true
}

func (s *SSet) Has(key string) bool {
	_, ok := s.value[key]
	return ok
}

func (s *SSet) Remove(key string) {
	delete(s.value, key)
}

func (s *SSet) Clear() {
	for key := range (s.value) {
		delete(s.value, key)
	}
}
