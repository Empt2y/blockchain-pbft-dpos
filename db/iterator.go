package db

// Iterator按升序迭代数据库中的键值对
// 迭代器使用后必须释放
type Iterator interface {
	// 将迭代器移动到下一个键值对，返回迭代器是否有效
	Next() bool

	// 返回任何错误
	Error() error

	// 返回当前键值对的键
	Key() []byte

	// 返回当前键值对的值
	Value() []byte

	// 释放相关资源
	Release()
}

// Iteratee包装了NewIterator方法
type Iteratee interface {
	// 创建针对整个数据库的迭代器
	NewIterator() Iterator

	// 创建针对具有特定前缀的数据库子集的迭代器
	NewIteratorWithPrefix(prefix []byte) Iterator
}
