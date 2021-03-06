package values

// UInt64 does the best to convert the value whose key is k to uint64.
func (m SMap) UInt64(k string) (v uint64, err error) {
	_v, ok := m[k]
	if !ok {
		err = ErrNoKey
		return
	}
	return ToUInt64(_v)
}

// IsUInt64 returns true when the type of the value whose key is k is uint64;
// or false.
func (m SMap) IsUInt64(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(uint64)
	return ok
}

// Byte does the best to convert the value whose key is k to byte.
func (m SMap) Byte(k string) (v byte, err error) {
	return m.UInt8(k)
}

// IsByte returns true when the type of the value whose key is k is byte;
// or false.
func (m SMap) IsByte(k string) bool {
	return m.IsUInt8(k)
}

// Uintptr does the best to convert the value whose key is k to uintptr.
func (m SMap) Uintptr(k string) (v uintptr, err error) {
	_v, err := m.UInt(k)
	return uintptr(_v), err
}

// IsUintptr returns true when the type of the value whose key is k is uintptr;
// or false.
func (m SMap) IsUintptr(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(uintptr)
	return ok
}

// UInt does the best to convert the value whose key is k to uint.
func (m SMap) UInt(k string) (v uint, err error) {
	_v, err := m.UInt64(k)
	return uint(_v), err
}

// IsUInt returns true when the type of the value whose key is k is uint;
// or false.
func (m SMap) IsUInt(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(uint)
	return ok
}

// UInt8 does the best to convert the value whose key is k to uint8.
func (m SMap) UInt8(k string) (v uint8, err error) {
	_v, err := m.UInt64(k)
	return uint8(_v), err
}

// IsUInt8 returns true when the type of the value whose key is k is uint8;
// or false.
func (m SMap) IsUInt8(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(uint8)
	return ok
}

// UInt16 does the best to convert the value whose key is k to uint16.
func (m SMap) UInt16(k string) (v uint16, err error) {
	_v, err := m.UInt64(k)
	return uint16(_v), err
}

// IsUInt16 returns true when the type of the value whose key is k is uint16;
// or false.
func (m SMap) IsUInt16(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(uint16)
	return ok
}

// UInt32 does the best to convert the value whose key is k to uint32.
func (m SMap) UInt32(k string) (v uint32, err error) {
	_v, err := m.UInt64(k)
	return uint32(_v), err
}

// IsUInt32 returns true when the type of the value whose key is k is uint32;
// or false.
func (m SMap) IsUInt32(k string) bool {
	_v, ok := m[k]
	if !ok {
		return false
	}
	_, ok = _v.(uint32)
	return ok
}
