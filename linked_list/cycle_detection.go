package linked_list

/*
Given the beginning of a linked list head, return true if there is a cycle in the linked list. Otherwise, return false.

There is a cycle in a linked list if at least one node in the list can be visited again by following the next pointer.

Internally, index determines the index of the beginning of the cycle, if it exists. The tail node of the list will set it's next pointer to the index-th node. If index = -1, then the tail node points to null and no cycle exists.

Note: index is not given to you as a parameter.

Example 1:



Input: head = [1,2,3,4], index = 1

Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:



Input: head = [1,2], index = -1

Output: false
Constraints:

1 <= Length of the list <= 1000.
-1000 <= Node.val <= 1000
index is -1 or a valid index in the linked list.
*/

func hasCycle(head *ListNode) bool {
	tortoise, hare := head, head
	for hare != nil && hare.Next != nil {
		hare = hare.Next.Next
		tortoise = tortoise.Next

		if hare == tortoise {
			return true
		}
	}
	return false
}

/*
After detecting a meeting point:

To find cycle start: reset slow to head and move both slow and fast one step at a time; the node where they meet is the first node in the cycle.

To find cycle length: keep one pointer fixed at the meeting point and move the other around the cycle counting steps until it returns to that node.

Intuition:

To visualize this, imagine the Linked List looks like a lollipop or the letter "P".

There is a straight stem (the list before the cycle) and a circle (the cycle).

The Visual Map

Let’s break the track down into three distinct segments using distances:

x: The length of the "stem" (from Head to the Cycle Start).

y: The distance inside the loop from Cycle Start to the Meeting Point (where the Hare caught the Tortoise).

z: The remaining distance in the loop (from the Meeting Point back to the Cycle Start).

Visual check:

The length of the full cycle L is y + z.

If you are at the Meeting Point and walk distance z, you arrive back at the Cycle Start.

The Math (Simplified)

When they meet, the Tortoise has traveled the stem plus part of the loop:

Distance_Tortoise = x + y

The Hare travels twice as fast, so he has gone twice the distance:

Distance_Hare = 2(x + y)

However, the Hare didn’t just go in a straight line. He went through the stem (x), the partial loop (y), and some number of full laps inside the circle before catching the Tortoise. Let’s say he did k full laps:

Distance_Hare = x + y + kL

Now we equate the two versions of the Hare’s distance:

2(x + y) = x + y + kL

The "Aha!" Algebra

We want to prove that the stem (x) is equal to the remainder of the loop (z). Let’s solve the equation above for x.

Expand the left side:
2x + 2y = x + y + kL

Subtract x and y from both sides:
x + y = kL

Isolate x:
x = kL - y

Decoding the Result Visually

The equation x = kL - y is the key. Let’s translate it into plain English:

kL means “some number of full circles”.

-y means “backtrack distance y”.

So mathematically, length x is the same as “spinning around the circle k times, then backing up distance y”.

Look at the diagram again:

If you are at the Cycle Start and you go around the circle (length L) but back up by distance y, where are you?
You are at the Meeting Point.

Conversely:

If you are at the Meeting Point and you want to finish the circle to get back to the Cycle Start, how far do you walk?
You walk distance z.

Since L = y + z, then kL - y is mathematically equivalent to just z (plus possibly some extra full laps).
*/
