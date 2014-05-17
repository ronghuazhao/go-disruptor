package disruptor

import "testing"

func BenchmarkWriterPublish(b *testing.B) {
	writer := NewWriter(NewCursor(), 1024, &Barrier{})
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		writer.Publish(i)
	}
}

func BenchmarkWriterNext(b *testing.B) {
	readerCursor := NewCursor()
	writerCursor := NewCursor()
	readerBarrier := NewBarrier(readerCursor)

	writer := NewWriter(writerCursor, 1024, readerBarrier)
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	for i := int64(0); i < iterations; i++ {
		claimed := writer.Next(1)
		readerCursor.Store(claimed)
	}
}

func BenchmarkWriterNextWrapPoint(b *testing.B) {
	readerCursor := NewCursor()
	writerCursor := NewCursor()
	readerBarrier := NewBarrier(readerCursor)

	writer := NewWriter(writerCursor, 1024, readerBarrier)
	iterations := int64(b.N)
	b.ReportAllocs()
	b.ResetTimer()

	readerCursor.Store(MaxCursorValue)
	for i := int64(0); i < iterations; i++ {
		writer.Next(1)
	}
}
