package binary

import gobinary "encoding/binary"
 

func PushU8(byBuf []byte, u8Value uint8) []byte {
	if len(byBuf) < 1 {
		return nil
	}
	byBuf[0] = u8Value
	return byBuf[1:]
}

func PopU8(byBuf []byte, u8Value *uint8) []byte {
	if len(byBuf) < 1 {
		return nil
	}
	*u8Value = byBuf[0]
	return byBuf[1:]
}

func PushU16(byBuf []byte, u16Value uint16) []byte {
	if len(byBuf) < 2 {
		return nil
	}
 
	gobinary.LittleEndian.PutUint16(byBuf, u16Value)
	return byBuf[2:]
}

func PopU16(byBuf []byte, u16Value *uint16) []byte {
	if len(byBuf) < 2 {
		return nil
	}

	*u16Value = gobinary.LittleEndian.Uint16(byBuf)
	return byBuf[2:]
}

func PushU32(byBuf []byte, u32Value uint32) []byte {
	if len(byBuf) < 4 {
		return nil
	}
	gobinary.LittleEndian.PutUint32(byBuf, u32Value)
	return byBuf[4:]
}

func PopU32(byBuf []byte, u32Value *uint32) []byte {
	if len(byBuf) < 4 {
		return nil
	}
	*u32Value = gobinary.LittleEndian.Uint32(byBuf)
	return byBuf[4:]
}

func PushU64(byBuf []byte, u64Value uint64) []byte {
	if len(byBuf) < 8 {
		return nil
	}
	gobinary.LittleEndian.PutUint64(byBuf, u64Value)
	return byBuf[8:]
}

func PopU64(byBuf []byte, u64Value *uint64) []byte {
	if len(byBuf) < 8 {
		return nil
	}
	*u64Value = gobinary.LittleEndian.Uint64(byBuf)
	return byBuf[8:]
}

func PushU8String(byBuf []byte, strValue string) []byte {
	var u8Len = len(strValue)
	var byValue = []byte(strValue)

	if len(byBuf) < u8Len {
		return nil
	}
	if uint32(u8Len) > uint32(0xff) {
		return nil
	}

	byBuf = PushU8(byBuf, uint8(u8Len))
	copy(byBuf, byValue)
	return byBuf[u8Len:]
}

func PopU8String(byBuf []byte, strValue *string) []byte {
	var u8Len uint8
	byBuf = PopU8(byBuf, &u8Len)
	if byBuf == nil {
		return nil
	}

	var byTemp []byte
	byTemp = make([]byte, u8Len)

	copy(byTemp, byBuf[:u8Len])
	*strValue = string(byTemp)
	return byBuf[u8Len:]
}

func PushU16String(byBuf []byte, strValue string) []byte {
	var u16Len = len(strValue)
	var byValue = []byte(strValue)

	if len(byBuf) < u16Len {
		return nil
	}
	if uint32(u16Len) > uint32(0xffff) {
		return nil
	}
 
	byBuf = PushU16(byBuf, uint16(u16Len))
	copy(byBuf, byValue)
	return byBuf[u16Len:]
}

func PopU16String(byBuf []byte, strValue *string) []byte {
	var u16Len uint16
	byBuf = PopU16(byBuf, &u16Len)
	if byBuf == nil {
		return nil
	}

	var byTemp []byte
	byTemp = make([]byte, u16Len)

	copy(byTemp, byBuf[:u16Len])
	*strValue = string(byTemp)
	return byBuf[u16Len:]
}
