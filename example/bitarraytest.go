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
	start1 := time.Now()
	bit := bitarray.NewBitArray(100000000)
	for i := uint64(0); i < 100000000; i += 2 {
		bit.Set(i)
	}
	start := time.Now()
	for i := uint64(0); i < bit.Size; i++ {
		if bit.IsSet(i) {
			count++
		}
	}
	end := time.Now()
	fmt.Println(count, start.Sub(start1), end.Sub(start))

	start = time.Now()
	bit.ReInit()
	end = time.Now()
	fmt.Println("reinit:", end.Sub(start))
	start = time.Now()
	bit.SetAll()
	end = time.Now()
	fmt.Println("setall:", end.Sub(start))
}

func main() {
	bit := bitarray.NewBitArray(75)
	bit.Set(3)
	bit.Set(6)
	bit.Set(7)
	bit.Set(0)
	bit.Set(68)
	bit.Set(69)
	fmt.Println(bit.String(), bit.HaveSet())
	bit.UnSet(3)
	fmt.Println(bit.String(), bit.HaveSet())
	bit.SetAll()
	fmt.Println(bit.String(), bit.HaveSet())
	bit.ReInit()
	fmt.Println(bit.String(), bit.HaveSet())
	bit.Set(6)
	bit.Set(3)
	bit.Set(9)
	bit.Set(9)
	fmt.Println(bit.String(), bit.HaveSet())
	bit2 := bitarray.NewBitArray(67)
	bit2.Set(65)
	//bit.Set(65)
	bit.Set(68)
	fmt.Println(bit.String(), bit.HaveSet())
	bit.And(bit2, 66)
	fmt.Println(bit2.String(), bit2.HaveSet())
	fmt.Println(bit.String(), bit.HaveSet())
	bit.Set(3)
	bit.Set(9)
	bit.Or(bit2, 66)
	fmt.Println(bit.String(), bit.HaveSet())

	//iter := bit.GetIter(0)
	//for i := int64(0); i < bit.Size; i++ {
	//	fmt.Println(i, iter.Next())
	//}

	//speedtest1() too slow
	//	speedtest2()
}
