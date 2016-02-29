package typeset

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Type=string,int"

import (
	"encoding/json"
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type Type generic.Type
type TypeToKeyFn func(Type) ([]byte, error)

type TypeSet struct {
	contents map[string]bool
	toKey    TypeToKeyFn
}

func TypeJSONKeyFn(t Type) ([]byte, error) {
	return json.Marshal(t)
}

func TypeStringyKeyFn(t Type) ([]byte, error) {
	return []byte(fmt.Sprintf("%v", t)), nil
}

func NewTypeSet(toKey TypeToKeyFn) TypeSet {
	return TypeSet{
		contents: make(map[string]bool),
		toKey:    toKey,
	}
}

func (ts TypeSet) Contains(t Type) (bool, error) {
	k, e := ts.toKey(t)
	if e != nil {
		return false, e
	}

	_, hasT := ts.contents[string(k)]
	return hasT, nil
}

func (ts TypeSet) Add(t Type) error {
	k, e := ts.toKey(t)
	if e != nil {
		return e
	}

	ts.contents[string(k)] = true
	return nil
}

func (ts TypeSet) MarshalJSON() ([]byte, error) {
	var c []string
	for k, _ := range ts.contents {
		c = append(c, k)
	}

	return json.Marshal(c)
}

func (ts TypeSet) UnmarshalJSON(data []byte) error {
	if ts.toKey == nil {
		return fmt.Errorf("Could not unmarshal into TypeSet, toKey function unset")
	}

	var s []string
	e := json.Unmarshal(data, &s)
	if e != nil {
		return e
	}

	ts.contents = make(map[string]bool)
	for _, v := range s {
		ts.contents[v] = true
	}

	return nil
}
