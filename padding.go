package padding

func Pad(src []byte, blocksize int) []byte {
	if blocksize > 0xff {
		return nil
	}

	length := blocksize - (len(src) % blocksize)
	if length == 0 {
		length = blocksize
	}

	pad := make([]byte, length)
	for i := 0; i < length; i++ {
		pad[i] = byte(length)
	}

	return append(src, pad...)
}

func Unpad(src []byte, blocksize int) []byte {
	if blocksize > 0xff {
		return nil
	}

	if len(src) == 0 {
		return nil
	}

	length := int(src[len(src)-1])
	if length == 0 || length > blocksize || length > len(src) {
		return nil
	}

	for _, v := range src[len(src)-length:] {
		if int(v) != length {
			return nil
		}
	}

	return src[:len(src)-length]
}
