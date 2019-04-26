package db

import "io"

// Reader包装了Has和Get方法
type Reader interface {
	// 检索key是否存在
	Has(key []byte) (bool, error)

	// 获取key对应的值
	Get(key []byte) ([]byte, error)
}

// Writer包装了Put和Delete方法
type Writer interface {
	// 插入键值对
	Put(key []byte, value []byte) error

	// 删除key对应的键值对
	Delete(key []byte) error
}

// Stater包装了Stat方法
type Stater interface {
	// 返回数据库的内部状态
	Stat(property string) (string, error)
}

// Compactor包装了Compact方法
type Compacter interface {
	// start和limit确定压缩范围
	// 如果start为nil，则start将视为所有key之前的一个key
	// 如果limit为nil，则limit将是为所有key之后的一个key
	// 如果start和limit都为nil，则会压缩整个数据集
	Compact(start []byte, limit []byte) error
}

// KeyValueStore包含了所有处理不同kv数据库的方法
type KeyValueStore interface {
	Reader
	Writer
	Batcher
	Iteratee
	Stater
	Compacter
	io.Closer
}
