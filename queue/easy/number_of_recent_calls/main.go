package main

import "fmt"

type RecentCounter struct {
	requests []int
	ranges   []int
}

func NewRecentCounter() RecentCounter {
	return RecentCounter{
		requests: make([]int, 0),
		ranges:   make([]int, 0),
	}
}

func (r *RecentCounter) ping(t int) int {
	r.requests = append(r.requests, t)
	if len(r.ranges) == 0 {
		r.ranges = append(r.ranges, t-3000)
	}
	for i := 0; i < len(r.requests); i++ {
		if r.requests[i] < r.ranges[0] {
			r.ranges[0] = r.requests[i]
		}
	}
	for i := 0; i < len(r.requests); i++ {
		if r.requests[i] > t-3000 {
			r.ranges[0] = r.requests[i]
			break
		}
	}
	return len(r.requests) - len(r.ranges)
}

func main() {
	r := NewRecentCounter()
	fmt.Println(r.ping(1))
	fmt.Println(r.ping(100))
	fmt.Println(r.ping(3001))
	fmt.Println(r.ping(3002))
	// Output: 3
	// Explanation:
	// RecentCounter.ping(1) returns 1.
	// RecentCounter.ping(100) returns 2.
	// RecentCounter.ping(3001) returns 3.
	// RecentCounter.ping(3002) returns 3.
}
