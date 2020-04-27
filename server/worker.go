package server

import "github.com/moonD4rk/morph-scan/common/lib/goworker"

func taskToMasscan(queue, class, ipRange, scanPort, scanRate string, isPublic bool) error {
	return goworker.Enqueue(&goworker.Job{
		Queue: queue,
		Payload: goworker.Payload{
			Class: class,
			Args:  []interface{}{ipRange, scanPort, scanRate, isPublic},
		},
	})
}
