package mysync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		var wantedCount int64 = 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(int(wantedCount))

		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int64) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
