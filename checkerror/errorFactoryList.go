package checkerror

import (
	"go/types"
	"sort"
)

type errorFactoryList struct {
	objpool  map[types.Object]bool
	namepool map[string]types.Object
}

func (t *errorFactoryList) ensureInit() {
	if t.objpool == nil {
		t.objpool = map[types.Object]bool{}
	}
	if t.namepool == nil {
		t.namepool = map[string]types.Object{}
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
