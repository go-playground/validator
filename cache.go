package validator

import "sync"

type cachedField struct {
	Idx       int
	Name      string
	AltName   string
	CachedTag *cachedTag
}

type cachedStruct struct {
	Name   string
	fields map[int]cachedField
}

type structCacheMap struct {
	lock sync.RWMutex
	m    map[string]*cachedStruct
}

func (s *structCacheMap) Get(key string) (*cachedStruct, bool) {
	s.lock.RLock()
	value, ok := s.m[key]
	s.lock.RUnlock()
	return value, ok
}

func (s *structCacheMap) Set(key string, value *cachedStruct) {
	s.lock.Lock()
	s.m[key] = value
	s.lock.Unlock()
}

type cachedTag struct {
	isOmitEmpty     bool
	isNoStructLevel bool
	isStructOnly    bool
	diveTag         string
	tags            []*tagVals
}

type tagVals struct {
	tagVals [][]string
	isOrVal bool
	isAlias bool
	tag     string
}

type tagCacheMap struct {
	lock sync.RWMutex
	m    map[string]*cachedTag
}

func (s *tagCacheMap) Get(key string) (*cachedTag, bool) {
	s.lock.RLock()
	value, ok := s.m[key]
	s.lock.RUnlock()

	return value, ok
}

func (s *tagCacheMap) Set(key string, value *cachedTag) {
	s.lock.Lock()
	s.m[key] = value
	s.lock.Unlock()
}
