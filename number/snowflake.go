package number

import (
	"github.com/goodlovesky/go-sdk/snowflake"
	"time"
)

// 雪花算法
var (
	node *snowflake.Node
)

// InitMachine 初始化 机器占10位 最大值为1023
func InitMachine(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

// https://tech.meituan.com/2017/04/21/mt-leaf.html

func GenerateUID() int64 {
	return node.Generate().Int64()
}
