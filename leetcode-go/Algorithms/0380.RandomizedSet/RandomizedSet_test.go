package _380_RandomizedSet

import (
	"fmt"
	"testing"
)

func TestRandomizedSet_Remove(t *testing.T) {
	c := Constructor()
	c.Insert(0)
	c.Insert(1)
	c.Remove(0)
	c.Insert(2)
	c.Remove(1)
	fmt.Println(c.GetRandom())
}
