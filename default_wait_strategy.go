package disruptor

import "time"

type DefaultWaitStrategy struct{}

type WaitStrategy interface {

}

const (
	BlockingWaitStrategy,					 
	SleepingWaitStrategy ,       	
	YieldingWaitStrategy  ,        
	BlueSpinWaitStrategy     ,      
)

func NewWaitStrategy() DefaultWaitStrategy        { return DefaultWaitStrategy{} }
func (this DefaultWaitStrategy) Gate(count int64) { time.Sleep(time.Nanosecond) }
func (this DefaultWaitStrategy) Idle(count int64) { time.Sleep(time.Microsecond * 50) }


func (this BlockingWaitStrategy) waitFor() {}
func (this SleepingWaitStrategy) waitFor() {}
func (this YieldingWaitStrategy) waitFor() {}
func (this BlueSpinStrategy) waitFor() {}

