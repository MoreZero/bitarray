package main

import (
	"fmt"
	"time"

	"github.com/MoreZero/bitarray"
)

/* too slow
func speedtest1() {
	count := 0
	bit := bitarray.NetBitArray(100000000)
	for i := int64(0); i < 100000000; i += 2 {
		bit.Set(i)
	}
	start := time.Now()
	iter := bit.GetIter(0)
	for i := int64(0); i < bit.Size; i++ {
		if iter.Next() {
			count++
		}
	}
	end := time.Now()
	fmt.Println(count, end.Sub(start))
}*/

func speedtest2() {
	count := 0
	bit := bitarray.NetBitArray(100000000)
	for i := int64(0); i < 100000000; i += 2 {
		bit.Set(i)
	}
	start := time.Now()
	for i := int64(0); i < bit.Size; i++ {
		if bit.IsSet(i) {
			count++
		}
	}
	end := time.Now()
	fmt.Println(count, end.Sub(start))
}

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

	iter := bit.GetIter(0)
	for i := int64(0); i < bit.Size; i++ {
		fmt.Println(i, iter.Next())
	}

	//speedtest1() too slow
	speedtest2()
}
