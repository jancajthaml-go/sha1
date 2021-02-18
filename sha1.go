package sha1

const (
	BLOCK_SIZE  int = 16
	DIGEST_SIZE int = 20
)

func bytes_to_words(bytes []byte) []uint {
	var (
		i     uint
		l     uint   = uint(len(bytes) * 8)
		words []uint = make([]uint, 1+((l+64>>9)<<4)+15)
	)
	for _, b := range bytes {
		words[i>>5] |= (uint(b) << (24 - (i - (i/32)*32)))
		i += 8
	}
	return words
}

func words_to_bytes(words []uint) []byte {
	var (
		i int
		j int    = len(words) * 32
		r []byte = make([]byte, len(words) * 4)
	)
	for b := 0; b < j; b += 8 {
		r[i] = byte((words[b>>5] >> uint(24-(b-(b/32)*32))) & 0xFF)
		i++
	}
	return r
}

// Digest returns SHA1 digest of given byte slice
func Digest(message []byte) []byte {
	var (
		l uint   = uint(len(message) * 8)
		m []uint = bytes_to_words(message)
		w []uint = make([]uint, 80)
		// FIXME problem with int vs uint here
		H0 int = 1732584193
		H1 int = -271733879
		H2 int = -1732584194
		H3 int = 271733878
		H4 int = -1009589776
		i  int
	)

	m[l>>5] |= 0x80 << (24 - (l - (l/32) * 32))
	m[((l+64>>9)<<4)+15] = l

loop:
	var (
		a int = H0
		b int = H1
		c int = H2
		d int = H3
		e int = H4
	)

	for j := 0; j < 80; j++ {

		if j < 16 {
			w[j] = m[i+j]
		} else {
			n := w[j-3] ^ w[j-8] ^ w[j-14] ^ w[j-16]
			w[j] = (n << 1) | (n >> 31)
		}

		var h int

		if j < 20 {
			h = (H1&H2 | ^H1&H3) + 1518500249
		} else if j < 40 {
			h = (H1 ^ H2 ^ H3) + 1859775393
		} else if j < 60 {
			h = (H1&H2 | H1&H3 | H2&H3) - 1894007588
		} else {
			h = (H1 ^ H2 ^ H3) - 899497514
		}

		H4 = H3
		H3 = H2
		H2 = (H1 << 30) | (H1 >> 2)
		H1 = H0
		H0 = ((H0 << 5) | (H0 >> 27)) + H4 + int(w[j]>>0) + h
	}

	H0 += a
	H1 += b
	H2 += c
	H3 += d
	H4 += e

	i += 16

	if i < len(m) {
		goto loop
	}

	r := make([]uint, 5)
	r[0] = uint(H0)
	r[1] = uint(H1)
	r[2] = uint(H2)
	r[3] = uint(H3)
	r[4] = uint(H4)
	return words_to_bytes(r)
}
