package bitarray

import ()

type Iter struct {
	slotRefer []byte
	cursor    int64
	index     int8
	cache     byte //生成时就初始化为第一个byte
}

func (this *Iter) Next() (isSet bool) {
	if this.index == 8 {
		this.cache = this.slotRefer[this.cursor>>3]
		this.index = 0
	}
	if (this.cache & bitmask[this.index]) == 0 {
		isSet = false
	} else {
		isSet = true
	}
	this.index++
	this.cursor++
	return

}

/////////////////////////////////////////////////////
type BitArray struct {
	Size int64
	slot []byte
}

func NetBitArray(size int64) *BitArray {
	bitarray := &BitArray{
		Size: size,
	}
	bytes := (size + 7) / 8
	bitarray.slot = make([]byte, bytes, bytes)
	return bitarray
}

var bitmask = []byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}
var bitumask = []byte{0xfe, 0xfd, 0xfb, 0xf7, 0xef, 0xdf, 0xbf, 0x7f}

func (this *BitArray) Set(index int64) {
	this.slot[index>>3] |= bitmask[index&0x00000007]
}
func (this *BitArray) UnSet(index int64) {
	this.slot[index>>3] &= bitumask[index&0x00000007]
}

func (this *BitArray) IsSet(index int64) bool {
	if (this.slot[index>>3] & bitmask[index&0x00000007]) > 0 {
		return true
	}
	return false
}

func (this *BitArray) ReInit() {
	for i := 0; i < len(this.slot); i++ {
		this.slot[i] = 0
	}
}

var tailmask = []byte{0x00, 0x01, 0x03, 0x07, 0x0f, 0x1f, 0x3f, 0x7f}

func (this *BitArray) SetAll() {
	length := len(this.slot) - 1
	i := 0
	for ; i < length; i++ {
		this.slot[i] = 0xff
	}
	this.slot[i] = tailmask[this.Size&0x00000007]
}

func (this *BitArray) HaveSet() bool {
	for i := 0; i < len(this.slot); i++ {
		if this.slot[i] != 0x00 {
			return true
		}
	}
	return false
}

func (this *BitArray) String() string {
	bytes := make([]byte, this.Size)
	for i := int64(0); i < this.Size; i++ {
		if this.IsSet(i) {
			bytes[i] = '1'
		} else {
			bytes[i] = '0'
		}
	}
	return string(bytes)
}

func (this *BitArray) GetIter() *Iter {
	return &Iter{
		slotRefer: this.slot,
		cache:     this.slot[0],
	}
}

/*type Iter struct {
	slotRefer []byte
	cursor    int64
	index     int8
	cache     byte //生成时就初始化为第一个byte
}*/
