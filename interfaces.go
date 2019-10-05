package disruptor

type Consumer interface {
	Consume(lower, upper int64)
}

type Barrier interface {
	Load() int64
}

type Writer interface {
	Reserve(count int64) int64
	Commit(lower, upper int64)
}

const (
	MaxSequenceValue     int64 = (1 << 63) - 1
	InitialSequenceValue int64 = -1
)