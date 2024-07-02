package buff

import "sync"

type ByteBuffer struct {
	data []byte
}

type ByteBufferPool struct {
	pool    sync.Pool
	maxSize int
}

func NewBufferPool(maxSize int) *ByteBufferPool {
	return &ByteBufferPool{
		pool: sync.Pool{
			New: func() any {
				return &ByteBuffer{
					data: make([]byte, 0, 100),
				}
			},
		},
		maxSize: maxSize,
	}
}

func (p *ByteBufferPool) Get() *ByteBuffer {
	return p.pool.Get().(*ByteBuffer)
}

func (p *ByteBufferPool) Put(b *ByteBuffer) {

	if cap(b.data) <= p.maxSize {
		b.data = b.data[:0]
		p.pool.Put(b)
	}

	// if it is larger then it is garbage collected
}

func (b *ByteBuffer) Write(d []byte) {
	b.data = append(b.data, d...)
}

func (b *ByteBuffer) Clear() {
	b.data = b.data[:0]
}

func (b *ByteBuffer) Bytes() []byte {
	return b.data
}
