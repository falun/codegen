package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/falun/go-genny-codegen/typeset"
)

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

	fmt.Println("\n\nTesting unmarshal")
	is2 := typeset.NewIntSet(typeset.IntStringyKeyFn)
	str := `["1234", "1235"]`
	e = json.Unmarshal([]byte(str), &is2)

	fmt.Println(e)
	fmt.Println(is2.Contains(1234))
	fmt.Println(is2.Contains(1235))
	fmt.Println(is2.Contains(29083))

	fmt.Println("\n\nTesting broken unmarshal")
	is3 := typeset.NewIntSet(typeset.IntStringyKeyFn)
	str = `["a1234", "b1235"]`
	e = json.Unmarshal([]byte(str), &is3)

	fmt.Println(e)
	fmt.Println(is3.Contains(1234))
	fmt.Println(is3.Contains(1235))
	fmt.Println(is3.Contains(29083))

	fmt.Println("\n\nTesting verified unmarshal")
	is4 := typeset.NewIntSet(typeset.IntStringyKeyFn)
	is4.SetVerifier(verifyingIntKeyFn)
	str = `["a1234", "b1235"]`
	e = json.Unmarshal([]byte(str), &is4)

	fmt.Println(e)
	fmt.Println(is4.Contains(1234))
	fmt.Println(is4.Contains(1235))
	fmt.Println(is4.Contains(29083))

}

func verifyingIntKeyFn(s string) error {
	_, e := strconv.ParseInt(s, 10, 64)
	return e
}
