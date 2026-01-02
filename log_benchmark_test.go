package teapot

import (
	"io"
	"testing"
)

func BenchmarkLogger_JSON(b *testing.B) {
	l := New(
		WithWriter(io.Discard), // does not measure os write performance
	)

	attrs := []Attr{
		String("user", "alice"),
		Int("status", 200),
		Bool("success", true),
	}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			l.Debug("received request", attrs...)
		}
	})
}
