package geecache

type ByteView struct {
	b []byte
}

func cloneBytes(bytes []byte) []byte {
	copyBytes := make([]byte, len(bytes))
	copy(copyBytes, bytes)
	return copyBytes
}

// 注意到 ByteView 的方法接收者都是对象 这样是为了不影响调用对象本身

func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice 返回一份[]byte的副本（深拷贝）
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}
