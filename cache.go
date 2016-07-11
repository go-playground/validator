package validator

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
)

// type cachedField struct {
// 	Idx       int
// 	Name      string
// 	AltName   string
// 	CachedTag *cachedTag
// }

// type cachedStruct struct {
// 	Name   string
// 	fields map[int]cachedField
// }

// type structCacheMap struct {
// 	lock sync.RWMutex
// 	m    map[reflect.Type]*cachedStruct
// }

// func (s *structCacheMap) Get(key reflect.Type) (*cachedStruct, bool) {
// 	s.lock.RLock()
// 	value, ok := s.m[key]
// 	s.lock.RUnlock()
// 	return value, ok
// }

// func (s *structCacheMap) Set(key reflect.Type, value *cachedStruct) {
// 	s.lock.Lock()
// 	s.m[key] = value
// 	s.lock.Unlock()
// }

// type cachedTag struct {
// 	tag             string
// 	isOmitEmpty     bool
// 	isNoStructLevel bool
// 	isStructOnly    bool
// 	diveTag         string
// 	tags            []*tagVals
// }

// type tagVals struct {
// 	tagVals [][]string
// 	isOrVal bool
// 	isAlias bool
// 	tag     string
// }

// type tagCacheMap struct {
// 	lock sync.RWMutex
// 	m    map[string]*cachedTag
// }

// func (s *tagCacheMap) Get(key string) (*cachedTag, bool) {
// 	s.lock.RLock()
// 	value, ok := s.m[key]
// 	s.lock.RUnlock()

// 	return value, ok
// }

// func (s *tagCacheMap) Set(key string, value *cachedTag) {
// 	s.lock.Lock()
// 	s.m[key] = value
// 	s.lock.Unlock()
// }

// ******* New Cache ***************

type structCache struct {
	lock sync.Mutex
	m    atomic.Value // map[reflect.Type]*cStruct
}

func (sc *structCache) Get(key reflect.Type) (c *cStruct, found bool) {
	c, found = sc.m.Load().(map[reflect.Type]*cStruct)[key]
	return
}

func (sc *structCache) Set(key reflect.Type, value *cStruct) {

	m := sc.m.Load().(map[reflect.Type]*cStruct)

	nm := make(map[reflect.Type]*cStruct, len(m)+1)
	for k, v := range m {
		nm[k] = v
	}
	nm[key] = value
	sc.m.Store(nm)
}

type tagCache struct {
	lock sync.Mutex
	m    atomic.Value // map[string]*cTag
}

func (tc *tagCache) Get(key string) (c *cTag, found bool) {
	c, found = tc.m.Load().(map[string]*cTag)[key]
	return
}

func (tc *tagCache) Set(key string, value *cTag) {

	m := tc.m.Load().(map[string]*cTag)

	nm := make(map[string]*cTag, len(m)+1)
	for k, v := range m {
		nm[k] = v
	}
	nm[key] = value
	tc.m.Store(nm)
}

type cStruct struct {
	Name   string
	fields map[int]*cField
	fn     StructLevelFunc
}

type cField struct {
	Idx     int
	Name    string
	AltName string
	cTags   *cTag
}

// TODO: investigate using enum instead of so many booleans, may be faster
// but let's get the new cache system working first
type cTag struct {
	tag             string
	aliasTag        string
	actualAliasTag  string
	param           string
	hasAlias        bool
	isOmitEmpty     bool
	isNoStructLevel bool
	isStructOnly    bool
	isDive          bool
	isOrVal         bool
	exists          bool
	hasTag          bool
	fn              Func
	next            *cTag
}

// TODO: eliminate get and set functions from cache, they are pure overhead for nicer syntax.

func (v *Validate) extractStructCache(current reflect.Value, sName string) *cStruct {

	v.structCache.lock.Lock()
	defer v.structCache.lock.Unlock()

	typ := current.Type()

	// could have been multiple trying to access, but once first is done this ensures struct
	// isn't parsed again.
	cs, ok := v.structCache.Get(typ)
	if ok {
		// v.structCache.lock.Unlock()
		return cs
	}

	cs = &cStruct{Name: sName, fields: make(map[int]*cField), fn: v.structLevelFuncs[typ]}

	numFields := current.NumField()

	var ctag *cTag
	var fld reflect.StructField
	var tag string
	var customName string

	for i := 0; i < numFields; i++ {

		fld = typ.Field(i)

		// if fld.PkgPath != blank {
		if !fld.Anonymous && fld.PkgPath != blank {
			continue
		}

		tag = fld.Tag.Get(v.tagName)

		// if len(tag) == 0 || tag == skipValidationTag {
		if tag == skipValidationTag {
			continue
		}

		customName = fld.Name

		if v.fieldNameTag != blank {

			name := strings.SplitN(fld.Tag.Get(v.fieldNameTag), ",", 2)[0]

			// dash check is for json "-" (aka skipValidationTag) means don't output in json
			if name != "" && name != skipValidationTag {
				customName = name
			}
		}

		// fmt.Println("Finding Struct Tag", sName, "FLD:", fld.Name)
		// ctag, ok := v.tagCache.Get(tag)
		// if !ok {
		// fmt.Println("Not Found, Lock then check again for parallel operations")
		// v.tagCache.lock.Lock()
		// defer func() {
		// 	fmt.Println("Ulocking")
		// 	v.tagCache.lock.Unlock()
		// 	fmt.Println("Unlocked")
		// }()
		// defer v.tagCache.lock.Unlock()

		// if ctag, ok = v.tagCache.Get(tag); !ok {

		// fmt.Println("parsing tag", tag)
		// ctag = v.parseFieldTags(tag, fld.Name)
		// NOTE: cannot use tag cache, because tag may be equal, but things like alias may be different
		// and so only struct level caching can be used

		if len(tag) > 0 {
			ctag, _ = v.parseFieldTagsRecursive(tag, fld.Name, blank, false)
		} else {
			// even if field doesn't have validations need cTag for traversing to potential inner/nested
			// elements of the field.
			ctag = new(cTag)
		}
		// fmt.Println("Done Parsing")
		// v.tagCache.Set(tag, ctag)
		// fmt.Println("Tag Cahed")
		// }

		// fmt.Println("Ulocking")
		// v.tagCache.lock.Unlock()
		// fmt.Println("Unlocked")
		// }

		// fmt.Println(tag, ctag)

		cs.fields[i] = &cField{Idx: i, Name: fld.Name, AltName: customName, cTags: ctag}
	}

	// If not anonymous type; they have to be parsed every time because if interface
	// a different struct could be used...
	// if len(sName) > 0 {
	// fmt.Println(typ)
	v.structCache.Set(typ, cs)
	// }

	// v.structCache.lock.Unlock()

	return cs
}

// func (v *Validate) parseFieldTags(tag, fieldName string) (ctag *cTag) {

// 	// ctag := &cTag{tag: tag}

// 	// fmt.Println(tag)
// 	ctag, _ = v.parseFieldTagsRecursive(tag, fieldName, blank, false)

// 	v.tagCache.Set(tag, ctag)

// 	// fmt.Println(ctag)
// 	return
// }

// TODO: Optimize for to not Split but ust for over string chunk, by chunk

func (v *Validate) parseFieldTagsRecursive(tag string, fieldName string, alias string, hasAlias bool) (firstCtag *cTag, current *cTag) {
	// func (v *Validate) parseFieldTagsRecursive(ctag *cTag, tag, fieldName, alias string, isAlias bool) bool {

	// if tag == blank {
	// 	return
	// }
	// fmt.Println("Parsing, depath:", depth)
	var t string
	var ok bool
	noAlias := len(alias) == 0
	// var tmpCtag *cTag
	// var start,end int
	// var ctag *cTag
	// var lastCtag *cTag
	// ctag := &cTag{tag: tag}
	tags := strings.Split(tag, tagSeparator)

	// fmt.Println(len(tags), tags)

	for i := 0; i < len(tags); i++ {
		// for i := 0; i < len(tags); i++ {

		t = tags[i]

		if noAlias {
			alias = t
		}
		// _, found := v.aliasValidators[t]
		// fmt.Println(i, t, found)

		// if len(t) == 0 {
		// 	continue
		// }
		// if i == 0 {
		// 	current = &cTag{aliasTag: alias, hasAlias: hasAlias}
		// 	firstCtag = current
		// } else {
		// 	current.next = &cTag{aliasTag: alias, hasAlias: hasAlias}
		// 	current = current.next
		// }

		if v.hasAliasValidators {
			// check map for alias and process new tags, otherwise process as usual
			if tagsVal, found := v.aliasValidators[t]; found {

				// fmt.Println(tagsVal)

				if i == 0 {
					firstCtag, current = v.parseFieldTagsRecursive(tagsVal, fieldName, t, true)
				} else {
					// fmt.Println("BEFORE ALIAS:", current)
					// diveCurr := current
					next, curr := v.parseFieldTagsRecursive(tagsVal, fieldName, t, true)
					// fmt.Println("ALIAS:", next, curr)
					current.next, current = next, curr

					// fmt.Println("AFTER current", diveCurr)
					// current.next, current = v.parseFieldTagsRecursive(tagsVal, fieldName, t, true)
				}

				continue

				// 		// tmpCtag, lastCtag = v.parseFieldTagsRecursive(tagsVal, fieldName, t, true)

				// 		// if ctag == nil {
				// 		// 	ctag = tmpCtag
				// 		// } else {
				// 		// 	ctag.next = tmpCtag
				// 		// }
			}
		}

		// if i == 0 {
		// 	firstCtag = current
		// }

		// 		type cTag struct {
		// 	tag             string
		// 	aliasTag        string
		// 	hasAlias        bool
		// 	isOmitEmpty     bool
		// 	isNoStructLevel bool
		// 	isStructOnly    bool
		// 	isDive          bool
		// 	isOrVal         bool
		// 	fn              Func
		// 	next            *cTag
		// }

		// if lastCtag == nil {
		// 	lastCtag = ctag
		// }

		if i == 0 {
			current = &cTag{aliasTag: alias, hasAlias: hasAlias, hasTag: true}
			firstCtag = current
		} else {
			current.next = &cTag{aliasTag: alias, hasAlias: hasAlias, hasTag: true}
			current = current.next
		}

		switch t {

		case diveTag:
			current.isDive = true
			// fmt.Println("DIVE CURRENT", current)
			continue
			// 	cTag.diveTag = tag
			// 	tVals := &tagVals{tagVals: [][]string{{t}}}
			// 	cTag.tags = append(cTag.tags, tVals)
			// 	return true

		case omitempty:
			current.isOmitEmpty = true
			continue

		case structOnlyTag:
			current.isStructOnly = true
			continue

		case noStructLevelTag:
			current.isNoStructLevel = true
			continue

		case existsTag:
			current.exists = true
			continue

		default:

			// if a pipe character is needed within the param you must use the utf8Pipe representation "0x7C"
			orVals := strings.Split(t, orSeparator)

			// // if no or values
			// if len(orVals) == 1 {
			// 	current.fn, ok = v.validationFuncs[t]
			// 	if !ok {
			// 		panic(strings.TrimSpace(fmt.Sprintf(undefinedValidation, fieldName)))
			// 	}

			// } else {

			// 		tagVal := &tagVals{isAlias: isAlias, isOrVal: len(orVals) > 1, tagVals: make([][]string, len(orVals))}
			// 		cTag.tags = append(cTag.tags, tagVal)

			// var key string
			// var param string

			for j := 0; j < len(orVals); j++ {
				// for i, val := range orVals {

				// if v.hasAliasValidators {
				// 	// check map for alias and process new tags, otherwise process as usual
				// 	if tagsVal, ok := v.aliasValidators[orVals[j]]; ok {

				// 		if i == 0 {
				// 			firstCtag, current = v.parseFieldTagsRecursive(tagsVal, fieldName, orVals[j], true)
				// 		} else {
				// 			current.next, current = v.parseFieldTagsRecursive(tagsVal, fieldName, orVals[j], true)
				// 		}

				// 		continue

				// 		// tmpCtag, lastCtag = v.parseFieldTagsRecursive(tagsVal, fieldName, t, true)

				// 		// if ctag == nil {
				// 		// 	ctag = tmpCtag
				// 		// } else {
				// 		// 	ctag.next = tmpCtag
				// 		// }
				// 	}
				// }

				vals := strings.SplitN(orVals[j], tagKeySeparator, 2)

				if noAlias {
					alias = vals[0]
					current.aliasTag = alias
				} else {
					// alias = t
					current.actualAliasTag = t
				}

				if j > 0 {
					current.next = &cTag{aliasTag: alias, actualAliasTag: current.actualAliasTag, hasAlias: hasAlias, hasTag: true}
					current = current.next
					// current.next=&
				}
				// else{
				current.tag = vals[0]
				if len(current.tag) == 0 {
					panic(strings.TrimSpace(fmt.Sprintf(invalidValidation, fieldName)))
				}

				// _, found := v.validationFuncs[current.tag]
				// fmt.Println("TAG", current.tag, "FOund:", found)

				if current.fn, ok = v.validationFuncs[current.tag]; !ok {
					// fmt.Println("I'm panicing!")
					panic(strings.TrimSpace(fmt.Sprintf(undefinedValidation, fieldName)))
				}

				current.isOrVal = len(orVals) > 1
				// }

				// tagVal.tag = key

				// if isAlias {
				// 	tagVal.tag = alias
				// }

				// if key == blank {
				// 	panic(strings.TrimSpace(fmt.Sprintf(invalidValidation, fieldName)))
				// }

				if len(vals) > 1 {
					current.param = strings.Replace(strings.Replace(vals[1], utf8HexComma, ",", -1), utf8Pipe, "|", -1)
				}

				// tagVal.tagVals[i] = []string{key, param}
			}
			// }
		}
	}

	// if depth > 0 {
	// 	// _, found := v.aliasValidators[t]
	// 	fmt.Println("WTF", len(tags), tags, firstCtag, current)
	// 	panic("WTF")
	// }

	return
}
