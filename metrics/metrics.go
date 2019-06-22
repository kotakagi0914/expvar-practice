package metrics

import (
	"expvar"
)

var (
	PracticeMap = expvar.NewMap("practice_map")

	Count1 = new(expvar.Int)
	Count2 = new(expvar.Int)
)

func InitMetrics() {
	PracticeMap.Set("count1", Count1)
	PracticeMap.Set("count2", Count2)
}
