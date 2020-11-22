package geecache

//存储真实的缓存值
type ByteView struct {
	b []byte
}

//实现Value接口的
func (v ByteView) Len() int {
	return len(v.b)
}

//返回的是拷贝
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)

	return c
}
