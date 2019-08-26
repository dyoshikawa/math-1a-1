package main

import (
	"fmt"
	"reflect"
	"sort"
)

type ByValue []string

func (s ByValue) Len() int           { return len(s) }
func (s ByValue) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByValue) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	strs := []string{"a", "b", "c", "d"}
	ptns := [][]string{}
	for _, v := range strs {
		for _, w := range strs {
			for _, z := range strs {
				if v == w || v == z || w == z {
					continue
				}
				ptn := []string{v, w, z}
				sort.Sort(ByValue(ptn))
				ptns = append(ptns, ptn)
			}
		}
	}
	res := [][]string{}
	for _, v := range ptns {
		next := false
		for _, w := range res {
			if reflect.DeepEqual(v, w) {
				next = true
				break
			}
		}
		if next {
			continue
		}
		res = append(res, v)
	}
	fmt.Println(len(res))
}
