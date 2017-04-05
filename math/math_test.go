package math

import "testing"

func Test_MaxU32(t *testing.T) {
	if MaxU32(1, 1) != 1 {
		t.Fatal()
	}

	if MaxU32(1, 2) != 2 {
		t.Fatal()
	}

	if MaxU32(2, 1) != 2 {
		t.Fatal()
	}
}

func Test_MinU32(t *testing.T) {
	if MinU32(1, 1) != 1 {
		t.Fatal()
	}

	if MinU32(1, 2) != 1 {
		t.Fatal()
	}

	if MinU32(2, 1) != 1 {
		t.Fatal()
	}
}

func Test_MinMaxU32(t *testing.T) {
	var u32L, u32R uint32

	u32L, u32R = MinMaxU32(1, 2)
	if u32L != 1 || u32R != 2 {
		t.Fatal()
	}

	u32L, u32R = MinMaxU32(2, 1)
	if u32L != 1 || u32R != 2 {
		t.Fatal()
	}

	u32L, u32R = MinMaxU32(1, 1)
	if u32L != 1 || u32R != 1 {
		t.Fatal()
	}
}

func Test_BoundU32(t *testing.T) {
	if BoundU32(1, 1, 1) != 1 {
		t.Fatal()
	}

	if BoundU32(1, 2, 1) != 1 {
		t.Fatal()
	}

	if BoundU32(1, 3, 2) != 2 {
		t.Fatal()
	}

	if BoundU32(2, 3, 1) != 2 {
		t.Fatal()
	}

	if BoundU32(2, 1, 3) != 2 {
		t.Fatal()
	}

	if BoundU32(3, 1, 2) != 2 {
		t.Fatal()
	}
}
