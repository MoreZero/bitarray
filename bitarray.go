package bitarray

type BitArray struct {
	Size uint64
	slot []uint64
}

func NewBitArray(size uint64) *BitArray {
	bitarray := &BitArray{
		Size: size,
	}
	allocSize := (size + 63) >> 6
	bitarray.slot = make([]uint64, allocSize, allocSize)
	return bitarray
}

func (this *BitArray) Clone() *BitArray {
	bitarray := &BitArray{
		Size: this.Size,
	}
	allocSize := (this.Size + 63) >> 6
	bitarray.slot = make([]uint64, allocSize, allocSize)
	return bitarray
}

var MODULO = uint64(0x000000000000003F)
var bitmask = []uint64{
	0x0000000000000001, 0x0000000000000002, 0x0000000000000004, 0x0000000000000008,
	0x0000000000000010, 0x0000000000000020, 0x0000000000000040, 0x0000000000000080,
	0x0000000000000100, 0x0000000000000200, 0x0000000000000400, 0x0000000000000800,
	0x0000000000001000, 0x0000000000002000, 0x0000000000004000, 0x0000000000008000,
	0x0000000000010000, 0x0000000000020000, 0x0000000000040000, 0x0000000000080000,
	0x0000000000100000, 0x0000000000200000, 0x0000000000400000, 0x0000000000800000,
	0x0000000001000000, 0x0000000002000000, 0x0000000004000000, 0x0000000008000000,
	0x0000000010000000, 0x0000000020000000, 0x0000000040000000, 0x0000000080000000,
	0x0000000100000000, 0x0000000200000000, 0x0000000400000000, 0x0000000800000000,
	0x0000001000000000, 0x0000002000000000, 0x0000004000000000, 0x0000008000000000,
	0x0000010000000000, 0x0000020000000000, 0x0000040000000000, 0x0000080000000000,
	0x0000100000000000, 0x0000200000000000, 0x0000400000000000, 0x0000800000000000,
	0x0001000000000000, 0x0002000000000000, 0x0004000000000000, 0x0008000000000000,
	0x0010000000000000, 0x0020000000000000, 0x0040000000000000, 0x0080000000000000,
	0x0100000000000000, 0x0200000000000000, 0x0400000000000000, 0x0800000000000000,
	0x1000000000000000, 0x2000000000000000, 0x4000000000000000, 0x8000000000000000}
var bitumask = []uint64{
	0xfffffffffffffffe, 0xfffffffffffffffd, 0xfffffffffffffffb, 0xfffffffffffffff7,
	0xffffffffffffffef, 0xffffffffffffffdf, 0xffffffffffffffbf, 0xffffffffffffff7f,
	0xfffffffffffffeff, 0xfffffffffffffdff, 0xfffffffffffffbff, 0xfffffffffffff7ff,
	0xffffffffffffefff, 0xffffffffffffdfff, 0xffffffffffffbfff, 0xffffffffffff7fff,
	0xfffffffffffeffff, 0xfffffffffffdffff, 0xfffffffffffbffff, 0xfffffffffff7ffff,
	0xffffffffffefffff, 0xffffffffffdfffff, 0xffffffffffbfffff, 0xffffffffff7fffff,
	0xfffffffffeffffff, 0xfffffffffdffffff, 0xfffffffffbffffff, 0xfffffffff7ffffff,
	0xffffffffefffffff, 0xffffffffdfffffff, 0xffffffffbfffffff, 0xffffffff7fffffff,
	0xfffffffeffffffff, 0xfffffffdffffffff, 0xfffffffbffffffff, 0xfffffff7ffffffff,
	0xffffffefffffffff, 0xffffffdfffffffff, 0xffffffbfffffffff, 0xffffff7fffffffff,
	0xfffffeffffffffff, 0xfffffdffffffffff, 0xfffffbffffffffff, 0xfffff7ffffffffff,
	0xffffefffffffffff, 0xffffdfffffffffff, 0xffffbfffffffffff, 0xffff7fffffffffff,
	0xfffeffffffffffff, 0xfffdffffffffffff, 0xfffbffffffffffff, 0xfff7ffffffffffff,
	0xffefffffffffffff, 0xffdfffffffffffff, 0xffbfffffffffffff, 0xff7fffffffffffff,
	0xfeffffffffffffff, 0xfdffffffffffffff, 0xfbffffffffffffff, 0xf7ffffffffffffff,
	0xefffffffffffffff, 0xdfffffffffffffff, 0xbfffffffffffffff, 0x7fffffffffffffff}

func (this *BitArray) Set(index uint64) {
	this.slot[index>>6] |= bitmask[index&MODULO]
}
func (this *BitArray) UnSet(index uint64) {
	this.slot[index>>6] &= bitumask[index&MODULO]
}

func (this *BitArray) IsSet(index uint64) bool {
	if (this.slot[index>>6] & bitmask[index&MODULO]) > 0 {
		return true
	}
	return false
}

func (this *BitArray) ReInit() {
	for i := 0; i < len(this.slot); i++ {
		this.slot[i] = 0
	}
}

var tailmask = []uint64{0x0000000000000000,
	0x0000000000000001, 0x0000000000000003, 0x0000000000000007, 0x000000000000000f,
	0x000000000000001f, 0x000000000000003f, 0x000000000000007f, 0x00000000000000ff,
	0x00000000000001ff, 0x00000000000003ff, 0x00000000000007ff, 0x0000000000000fff,
	0x0000000000001fff, 0x0000000000003fff, 0x0000000000007fff, 0x000000000000ffff,
	0x000000000001ffff, 0x000000000003ffff, 0x000000000007ffff, 0x00000000000fffff,
	0x00000000001fffff, 0x00000000003fffff, 0x00000000007fffff, 0x0000000000ffffff,
	0x0000000001ffffff, 0x0000000003ffffff, 0x0000000007ffffff, 0x000000000fffffff,
	0x000000001fffffff, 0x000000003fffffff, 0x000000007fffffff, 0x00000000ffffffff,
	0x00000001ffffffff, 0x00000003ffffffff, 0x00000007ffffffff, 0x0000000fffffffff,
	0x0000001fffffffff, 0x0000003fffffffff, 0x0000007fffffffff, 0x000000ffffffffff,
	0x000001ffffffffff, 0x000003ffffffffff, 0x000007ffffffffff, 0x00000fffffffffff,
	0x00001fffffffffff, 0x00003fffffffffff, 0x00007fffffffffff, 0x0000ffffffffffff,
	0x0001ffffffffffff, 0x0003ffffffffffff, 0x0007ffffffffffff, 0x000fffffffffffff,
	0x001fffffffffffff, 0x003fffffffffffff, 0x007fffffffffffff, 0x00ffffffffffffff,
	0x01ffffffffffffff, 0x03ffffffffffffff, 0x07ffffffffffffff, 0x0fffffffffffffff,
	0x1fffffffffffffff, 0x3fffffffffffffff, 0x7fffffffffffffff}

var tailumask = []uint64{0xffffffffffffffff,
	0xfffffffffffffffe, 0xfffffffffffffffc, 0xfffffffffffffff8, 0xfffffffffffffff0,
	0xffffffffffffffe0, 0xffffffffffffffc0, 0xffffffffffffff80, 0xffffffffffffff00,
	0xfffffffffffffe00, 0xfffffffffffffc00, 0xfffffffffffff800, 0xfffffffffffff000,
	0xffffffffffffe000, 0xffffffffffffc000, 0xffffffffffff8000, 0xffffffffffff0000,
	0xfffffffffffe0000, 0xfffffffffffc0000, 0xfffffffffff80000, 0xfffffffffff00000,
	0xffffffffffe00000, 0xffffffffffc00000, 0xffffffffff800000, 0xffffffffff000000,
	0xfffffffffe000000, 0xfffffffffc000000, 0xfffffffff8000000, 0xfffffffff0000000,
	0xffffffffe0000000, 0xffffffffc0000000, 0xffffffff80000000, 0xffffffff00000000,
	0xfffffffe00000000, 0xfffffffc00000000, 0xfffffff800000000, 0xfffffff000000000,
	0xffffffe000000000, 0xffffffc000000000, 0xffffff8000000000, 0xffffff0000000000,
	0xfffffe0000000000, 0xfffffc0000000000, 0xfffff80000000000, 0xfffff00000000000,
	0xffffe00000000000, 0xffffc00000000000, 0xffff800000000000, 0xffff000000000000,
	0xfffe000000000000, 0xfffc000000000000, 0xfff8000000000000, 0xfff0000000000000,
	0xffe0000000000000, 0xffc0000000000000, 0xff80000000000000, 0xff00000000000000,
	0xfe00000000000000, 0xfc00000000000000, 0xf800000000000000, 0xf000000000000000,
	0xe000000000000000, 0xc000000000000000, 0x8000000000000000, 0x0000000000000000}

func (this *BitArray) SetAll() {
	length := len(this.slot) - 1
	i := 0
	for ; i < length; i++ {
		this.slot[i] = 0xffffffffffffffff
	}
	this.slot[i] = tailmask[this.Size&MODULO]
}

func (this *BitArray) HaveSet() bool {
	for i := 0; i < len(this.slot); i++ {
		if this.slot[i] != 0x0000000000000000 {
			return true
		}
	}
	return false
}

//size：或操作的数量
func (this *BitArray) And(obj *BitArray, size uint64) {
	slotsize := size >> 6
	var i uint64
	for ; i < slotsize; i++ {
		this.slot[i] &= obj.slot[i]
	}
	tailsize := size & MODULO
	if tailsize > 0 {
		this.slot[i] &= (obj.slot[i] | tailumask[tailsize])
	}
}

func (this *BitArray) Or(obj *BitArray, size uint64) {
	slotsize := size >> 6
	var i uint64
	for ; i < slotsize; i++ {
		this.slot[i] |= obj.slot[i]
	}
	tailsize := size & MODULO
	if tailsize > 0 {
		this.slot[i] |= (obj.slot[i] & tailmask[tailsize])
	}

}

func (this *BitArray) String() string {
	bytes := make([]byte, this.Size)
	for i := uint64(0); i < this.Size; i++ {
		if this.IsSet(i) {
			bytes[i] = '1'
		} else {
			bytes[i] = '0'
		}
	}
	return string(bytes)
}
