package heappq

import "container/heap"

/*
You are given an array of CPU tasks tasks, where tasks[i] is an uppercase english character from A to Z. You are also given an integer n.

Each CPU cycle allows the completion of a single task, and tasks may be completed in any order.

The only constraint is that identical tasks must be separated by at least n CPU cycles, to cooldown the CPU.

Return the minimum number of CPU cycles required to complete all tasks.

Example 1:

Input: tasks = ["X","X","Y","Y"], n = 2

Output: 5
Explanation: A possible sequence is: X -> Y -> idle -> X -> Y.

Example 2:

Input: tasks = ["A","A","A","B","C"], n = 3

Output: 9
Explanation: A possible sequence is: A -> B -> C -> Idle -> A -> Idle -> Idle -> Idle -> A.

Constraints:

1 <= tasks.length <= 1000
0 <= n <= 100
*/

type Task struct {
	freq  int
	sched int
}

type CooldownQueue []*Task

func leastInterval(tasks []byte, n int) int {
	// prioritise the most frequent tasks since they have the most cooldown cycles
	// which gives us more time to do other tasks
	freqs := map[byte]int{}
	for _, task := range tasks {
		freqs[task]++
	}

	h := &IntHeapMax{}
	heap.Init(h)
	for _, freq := range freqs {
		heap.Push(h, freq)
	}

	// have a cooldown queue to store the tasks that are on cooldown.
	// once the cooldown is over, we can add the task back to the heap.
	q := CooldownQueue{}

	cycles := 0
	for h.Len() > 0 || len(q) > 0 {
		cycles++

		if h.Len() == 0 {
			// if the heap is empty but there are still tasks in the cooldown queue,
			// fast-forward to the next task that is ready to be executed. The first
			// task in the cooldown queue is guaranteed to be the next task to be executed
			// since it is the earliest scheduled task.
			cycles = q[0].sched
		} else {
			// otherwise, we prioritise the most frequent task so far
			// and add it to the cooldown queue if there are more tasks to execute
			freq := (heap.Pop(h)).(int)
			freq--
			if freq > 0 {
				task := &Task{
					freq:  freq,
					sched: cycles + n,
				}
				q = append(q, task)
			}
		}

		// if the cooldown for a task is over, we can add it back to the heap.
		if len(q) > 0 && q[0].sched == cycles {
			heap.Push(h, q[0].freq)
			q = q[1:]
		}
	}
	return cycles
}
