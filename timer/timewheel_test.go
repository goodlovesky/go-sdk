package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	tw := DefaultTimeWheel()
	err := tw.AddTask("task-1", func(key string) {
		fmt.Println("task run :", key, " > ", time.Now().Format(time.DateTime))
	}, 1*time.Second)
	t.Log("err", err)

	time.Sleep(60 * time.Second)
	tw.RemoveTask("task-1")

	time.Sleep(10 * time.Second)
	tw.Stop()
}
