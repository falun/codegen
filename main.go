package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/falun/go-genny-codegen/typeset"
)

type T struct {
	A string
}

func main() {
	fmt.Println("testing string set")
	ts := typeset.NewStringSet(typeset.StringStringyKeyFn)

	t1 := "aoeu"
	v, e := ts.Contains(t1)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("ts contains t1?", v)

	ts.Add(t1)
	v, e = ts.Contains(t1)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("ts contains t1?", v)

	s, e := json.Marshal(ts)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("ts is: %s\n", string(s))

	fmt.Println("\n\ntesting int set")
	is := typeset.NewIntSet(typeset.IntStringyKeyFn)

	i1 := 1234
	v, e = is.Contains(i1)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("is contains i1?", v)

	is.Add(i1)
	v, e = is.Contains(i1)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("is contains i1?", v)

	is.Add(234)
	is.Add(2882)
	s, e = json.Marshal(is)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("is is: %s\n", string(s))
}
