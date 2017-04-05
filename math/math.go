package math

func MinU32(u32L, u32R uint32) uint32 {
	if u32L <= u32R {
		return u32L
	}
	return u32R
}

func MaxU32(u32L, u32R uint32) uint32 {
	if u32L >= u32R {
		return u32L
	}
	return u32R
}

// return min,max
func MinMaxU32(u32L, u32R uint32) (uint32, uint32) {
	return MinU32(u32L, u32R), MaxU32(u32L, u32R)
}

// return u32M between u32L, u32R
func BoundU32(u32L, u32M, u32R uint32) uint32 {
	u32L, u32R = MinMaxU32(u32L, u32R)
	return MinU32(MaxU32(u32L, u32M), u32R)
}

// 一个数是否是2的幂次方1 2 4 8 16
// 二进制表示的2的幂次方数中只有一个1，后面跟的是n个0
// 将这个数减去1后会发现，仅有的那个1会变为0，而原来的那n个0会变为1
// 此将原来的数与上(&)减去1后的数字，结果为零
func Is2Power(u64Num uint64) bool {
	return (u64Num & (u64Num - 1)) != 0
}

// 是偶数?
// 偶数最后一位一定是0，所以与1 = 0
func IsEven(u64Num uint64) bool {
	return (u64Num & 1) == 0
}

// roundUp takes a uint64 greater than 0 and rounds it up to the next
// power of 2.
func RoundUp(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}
