package timer

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Job func(key string)

type TimeWheel struct {
	interval     time.Duration
	slots        []*list.List
	slotsNum     int64
	currentSlots int64
	ticker       *time.Ticker
	lock         sync.Mutex
	running      bool
	tasks        sync.Map
	addTaskCh    chan *Task
	removeTaskCh chan string
	closeCh      chan struct{}
}

type Task struct {
	ID         string
	createTime time.Time
	delay      time.Duration
	slotIndex  int64
	circle     int64 // 多少圈
	job        Job
	times      int64 //执行多少次 -1 一直执行
}

func DefaultTimeWheel() *TimeWheel {
	tw, _ := NewTimeWheel(time.Second, 60*60*24)
	return tw
}

func NewTimeWheel(interval time.Duration, slotsNum int64) (*TimeWheel, error) {
	if interval < time.Second {
		return nil, errors.New("minimum interval is 1 second")
	}
	if slotsNum <= 0 {
		return nil, errors.New("slots num must be greater than 0")
	}
	tw := &TimeWheel{
		interval:     interval,
		slotsNum:     slotsNum,
		slots:        make([]*list.List, slotsNum),
		addTaskCh:    make(chan *Task),
		removeTaskCh: make(chan string),
		closeCh:      make(chan struct{}),
	}
	tw.start()
	return tw, nil
}

func (t *TimeWheel) start() {
	if !t.running {
		t.slots = make([]*list.List, t.slotsNum)
		for i := int64(0); i < t.slotsNum; i++ {
			t.slots[i] = list.New()
		}
		// 创建定时器
		t.ticker = time.NewTicker(t.interval)
		t.lock.Lock()
		t.running = true
		go t.run()
		t.lock.Unlock()
	}
}
func (t *TimeWheel) Stop() {
	if t.running {
		t.lock.Lock()
		t.running = false
		t.lock.Unlock()
		t.closeCh <- struct{}{}
	}
}

func (t *TimeWheel) AddTask(ID string, job Job, delay time.Duration, times ...int64) error {
	if ID == "" {
		return errors.New("ID is empty")
	}
	if delay < t.interval {
		return errors.New("the delay time must be greater than the interval time")
	}
	// 默认永久执行
	var timesInt64 int64 = -1
	if len(times) > 0 {
		timesInt64 = times[0]
	}
	_, ok := t.tasks.Load(ID)
	if ok {
		return errors.New("ID already exists")
	}
	task := &Task{
		ID:         ID,
		createTime: time.Now(),
		job:        job,
		delay:      delay,
		times:      timesInt64,
	}
	t.addTaskCh <- task
	return nil
}

func (t *TimeWheel) RemoveTask(ID string) error {
	_, ok := t.tasks.Load(ID)
	if !ok {
		return errors.New("ID does not exist")
	}
	t.removeTaskCh <- ID
	return nil
}

func (t *TimeWheel) addTask(task *Task, first bool) {
	task.circle, task.slotIndex = t.getCircleAndSlotIndex(task.delay, first)
	// 逻辑存储slots
	ele := t.slots[task.slotIndex].PushBack(task)
	// 缓存tasks用于记录任务
	t.tasks.Store(task.ID, ele)
}

func (t *TimeWheel) delTask(id string) {
	if val, ok := t.tasks.Load(id); ok {
		task := val.(*list.Element).Value.(*Task)
		t.slots[task.slotIndex].Remove(val.(*list.Element))
		t.tasks.Delete(task.ID)
	}
}
func (t *TimeWheel) run() {
	for {
		select {
		case _ = <-t.ticker.C:
			fmt.Println("触发定时任务")
			t.runTask()
		case task := <-t.addTaskCh:
			t.addTask(task, true)
		case id := <-t.removeTaskCh:
			t.delTask(id)
		case _ = <-t.closeCh:
			t.ticker.Stop()
			break
		}
	}
}

func (t *TimeWheel) runTask() {
	tasks := t.slots[t.currentSlots]
	if tasks != nil {
		for item := tasks.Front(); item != nil; item = item.Next() {
			task := item.Value.(*Task)
			if task.circle > 0 {
				task.circle--
				continue
			}
			go task.job(task.ID)
			t.tasks.Delete(task.ID)
			tasks.Remove(item)
			if task.times == -1 {
				t.addTask(task, false)
			} else {
				if task.times > 1 {
					task.times--
					t.addTask(task, false)
				} else {

				}
			}
		}
	}

	// 循环
	if t.currentSlots == t.currentSlots-1 {
		t.currentSlots = 0
	} else {
		t.currentSlots++
	}
}

func (t *TimeWheel) getInitSlots() int64 {
	return time.Now().Unix() % t.slotsNum
}

func (t *TimeWheel) getCircleAndSlotIndex(delay time.Duration, first bool) (circle, slotIndex int64) {

	delaySed := int64(delay.Seconds())
	intervalSed := int64(t.interval.Seconds())
	circle = delaySed / intervalSed / t.slotsNum
	//fmt.Println("circles:  " + strconv.FormatInt(circle, 10))
	slotIndex = delaySed - (t.slotsNum * intervalSed * circle) + t.currentSlots
	//fmt.Println("slotIndex:  " + strconv.FormatInt(slotIndex, 10))
	if slotIndex == t.currentSlots && circle > 0 {
		circle--
	}
	//第一次加入时 当前秒（currentSlots）还未执行，比如当前是第一秒的slot(0) 延迟5秒计算得出为5 （0～5有6格所有需要-1）
	//第二次加入时 当前秒（currentSlots）已经执行，就不需要-1
	if slotIndex > 0 && first {
		slotIndex--
	}
	return
}
