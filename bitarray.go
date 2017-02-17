package bitarray

/* too slow
type Iter struct {
	slotRefer  []byte
	index      int8
	slotcursor int64
}

func (this *Iter) Next() (isSet bool) {
	if (this.slotRefer[this.slotcursor] & bitmask[this.index]) > 0 {
		isSet = true
	} else {
		isSet = false
	}
	this.index += 1
	if this.index == 8 {
		this.slotcursor += 1
		this.index = 0
	}
	return

}*/

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

var MODULO = int64(0x0000000000000007)
var bitmask = []byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}
var bitumask = []byte{0xfe, 0xfd, 0xfb, 0xf7, 0xef, 0xdf, 0xbf, 0x7f}

func (this *BitArray) Set(index int64) {
	this.slot[index>>3] |= bitmask[index&MODULO]
}
func (this *BitArray) UnSet(index int64) {
	this.slot[index>>3] &= bitumask[index&MODULO]
}

func (this *BitArray) IsSet(index int64) bool {
	if (this.slot[index>>3] & bitmask[index&MODULO]) > 0 {
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
	this.slot[i] = tailmask[this.Size&MODULO]
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

/* too slow
func (this *BitArray) GetIter(skip int64) *Iter {
	if skip == 0 {
		return &Iter{
			slotRefer: this.slot,
		}
	} else {
		return &Iter{
			slotRefer: this.slot,
			index:     int8(skip & MODULO),
		}
	}
}*/
