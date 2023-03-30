package main

import (
	"fmt"
	"reflect"
)

// START OMIT
func main() {
	v1 := "hi"
	v2 := []string{"hi"}
	vv1, vv2 := reflect.ValueOf(v1), reflect.ValueOf(v2)
	fmt.Println(vv1.Comparable(), vv1.Equal(vv2))
	fmt.Println(vv2.Comparable())
	// fmt.Println(vv2.Interface() == vv2.Interface())
	// vv1, vv2 = reflect.ValueOf(&v1), reflect.ValueOf(&v2)
	// vv2.Elem().Grow(2)
	// fmt.Println(len(v2), cap(v2))
	// vv2.Elem().SetZero()
	// fmt.Println(len(v2), cap(v2))
	// vv1.Elem().SetZero()
	// fmt.Println(v1)
}

// END OMIT
