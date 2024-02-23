package evm

import "github.com/0xPolygon/polygon-edge/helper/common"

const bitmapSize = 8

type bitmap struct {
	buf []byte
}

func (b *bitmap) isSet(i uint64) bool {
	return b.buf[i/bitmapSize]&(1<<(i%bitmapSize)) != 0
}

func (b *bitmap) set(i uint64) {
	b.buf[i/bitmapSize] |= 1 << (i % bitmapSize)
}

func (b *bitmap) reset() {
	for i := range b.buf {
		b.buf[i] = 0
	}

	b.buf = b.buf[:0]
}

func (b *bitmap) setCode(code []byte) {
	codeSize := len(code)
	b.buf = common.ExtendByteSlice(b.buf, codeSize/bitmapSize+1)

	for i := 0; i < codeSize; {
		c := code[i]

		if isPushOp(c) {
			// push op
			pushLength := int(c) - int(PUSH1) + 2 // Calculate the length of the push instruction
			i += pushLength                         // Increment i by the length of the push instruction
		} else {
			if c == JUMPDEST {
				// jumpdest
				b.set(uint64(i))
			}
			i++
		}
	}
}


func isPushOp(i byte) bool {
	// From PUSH0 (0x60) to PUSH32(0x7F)
	return i >= PUSH0 && i <= PUSH32
}


