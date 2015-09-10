package wolfang

import (
	"time"
	"fmt"
)

var Threads chan int

type Spider struct {
	taskName   string
	threadNum  int
	timeout    time.Duration
	scheduler  interface{}
	downloader interface{}
	processor  interface{}
	pipeliner  interface{}
}

func NewSpider(taskName string) *Spider {
	return &Spider{
		threadNum:2,
		taskName:taskName,
	}
}

func (self *Spider) SetTaskName(name string) {
	if len(name) < 1 {
		panic("task name should not be empty.")
	}else {
		self.taskName = name
	}
}

func (self *Spider) SetThreadNum(num int) {
	if num < 1 {
		panic("invalid thread number.")
	}else {
		self.threadNum = num
	}
}

func (self *Spider) SetTimeout(time time.Duration) {
	self.timeout = time
}

func (self *Spider) SetScheduler(scheduler interface{}) {
	self.scheduler = scheduler
}
func (self *Spider) SetDownloader(downloader interface{}) {
	self.downloader = downloader
}
func (self *Spider) SetProcessor(processor interface{}) {
	self.processor = processor
}
func (self *Spider) SetPipeliner(pipeliner interface{}) {
	self.pipeliner = pipeliner
}

func (self *Spider) Run() {
	fmt.Printf("%v", self)
}

