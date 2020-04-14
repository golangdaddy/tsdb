package metrick

type hashableUint64 uint64

func (h hashableUint64) Write([]byte) (int, error) {
	panic("Unimplemented")
}

func (h hashableUint64) Sum([]byte) []byte {
	panic("Unimplemented")
}

func (h hashableUint64) Reset() {
	panic("Unimplemented")
}

func (h hashableUint64) BlockSize() int {
	panic("Unimplemented")
}

func (h hashableUint64) Size() int {
	panic("Unimplemented")
}

func (h hashableUint64) Sum64() uint64 {
	return uint64(h)
}
