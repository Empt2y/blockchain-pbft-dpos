package db

// 理想情况下一次写入数据的大小
const IdealBatchSize = 100 * 1024

// Batch是一个只写数据库，当调用Write方法时，它将改动提交到数据库
type Batch interface {
	Writer

	// 检索排队等待写入的数据量
	ValueSize() int

	// 将所有累积的数据写到磁盘
	Write() error

	// 将Batch重置为可用
	Reset()

	// 重放batch的内容
	Replay(w Writer) error
}

// Batcher包装了NewBatch方法
type Batcher interface {
	// 创建一个Batch
	NewBatch() Batch
}
