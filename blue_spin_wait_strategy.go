package disruptor

type BlueSpinWaitStrategy struct{}

// func WaitFor(cursor *Cursor) {
	// for cursor.Get() != cursor.GetBufferSize() {
	// 	time.Sleep(time.Nanosecond)
	// }
// }