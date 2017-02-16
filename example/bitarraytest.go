package main

import (
	"fmt"
	"github.com/MoreZero/bitarray"
)

func main() {
	bit := bitarray.NetBitArray(10)
	bit.Set(3)
	bit.Set(6)
	bit.Set(7)
	bit.Set(0)
	bit.UnSet(3)
	fmt.Println(bit.String(), bit.HaveSet())
	bit.SetAll()
	fmt.Println(bit.String(), bit.HaveSet())
	bit.ReInit()
	fmt.Println(bit.String(), bit.HaveSet())
	bit.Set(6)
	bit.Set(3)
	bit.Set(9)
	fmt.Println(bit.String(), bit.HaveSet())

	iter := bit.GetIter()
	for i := int64(0); i < bit.Size; i++ {
		fmt.Println(i, iter.Next())
	}

}
