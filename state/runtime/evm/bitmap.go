package evm

import "github.com/0xPolygon/polygon-edge/helper/common"

const (
	bitmapSize = 8
	pushOpcode = 0x60
)

type bitmap struct {
	buf []byte
}

func (b *bitmap) isSet(i uint64) bool {
	return len(b.buf) > 0 && int(i/bitmapSize) < len(b.buf) && b.buf[i/bitmapSize]&(1<<(i%bitmapSize)) != 0
}

func (b *bitmap) set(i uint64) {
	idx := int(i / bitmapSize)
	if idx >= len(b.buf) {
		b.buf = append(b.buf, make([]byte, idx-len(b.buf)+1)...)
	}
	b.buf[idx] |= 1 << (i % bitmapSize)
}

func (b *bitmap) reset() {
	for i := range b.buf {
		b.buf[i] = 0
	}
	b.buf = b.buf[:0]
}

func (b *bitmap) setCode(code []byte) {
	b.reset()
	for i := 0; i < len(code); {
		c := code[i]
		if isPushOp(c) {
			size := int(c-pushOpcode) + 1
			if i+size >= len(code) {
				break // Avoid potential out-of-bounds access
			}
			for j := 1; j <= size; j++ {
				b.set(uint64(i + j))
			}
			i += size + 1
		} else {
			if c == JUMPDEST {
				b.set(uint64(i))
			}
			i++
		}
	}
}

func isPushOp(opcode byte) bool {
	return opcode >= pushOpcode && opcode < pushOpcode+32
}
