package agent

import (
	"github.com/moonD4rk/morph-scan/common/lib/go-masscan"
	"github.com/moonD4rk/morph-scan/common/lib/goworker"
)

func resultToNmapQueue(results []masscan.Result, queue, class string) error {
	pushList := make([]*goworker.Payload, len(results))
	for i, v := range results {
		pushList[i] = &goworker.Payload{
			Class: class,
			Args:  []interface{}{v.IP, v.Port, v.IsPublic},
		}
	}
	return goworker.EnqueueMoreOne(queue, pushList...)
}
