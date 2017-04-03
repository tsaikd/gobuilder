package checkerror

import (
	"go/types"
	"sort"
)

type errorFactoryList struct {
	objpool   map[types.Object]bool
	namepool  map[string]types.Object
	twicepool map[string]int8
}

func (t *errorFactoryList) ensureInit() {
	if t.objpool == nil {
		t.objpool = map[types.Object]bool{}
	}
	if t.namepool == nil {
		t.namepool = map[string]types.Object{}
	}
	if t.twicepool == nil {
		t.twicepool = map[string]int8{}
	}
}

func (t *errorFactoryList) addObject(obj types.Object) {
	t.ensureInit()
	t.objpool[obj] = true
	t.namepool[getTypesObjectName(obj)] = obj
}

func (t *errorFactoryList) removeName(name string) {
	t.ensureInit()
	if obj, exist := t.namepool[name]; exist {
		delete(t.namepool, name)
		delete(t.objpool, obj)
	}
}

// removeNameTwice remove name if called twice
func (t *errorFactoryList) removeNameTwice(name string) {
	t.ensureInit()
	if obj, exist := t.namepool[name]; exist {
		if count, twiceExist := t.twicepool[name]; twiceExist {
			t.twicepool[name] = count + 1
			delete(t.namepool, name)
			delete(t.objpool, obj)
		} else {
			t.twicepool[name] = 1
		}
	}
}

func (t *errorFactoryList) isEmpty() bool {
	t.ensureInit()
	return len(t.objpool) < 1 && len(t.namepool) < 1
}

func (t *errorFactoryList) sortedObjects() []types.Object {
	sorted := []types.Object{}
	names := []string{}
	for name := range t.namepool {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		sorted = append(sorted, t.namepool[name])
	}
	return sorted
}
