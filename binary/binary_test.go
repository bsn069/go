package binary

import "testing"

func Test_PushPopU8(t *testing.T) {
	var byBuf []byte
	var byTemp []byte
	byBuf = make([]byte, 10)

	byTemp = byBuf
	for i := 0; i < len(byBuf); i++ {
		byTemp = PushU8(byTemp, uint8(i))
	}

	var u8Temp uint8
	byTemp = byBuf
	for i := 0; i < len(byBuf); i++ {
		byTemp = PopU8(byTemp, &u8Temp)
		if u8Temp != uint8(i) {
			t.Fatal()
		}
	}
}

func Test_PushPopU16(t *testing.T) {
	var byBuf []byte
	var byTemp []byte
	byBuf = make([]byte, 10*2)

	byTemp = byBuf
	for i := 0; i < len(byBuf); i++ {
		byTemp = PushU16(byTemp, uint16(i))
		if byTemp == nil {
			break
		}
	}

	var u16Temp uint16
	byTemp = byBuf
	for i := 0; i < len(byBuf); i++ {
		byTemp = PopU16(byTemp, &u16Temp)
		if byTemp == nil {
			break
		}
		if u16Temp != uint16(i) {
			t.Fatal()
		}
	}
}

func Test_PushPopU32(t *testing.T) {
	var byBuf []byte
	var byTemp []byte
	byBuf = make([]byte, 10*4)

	byTemp = byBuf
	for i := 0; i < len(byBuf); i++ {
		byTemp = PushU32(byTemp, uint32(i))
		if byTemp == nil {
			break
		}
	}

	var u32Temp uint32
	byTemp = byBuf
	for i := 0; i < len(byBuf); i++ {
		byTemp = PopU32(byTemp, &u32Temp)
		if byTemp == nil {
			break
		}
		if u32Temp != uint32(i) {
			t.Fatal()
		}
	}
}

func Test_PushPopU64(t *testing.T) {
	var byBuf []byte
	var byTemp []byte
	byBuf = make([]byte, 10*8)

	byTemp = byBuf
	for i := 0; i < len(byBuf); i++ {
		byTemp = PushU64(byTemp, uint64(i))
		if byTemp == nil {
			break
		}
	}

	var u64Temp uint64
	byTemp = byBuf
	for i := 0; i < len(byBuf); i++ {
		byTemp = PopU64(byTemp, &u64Temp)
		if byTemp == nil {
			break
		}
		if u64Temp != uint64(i) {
			t.Fatal()
		}
	}
}

func Test_PushPopU16String(t *testing.T) {
	var byBuf []byte
	var byPush []byte
	var byPop []byte
	var strValue string
	byBuf = make([]byte, 30)
	byPush = byBuf

	strTemp0 := string("我abc+23")
	byPush = PushU16String(byPush, strTemp0)
	if byPush == nil {
		t.Fatal()
	}

	byPop = byBuf
	byPop = PopU16String(byPop, &strValue)
	if strValue != strTemp0 {
		t.Fatal()
	}

	strTemp1 := string("11111")
	byPush = PushU16String(byPush, strTemp1)
	if byPush == nil {
		t.Fatal()
	}

	byPop = byBuf
	byPop = PopU16String(byPop, &strValue)
	if strValue != strTemp0 {
		t.Fatal()
	}
	byPop = PopU16String(byPop, &strValue)
	if strValue != strTemp1 {
		t.Fatal()
	}

	strTemp2 := string("我33fddf")
	byPush = PushU16String(byPush, strTemp2)
	if byPush == nil {
		t.Fatal()
	}

	byPop = byBuf
	byPop = PopU16String(byPop, &strValue)
	if strValue != strTemp0 {
		t.Fatal()
	}
	byPop = PopU16String(byPop, &strValue)
	if strValue != strTemp1 {
		t.Fatal()
	}
	byPop = PopU16String(byPop, &strValue)
	if strValue != strTemp2 {
		t.Fatal()
	}
}

func Test_PushPopU8String(t *testing.T) {
	var byBuf []byte
	var byPush []byte
	var byPop []byte
	var strValue string
	byBuf = make([]byte, 30)
	byPush = byBuf

	strTemp0 := string("我abc+23")
	byPush = PushU8String(byPush, strTemp0)
	if byPush == nil {
		t.Fatal()
	}

	byPop = byBuf
	byPop = PopU8String(byPop, &strValue)
	if strValue != strTemp0 {
		t.Fatal()
	}

	strTemp1 := string("11111")
	byPush = PushU8String(byPush, strTemp1)
	if byPush == nil {
		t.Fatal()
	}

	byPop = byBuf
	byPop = PopU8String(byPop, &strValue)
	if strValue != strTemp0 {
		t.Fatal()
	}
	byPop = PopU8String(byPop, &strValue)
	if strValue != strTemp1 {
		t.Fatal()
	}

	strTemp2 := string("我33fddf")
	byPush = PushU8String(byPush, strTemp2)
	if byPush == nil {
		t.Fatal()
	}

	byPop = byBuf
	byPop = PopU8String(byPop, &strValue)
	if strValue != strTemp0 {
		t.Fatal()
	}
	byPop = PopU8String(byPop, &strValue)
	if strValue != strTemp1 {
		t.Fatal()
	}
	byPop = PopU8String(byPop, &strValue)
	if strValue != strTemp2 {
		t.Fatal()
	}
}
