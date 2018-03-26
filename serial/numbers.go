package serial



// Uint16ToBytes serializes an uint16 into bytes in big endian order.
func Uint16ToBytes(value uint16) []byte {
	var b []byte
	return append(b, byte(value>>8), byte(value))
}


// Uint16ToBytes2  坤哥版本
func Uint16ToBytes2(value uint16) (b [2]byte){
	b[0] = byte(value >> 8 & 0xff)
	b[1] = byte(value & 0xff)
	return
}