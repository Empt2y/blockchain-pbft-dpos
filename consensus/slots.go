package consensus

import "time"

const (
	BlockInterval  = 10 // 产块时间间隔
	NumOfDelegates = 3  // 代理个数
)

var (
	InitTime = time.Date(2019, 4, 25, 0, 0, 0, 0, time.Local).Unix()
)

//
func getEpochNumber(t int64) int64 {
	return (t - InitTime) / BlockInterval
}
