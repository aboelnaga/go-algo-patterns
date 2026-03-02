# DSA Learning Roadmap — Go

## How to Use This File

- Check off `[ ]` when you attempt a problem
- Mark level achieved after each session: `[B]` Brute | `[O]` Optimal | `[M]` Most Optimized
- Add the date you first solved it and the date you last reviewed it
- One problem at a time — do not jump ahead
- Follow the coaching flow: approach verbally first → pseudocode → code → edge cases → complexity → improve

## Coaching Flow (Per Problem — Google Interview Style)

1. Read the problem aloud and restate it in your own words
2. Ask clarifying questions (input types, constraints, edge cases)
3. Think of the brute force solution and **state it before coding**
4. Code the brute force, add comments explaining each step
5. Analyze time and space complexity of your brute force
6. Identify the bottleneck — what makes it slow?
7. Propose an optimal approach verbally before implementing
8. Implement the optimal solution with comments
9. Run tests, cover edge cases
10. Mark levels achieved and schedule for spaced repetition review

---

## Phase 0 — Foundations

> Goal: Build the mental models you need before touching LeetCode. These are not coding problems — they are thinking exercises. Do not skip this phase.

### 0.1 Big O Notation

- [x] **Concept: Time Complexity Basics** — Easy
  - Focus: What does O(n), O(log n), O(n²) mean visually and intuitively?
  - Exercise: Given a for loop nested inside another, derive the complexity. No code, just reasoning.
  - Visual: https://www.bigocheatsheet.com/ — study the chart at the top
  - Levels: N/A (concept study)
  - First studied: 2026-02-21 | Last reviewed: _____

- [x] **Concept: Space Complexity Basics** — Easy
  - Focus: What memory does your program use? Stack vs heap, input space vs auxiliary space.
  - Exercise: Trace through a recursive function and count stack frames.
  - Levels: N/A (concept study)
  - First studied: 2026-02-23 | Last reviewed: _____

- [x] **Concept: Best / Average / Worst Case** — Easy
  - Focus: Why we almost always discuss worst case. When average case matters.
  - Exercise: Analyze linear search in all three cases.
  - Levels: N/A (concept study)
  - First studied: 2026-02-23 | Last reviewed: _____

- [x] **Concept: Comparing Complexities** — Easy
  - Focus: Ordering — O(1) < O(log n) < O(n) < O(n log n) < O(n²) < O(2ⁿ) < O(n!)
  - Exercise: For each algorithm you write going forward, always ask: where is the bottleneck?
  - Levels: N/A (concept study)
  - First studied: 2026-02-23 | Last reviewed: _____
  - Notes: notes/phase0.1_big_o.md

### 0.2 Go Language Fundamentals

- [x] **Concept: Variables and Constants** — Easy
  - Focus: `var` vs `:=`, Zero Values, `const` and `iota`, scope and shadowing
  - Exercise: Declare variables using both styles, observe zero values, use `iota` to define a named set of constants (e.g., directions: North, South, East, West).
  - Levels: N/A (concept study)
  - First studied: 2026-02-28 | Last reviewed: _____

- [x] **Concept: Numeric Types and Booleans** — Easy
  - Focus: Signed integers (`int`, `int8/16/32/64`), unsigned (`uint`, `uint8/16/32/64`), `float32`/`float64`, `complex64`/`complex128`, `bool`, `rune` (alias for `int32`). Type conversions are always explicit in Go — never implicit.
  - Exercise: Try adding an `int` and an `int64` without conversion — observe the compile error. Then fix it with explicit conversion.
  - Levels: N/A (concept study)
  - First studied: 2026-02-28 | Last reviewed: _____

- [x] **Concept: Strings in Go** — Easy
  - Focus: Interpreted string literals (`"..."` — supports escape sequences like `\n`) vs raw string literals (`` `...` `` — no escaping, can span lines). Strings are immutable byte sequences. Iterating with `range` gives runes; indexing gives bytes. Type conversions: `string ↔ []byte`, `string ↔ []rune`.
  - Exercise: Iterate over a string containing an Arabic or emoji character — compare `range` (rune) vs index (byte) iteration.
  - Levels: N/A (concept study)
  - First studied: 2026-02-28 | Last reviewed: _____

- [x] **Concept: Functions in Go** — Easy
  - Focus: Function syntax, multiple return values (the Go way), named returns, variadic functions (`...`). Why multiple returns replace exceptions in Go.
  - Exercise: Write a `divide(a, b int) (int, error)` function that returns an error when b is zero.
  - Levels: N/A (concept study)
  - First studied: 2026-02-28 | Last reviewed: _____

- [x] **Concept: Control Flow in Go** — Easy
  - Focus: `for` is the only loop (no `while`, no `do-while`). Four forms: classic `for`, while-style `for`, infinite `for {}`, and `for range`. `if/else` with optional init statement. `switch` without fallthrough by default.
  - Exercise: Rewrite a classic `while` loop using Go's `for`. Use `for range` to iterate over a slice and a map.
  - Levels: N/A (concept study)
  - First studied: 2026-03-01 | Last reviewed: _____

### 0.3 Go Basics for DSA

- [x] **Concept: Slices vs Arrays in Go** — Easy
  - Focus: How Go slices work under the hood. Append, length, capacity. Why this matters for DSA.
  - Exercise: Write a function that doubles a slice without modifying the original.
  - File: `slice_basics.go` + `slice_basics_test.go`
  - Levels: [B] [x] [O] [x] [M] [x]
  - First solved: 2026-03-02 | Last reviewed: _____

- [ ] **Concept: Maps in Go** — Easy
  - Focus: Hash map internals (conceptual). How Go maps are declared, read, written, and checked for key existence.
  - Exercise: Count word frequency in a string using a map.
  - File: `map_basics.go` + `map_basics_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Concept: Structs and Pointers for DSA** — Easy
  - Focus: How to build linked list nodes, tree nodes. Why pointers matter.
  - Exercise: Define a `ListNode` struct and manually link three nodes together.
  - File: `struct_basics.go` + `struct_basics_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

### 0.4 Recursion Basics

- [ ] **Concept: The Call Stack** — Easy
  - Focus: How function calls stack up. Base case vs recursive case. What happens without a base case.
  - Exercise: Draw the call stack for `factorial(4)` by hand before writing any code.
  - Levels: N/A (concept study)
  - First studied: _____ | Last reviewed: _____

- [ ] **Problem: Fibonacci (Recursive → Memoized)** — Easy
  - Focus: Classic recursion. Identify the repeated subproblems — this foreshadows DP.
  - File: `fibonacci.go` + `fibonacci_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 1 — Arrays & Hashing

> **Pattern**: Use a hash map to trade time for space. When you need O(1) lookup, reach for a map. The question to always ask: "Have I seen this before?"

- [ ] **Contains Duplicate** — Easy (LC 217)
  - Why: First introduction to using a map as a "seen" set. Foundational pattern.
  - Focus: Brute is O(n²) comparison. Optimal is O(n) with a hash set.
  - File: `contains_duplicate.go` + `contains_duplicate_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Two Sum** — Easy (LC 1)
  - Why: The canonical hash map problem. Every interviewer knows it.
  - Focus: Brute is nested loops O(n²). Optimal is one-pass map O(n).
  - File: `two_sum.go` + `two_sum_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Valid Anagram** — Easy (LC 242)
  - Why: Character frequency counting. Same map-as-counter pattern.
  - Focus: Two maps vs one map vs sorted string comparison — discuss all three.
  - File: `valid_anagram.go` + `valid_anagram_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Group Anagrams** — Medium (LC 49)
  - Why: Extends anagram concept to grouping. Introduces sorted-string as map key.
  - Focus: How to choose a canonical key to represent a group.
  - File: `group_anagrams.go` + `group_anagrams_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Top K Frequent Elements** — Medium (LC 347)
  - Why: Combines frequency map with sorting or bucket sort. Bridge to heaps later.
  - Focus: Bucket sort approach achieves O(n) — a rare trick worth knowing.
  - File: `top_k_frequent.go` + `top_k_frequent_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Product of Array Except Self** — Medium (LC 238)
  - Why: Classic prefix/suffix product. No division allowed — forces creative thinking.
  - Focus: Two-pass prefix product then suffix product in O(n) time, O(1) extra space.
  - File: `product_except_self.go` + `product_except_self_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Valid Sudoku** — Medium (LC 36)
  - Why: Multi-dimensional hash set tracking. Index math to identify 3×3 boxes.
  - Focus: Encoding seen values with composite keys.
  - File: `valid_sudoku.go` + `valid_sudoku_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Longest Consecutive Sequence** — Medium (LC 128)
  - Why: Set membership + sequence building. Looks like O(n²) but is actually O(n).
  - Focus: Only start counting from the beginning of a sequence (no left neighbor in the set).
  - File: `longest_consecutive.go` + `longest_consecutive_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 2 — Two Pointers

> **Pattern**: When the array is sorted (or you need to find pairs/triplets), use two pointers moving toward or away from each other. Eliminates nested loops: O(n²) → O(n).

- [ ] **Valid Palindrome** — Easy (LC 125)
  - Why: Introduce the two-pointer pattern on strings. Simple and confidence-building.
  - Focus: Left pointer moves right, right pointer moves left, compare characters.
  - File: `valid_palindrome.go` + `valid_palindrome_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Two Sum II (Sorted Array)** — Medium (LC 167)
  - Why: Sorted input unlocks two-pointer. Connects back to Two Sum.
  - Focus: Why sorting lets you confidently move pointers based on sum comparison.
  - File: `two_sum_sorted.go` + `two_sum_sorted_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **3Sum** — Medium (LC 15)
  - Why: Combines sorting + two pointers + deduplication. A classic hard-feeling medium.
  - Focus: Fix one element, use two pointers for the remaining two. Handle duplicates carefully.
  - File: `three_sum.go` + `three_sum_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Container With Most Water** — Medium (LC 11)
  - Why: Greedy two-pointer. Teaches why moving the shorter side is always correct.
  - Focus: Proof — moving the taller side can only decrease or maintain area.
  - File: `container_with_water.go` + `container_with_water_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Trapping Rain Water** — Hard (LC 42)
  - Why: Classic hard problem. Two-pointer beats the prefix array approach in space.
  - Focus: At each index, water = min(max_left, max_right) - height[i]. Track maxes with pointers.
  - File: `trapping_rain_water.go` + `trapping_rain_water_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 3 — Sliding Window

> **Pattern**: Maintain a window of elements with two pointers (left, right). Expand right to grow, shrink left when a constraint is violated. Key question: "What is the window's invariant?"

- [ ] **Best Time to Buy and Sell Stock** — Easy (LC 121)
  - Why: Introduces the sliding window mindset. Track min so far, compute max profit.
  - Focus: Why we never need to look back past the current minimum.
  - File: `buy_sell_stock.go` + `buy_sell_stock_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Longest Substring Without Repeating Characters** — Medium (LC 3)
  - Why: Dynamic window with a set. Expand right, shrink left on repeat.
  - Focus: Use a set to track window contents. When duplicate found, shrink until it's gone.
  - File: `longest_substring.go` + `longest_substring_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Longest Repeating Character Replacement** — Medium (LC 424)
  - Why: Window with a frequency count and a key insight about max frequency.
  - Focus: Window is valid when `(window_size - max_freq) <= k`. Why this works.
  - File: `char_replacement.go` + `char_replacement_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Permutation in String** — Medium (LC 567)
  - Why: Fixed-size sliding window. Frequency map comparison.
  - Focus: Slide a window of size `len(s1)` across `s2`, compare character counts.
  - File: `permutation_in_string.go` + `permutation_in_string_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Minimum Window Substring** — Hard (LC 76)
  - Why: Variable window with character requirement tracking. A top-tier interview problem.
  - Focus: Track how many characters from `t` are "satisfied" in the window. Two-pointer shrink.
  - File: `min_window_substring.go` + `min_window_substring_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Sliding Window Maximum** — Hard (LC 239)
  - Why: Monotonic deque. Teaches a data structure trick useful in many problems.
  - Focus: Maintain a deque of indices in decreasing order of values. Front is always the max.
  - File: `sliding_window_max.go` + `sliding_window_max_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 4 — Stack

> **Pattern**: Use a stack when you need to process elements and revisit the most recent one (LIFO). Great for matching, nesting, monotonic order, and "next greater element" problems.

- [ ] **Valid Parentheses** — Easy (LC 20)
  - Why: The canonical stack problem. Push open brackets, pop and match on close.
  - Focus: What to do when the stack is empty on a close bracket. Edge cases matter.
  - File: `valid_parentheses.go` + `valid_parentheses_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Min Stack** — Medium (LC 155)
  - Why: Design problem. How to track auxiliary state (the minimum) alongside the stack.
  - Focus: Use a parallel stack that stores the current min at each level.
  - File: `min_stack.go` + `min_stack_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Evaluate Reverse Polish Notation** — Medium (LC 150)
  - Why: Stack for expression evaluation. Push operands, operator pops two and pushes result.
  - Focus: Clean separation of operand and operator cases.
  - File: `reverse_polish.go` + `reverse_polish_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Daily Temperatures** — Medium (LC 739)
  - Why: Monotonic stack. Stack of indices waiting for a warmer day.
  - Focus: Maintain a decreasing stack. When a warmer day arrives, pop and record distance.
  - File: `daily_temperatures.go` + `daily_temperatures_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Car Fleet** — Medium (LC 853)
  - Why: Simulation with a stack. Sorting + logical grouping.
  - Focus: Sort by position, compute time to target, stack merges cars that catch up.
  - File: `car_fleet.go` + `car_fleet_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Largest Rectangle in Histogram** — Hard (LC 84)
  - Why: Monotonic stack at its hardest. Tests full mastery of the pattern.
  - Focus: For each bar, find how far left and right it can extend as the shortest bar.
  - File: `largest_rectangle.go` + `largest_rectangle_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 5 — Binary Search

> **Pattern**: When the search space is sorted (or can be treated as monotonic), binary search cuts it in half each step — O(log n) instead of O(n). Key insight: "Can I search on the *answer* instead of the array?"

- [ ] **Binary Search** — Easy (LC 704)
  - Why: The purest form. Nail the loop invariant and off-by-one errors first.
  - Focus: `left <= right`, `mid = left + (right-left)/2`, when to return -1.
  - File: `binary_search.go` + `binary_search_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Search a 2D Matrix** — Medium (LC 74)
  - Why: Treat the 2D matrix as a flat sorted array. Index math to convert mid to row/col.
  - Focus: `row = mid / cols`, `col = mid % cols`. Same binary search, different indexing.
  - File: `search_2d_matrix.go` + `search_2d_matrix_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Koko Eating Bananas** — Medium (LC 875)
  - Why: Binary search on the *answer space*, not the array. A huge conceptual leap.
  - Focus: "What is the minimum speed k such that Koko finishes?" — search on k, not the array.
  - File: `koko_bananas.go` + `koko_bananas_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Find Minimum in Rotated Sorted Array** — Medium (LC 153)
  - Why: Binary search on a modified sorted array. Identify which half is sorted.
  - Focus: Compare mid to right to determine which side the minimum is on.
  - File: `min_rotated_array.go` + `min_rotated_array_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Search in Rotated Sorted Array** — Medium (LC 33)
  - Why: Harder variation. Two conditions: which half is sorted, then where is target.
  - Focus: Always one half is fully sorted. Use that to decide where to search.
  - File: `search_rotated.go` + `search_rotated_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Time Based Key-Value Store** — Medium (LC 981)
  - Why: Binary search in a real design context. Builds engineering intuition.
  - Focus: Store values in a sorted list per key. Binary search for the largest timestamp ≤ t.
  - File: `time_key_value.go` + `time_key_value_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 6 — Linked Lists

> **Pattern**: Slow and fast pointers, dummy head nodes, and in-place reversal. Think in terms of pointer manipulation, not index access. Draw it out before coding.

- [ ] **Reverse Linked List** — Easy (LC 206)
  - Why: The foundational linked list operation. Everything builds on this.
  - Focus: Three pointers: prev, curr, next. Iterative first, then recursive.
  - File: `reverse_linked_list.go` + `reverse_linked_list_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Merge Two Sorted Lists** — Easy (LC 21)
  - Why: Dummy head pattern. Compare heads, advance the smaller one.
  - Focus: The dummy node eliminates special casing for the head.
  - File: `merge_sorted_lists.go` + `merge_sorted_lists_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Linked List Cycle** — Easy (LC 141)
  - Why: Floyd's tortoise and hare. First use of slow/fast pointers.
  - Focus: If there is a cycle, fast will eventually lap slow. If not, fast reaches nil.
  - File: `linked_list_cycle.go` + `linked_list_cycle_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Reorder List** — Medium (LC 143)
  - Why: Combines find-middle (slow/fast) + reverse second half + merge. Three sub-problems.
  - Focus: Break into clear steps. Each step is a problem you have already solved.
  - File: `reorder_list.go` + `reorder_list_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Remove Nth Node From End** — Medium (LC 19)
  - Why: Two-pointer trick on a linked list. Gap of n between fast and slow.
  - Focus: Advance fast n steps, then move both until fast reaches nil.
  - File: `remove_nth_node.go` + `remove_nth_node_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Copy List with Random Pointer** — Medium (LC 138)
  - Why: Map from old nodes to new nodes. Deep copy with non-trivial pointers.
  - Focus: Two passes: first create all new nodes, then wire up next and random.
  - File: `copy_random_list.go` + `copy_random_list_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Find the Duplicate Number** — Medium (LC 287)
  - Why: Array-as-linked-list + Floyd's cycle detection. Elegant reuse of a pattern.
  - Focus: Treat each value as a pointer. The duplicate creates a cycle.
  - File: `find_duplicate.go` + `find_duplicate_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 7 — Trees

> **Pattern**: DFS with recursion is the default. Ask: "What should this node return to its parent?" BFS with a queue handles level-by-level problems. Always handle the nil base case first.

- [ ] **Invert Binary Tree** — Easy (LC 226)
  - Why: First recursive tree problem. The recursive structure mirrors the tree structure.
  - Focus: Swap left and right, recurse on each. Trust the recursion.
  - File: `invert_tree.go` + `invert_tree_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Maximum Depth of Binary Tree** — Easy (LC 104)
  - Why: DFS returning a value upward. Each node's answer uses its children's answers.
  - Focus: Base case: nil returns 0. Each node returns `1 + max(left, right)`.
  - File: `max_depth_tree.go` + `max_depth_tree_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Same Tree** — Easy (LC 100)
  - Why: Structural comparison. Three base cases: both nil, one nil, values differ.
  - Focus: Clean recursive logic. Handle all base cases before recursive calls.
  - File: `same_tree.go` + `same_tree_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Subtree of Another Tree** — Easy (LC 572)
  - Why: Recursive search + isSameTree reuse. Composition of solutions.
  - Focus: At each node in the main tree, check if the subtree rooted there equals the target.
  - File: `subtree_of_tree.go` + `subtree_of_tree_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Diameter of Binary Tree** — Easy (LC 543)
  - Why: Return height, update a global answer. Pattern for problems where the answer isn't just returned.
  - Focus: The diameter through a node = left_height + right_height. Track max across all nodes.
  - File: `diameter_tree.go` + `diameter_tree_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Lowest Common Ancestor of BST** — Medium (LC 235)
  - Why: BST property guides the search. No full traversal needed.
  - Focus: Both values less than root → go left. Both greater → go right. Otherwise root is LCA.
  - File: `lca_bst.go` + `lca_bst_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Binary Tree Level Order Traversal** — Medium (LC 102)
  - Why: BFS with a queue. First use of level-by-level thinking.
  - Focus: At each level, process exactly `len(queue)` nodes, then add their children.
  - File: `level_order_traversal.go` + `level_order_traversal_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Validate Binary Search Tree** — Medium (LC 98)
  - Why: Pass bounds down the recursion. Common mistake: only comparing to the direct parent.
  - Focus: Each node must satisfy `min < node.val < max`. Pass tighter bounds to children.
  - File: `validate_bst.go` + `validate_bst_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Kth Smallest Element in BST** — Medium (LC 230)
  - Why: In-order traversal of a BST gives sorted order. Counting during traversal.
  - Focus: In-order DFS with a counter. Stop when k is reached.
  - File: `kth_smallest_bst.go` + `kth_smallest_bst_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Binary Tree Right Side View** — Medium (LC 199)
  - Why: BFS level-order, taking the last element at each level.
  - Focus: Combines level-order traversal with result extraction.
  - File: `right_side_view.go` + `right_side_view_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Binary Tree Maximum Path Sum** — Hard (LC 124)
  - Why: Harder version of diameter. Paths can have negative values — must decide to include or not.
  - Focus: Same pattern as diameter but with values. Clamp negative contributions to 0.
  - File: `max_path_sum.go` + `max_path_sum_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Serialize and Deserialize Binary Tree** — Hard (LC 297)
  - Why: Design + BFS/DFS. Encodes the full structure including nulls.
  - Focus: Use pre-order traversal, mark nulls. Reconstruct by consuming the same stream.
  - File: `serialize_deserialize.go` + `serialize_deserialize_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 8 — Heap / Priority Queue

> **Pattern**: When you need the k-th largest/smallest, or a running median, use a heap. Go's `container/heap` package requires implementing an interface — learn it once.

- [ ] **Concept: Go Heap Interface** — Easy
  - Focus: Go has no built-in heap type. You must implement `Len`, `Less`, `Swap`, `Push`, `Pop`.
  - Exercise: Write a min-heap and a max-heap from scratch using `container/heap`.
  - File: `heap_practice.go` + `heap_practice_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Kth Largest Element in a Stream** — Easy (LC 703)
  - Why: Maintain a min-heap of size k. The root is always the kth largest.
  - Focus: If heap grows beyond k, pop the minimum. Root = kth largest.
  - File: `kth_largest_stream.go` + `kth_largest_stream_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Last Stone Weight** — Easy (LC 1046)
  - Why: Max-heap simulation. Good first use of a heap for a simulation problem.
  - Focus: Always smash the two largest stones. Max-heap gives O(log n) per operation.
  - File: `last_stone_weight.go` + `last_stone_weight_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **K Closest Points to Origin** — Medium (LC 973)
  - Why: Max-heap of size k to track k closest. Or sort — compare the tradeoffs.
  - Focus: Maintain a max-heap. When size exceeds k, pop the farthest.
  - File: `k_closest_points.go` + `k_closest_points_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Kth Largest Element in an Array** — Medium (LC 215)
  - Why: Quickselect O(n) avg vs heap O(n log k). A great tradeoff discussion.
  - Focus: Heap approach is simpler and predictable. Quickselect is faster on average.
  - File: `kth_largest_array.go` + `kth_largest_array_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Task Scheduler** — Medium (LC 621)
  - Why: Greedy + max-heap simulation. Idle time calculation.
  - Focus: Always schedule the most frequent remaining task. Use a cooldown queue.
  - File: `task_scheduler.go` + `task_scheduler_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Find Median from Data Stream** — Hard (LC 295)
  - Why: Two heaps (max-heap for lower half, min-heap for upper half). Advanced design.
  - Focus: Maintain balance between heaps. Median comes from the tops of both heaps.
  - File: `median_data_stream.go` + `median_data_stream_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 9 — Backtracking

> **Pattern**: Explore all possibilities recursively. At each step: make a choice, recurse, undo the choice. Think of it as DFS on a decision tree. Template: Choose → Recurse → Unchoose.

- [ ] **Subsets** — Medium (LC 78)
  - Why: First backtracking problem. Include or exclude each element.
  - Focus: Binary decision at each element. No pruning needed here — pure exploration.
  - File: `subsets.go` + `subsets_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Combination Sum** — Medium (LC 39)
  - Why: Backtracking with repetition allowed. Prune when running sum exceeds target.
  - Focus: Start index prevents revisiting earlier elements (avoids duplicate combos).
  - File: `combination_sum.go` + `combination_sum_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Permutations** — Medium (LC 46)
  - Why: Every element can go in every position. Use a "used" array to track.
  - Focus: Unlike subsets, order matters. Try all unused elements at each level.
  - File: `permutations.go` + `permutations_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Subsets II (with duplicates)** — Medium (LC 90)
  - Why: Deduplication in backtracking. Sort first, skip duplicate branches.
  - Focus: After sorting, if current element equals previous and previous was not used at this level, skip.
  - File: `subsets_ii.go` + `subsets_ii_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Word Search** — Medium (LC 79)
  - Why: Backtracking on a 2D grid. Mark visited, try all four directions, unmark on backtrack.
  - Focus: In-place marking (set cell to `'#'`, restore after). Avoids extra visited array.
  - File: `word_search.go` + `word_search_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Palindrome Partitioning** — Medium (LC 131)
  - Why: Backtracking + palindrome check. Builds full partitions.
  - Focus: At each index, try all valid palindrome substrings starting there, then recurse.
  - File: `palindrome_partition.go` + `palindrome_partition_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **N-Queens** — Hard (LC 51)
  - Why: The classic backtracking hard problem. Constraint propagation + careful pruning.
  - Focus: Track which columns, diagonals, and anti-diagonals are occupied. Prune immediately.
  - File: `n_queens.go` + `n_queens_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 10 — Graphs

> **Pattern**: Most graph problems are DFS or BFS. Always maintain a visited set. For shortest path, BFS is your default. For connectivity/components, DFS or Union-Find. Draw the graph before coding.

- [ ] **Concept: Graph Representations in Go** — Easy
  - Focus: Adjacency list (`map[int][]int`) vs adjacency matrix. How to build from an edge list.
  - Exercise: Build a graph from a list of edges and traverse it with DFS and BFS.
  - Levels: N/A (concept study)
  - First studied: _____ | Last reviewed: _____

- [ ] **Number of Islands** — Medium (LC 200)
  - Why: DFS/BFS on a grid. Sink visited land to avoid re-visiting.
  - Focus: For each unvisited `'1'` cell, run DFS to mark the whole island, increment count.
  - File: `number_of_islands.go` + `number_of_islands_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Clone Graph** — Medium (LC 133)
  - Why: Graph traversal + node mapping. DFS with a visited map from old to new nodes.
  - Focus: Map old nodes to clones. When you encounter a visited node, return its clone.
  - File: `clone_graph.go` + `clone_graph_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Max Area of Island** — Medium (LC 695)
  - Why: DFS returning a size. Variation of number of islands.
  - Focus: DFS returns the number of cells it sinks. Take the max.
  - File: `max_area_island.go` + `max_area_island_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Rotting Oranges** — Medium (LC 994)
  - Why: Multi-source BFS. Start BFS from all rotten oranges simultaneously.
  - Focus: Add all initially rotten oranges to the queue. BFS spreads rot layer by layer.
  - File: `rotting_oranges.go` + `rotting_oranges_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Pacific Atlantic Water Flow** — Medium (LC 417)
  - Why: Reverse BFS from both coasts. Meeting cells can reach both oceans.
  - Focus: Instead of simulating flow forward, reverse: BFS inward from ocean borders.
  - File: `pacific_atlantic.go` + `pacific_atlantic_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Course Schedule** — Medium (LC 207)
  - Why: Cycle detection in a directed graph. DFS with a "visiting" state.
  - Focus: Three states: unvisited, visiting (in current path), visited. Cycle if you reach "visiting."
  - File: `course_schedule.go` + `course_schedule_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Number of Connected Components** — Medium (LC 323)
  - Why: Union-Find in its most direct form. Good introduction to the pattern.
  - Focus: Union each edge. Count distinct roots in the final parent array.
  - File: `connected_components.go` + `connected_components_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Word Ladder** — Hard (LC 127)
  - Why: BFS for shortest path in an implicit graph. Each word is a node.
  - Focus: Each one-letter transformation is an edge. BFS gives minimum transformations.
  - File: `word_ladder.go` + `word_ladder_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 11 — Dynamic Programming 1D

> **Pattern**: Break the problem into overlapping subproblems. Define `dp[i]` clearly in plain English first. Find the recurrence relation. Fill bottom-up. Ask: "What decision is made at each step?"

- [ ] **Concept: Memoization vs Tabulation** — Easy
  - Focus: Two ways to implement DP. Top-down with cache vs bottom-up with table.
  - Exercise: Convert a recursive Fibonacci to memoized, then to tabulated.
  - Levels: N/A (concept study)
  - First studied: _____ | Last reviewed: _____

- [ ] **Climbing Stairs** — Easy (LC 70)
  - Why: The Fibonacci pattern in DP form. Base case + recurrence.
  - Focus: `dp[i] = dp[i-1] + dp[i-2]`. Optimize space to two variables.
  - File: `climbing_stairs.go` + `climbing_stairs_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Min Cost Climbing Stairs** — Easy (LC 746)
  - Why: Adds a cost to the Fibonacci pattern. Choice matters.
  - Focus: `dp[i] = min(dp[i-1], dp[i-2]) + cost[i]`.
  - File: `min_cost_stairs.go` + `min_cost_stairs_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **House Robber** — Medium (LC 198)
  - Why: Classic 1D DP. Cannot rob adjacent houses.
  - Focus: `dp[i] = max(dp[i-1], dp[i-2] + nums[i])`.
  - File: `house_robber.go` + `house_robber_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **House Robber II (Circular)** — Medium (LC 213)
  - Why: Circular array breaks the simple recurrence. Split into two linear passes.
  - Focus: Run house robber on `[0, n-2]` and `[1, n-1]`, take the max.
  - File: `house_robber_ii.go` + `house_robber_ii_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Longest Palindromic Substring** — Medium (LC 5)
  - Why: Expand-around-center approach — elegant O(n²) time, O(1) space.
  - Focus: Expand around each center (odd and even length palindromes).
  - File: `longest_palindromic.go` + `longest_palindromic_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Palindromic Substrings** — Medium (LC 647)
  - Why: Count palindromes using the same expand-around-center technique.
  - Focus: Same as above but count every valid palindrome found.
  - File: `palindromic_substrings.go` + `palindromic_substrings_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Decode Ways** — Medium (LC 91)
  - Why: String DP with validity conditions. Two choices at each step.
  - Focus: `dp[i]` depends on `dp[i-1]` (valid 1-digit) and `dp[i-2]` (valid 2-digit).
  - File: `decode_ways.go` + `decode_ways_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Coin Change** — Medium (LC 322)
  - Why: Classic unbounded knapsack. `dp[amount]` = min coins to make that amount.
  - Focus: `dp[i] = min over all coins c: dp[i - c] + 1`. Initialize to infinity.
  - File: `coin_change.go` + `coin_change_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Maximum Product Subarray** — Medium (LC 152)
  - Why: Must track both max and min (negatives flip the sign). Two states.
  - Focus: `curMax = max(num, curMax*num, curMin*num)`. Same for curMin.
  - File: `max_product_subarray.go` + `max_product_subarray_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Word Break** — Medium (LC 139)
  - Why: DP on a string. `dp[i]` = true if `s[0:i]` can be segmented into valid words.
  - Focus: For each i, try all j < i where `dp[j]` is true and `s[j:i]` is in the dictionary.
  - File: `word_break.go` + `word_break_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Longest Increasing Subsequence** — Medium (LC 300)
  - Why: O(n²) DP vs O(n log n) patience sorting. Study both.
  - Focus: `dp[i]` = length of LIS ending at index i. For each j < i where `nums[j] < nums[i]`, update.
  - File: `longest_increasing_subseq.go` + `longest_increasing_subseq_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Phase 12 — Dynamic Programming 2D

> **Pattern**: Use a 2D table when the problem depends on two inputs (two strings, a string and a target, a grid row/col). Define `dp[i][j]` clearly in plain English first. Think about what the boundaries mean.

- [ ] **Unique Paths** — Medium (LC 62)
  - Why: Grid DP introduction. Only two directions. `dp[i][j] = dp[i-1][j] + dp[i][j-1]`.
  - Focus: Initialize first row and column to 1. Then optimize space to a 1D array.
  - File: `unique_paths.go` + `unique_paths_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Longest Common Subsequence** — Medium (LC 1143)
  - Why: The textbook 2D DP problem. Foundation for edit distance and many others.
  - Focus: `dp[i][j]` = LCS of `s1[0:i]` and `s2[0:j]`. Match: `dp[i-1][j-1]+1`. No match: `max(skip either)`.
  - File: `longest_common_subseq.go` + `longest_common_subseq_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Best Time to Buy/Sell Stock with Cooldown** — Medium (LC 309)
  - Why: State machine DP. Three states: holding, sold, rest.
  - Focus: Define each state clearly. `held = max(held, rest-price)`, `sold = held+price`, `rest = max(rest, sold)`.
  - File: `stock_cooldown.go` + `stock_cooldown_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Coin Change II** — Medium (LC 518)
  - Why: Count combinations vs minimum coins. Loop order matters (combinations vs permutations).
  - Focus: `dp[i][j]` = ways to make amount j using first i coins. Outer loop on coins.
  - File: `coin_change_ii.go` + `coin_change_ii_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Target Sum** — Medium (LC 494)
  - Why: Knapsack with +/- assignment. Can be solved recursively or with DP.
  - Focus: Reduce to subset-sum: count subsets where sum = `(target + total) / 2`.
  - File: `target_sum.go` + `target_sum_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Edit Distance** — Medium (LC 72)
  - Why: Classic 2D DP on two strings. Insert, delete, replace.
  - Focus: `dp[i][j]` = edit distance between `s1[0:i]` and `s2[0:j]`. Three operations, three transitions.
  - File: `edit_distance.go` + `edit_distance_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Distinct Subsequences** — Hard (LC 115)
  - Why: Count ways, not just boolean. `dp[i][j]` = ways to form `t[0:j]` from `s[0:i]`.
  - Focus: If `s[i]==t[j]`, `dp[i][j] = dp[i-1][j-1] + dp[i-1][j]`. Else `dp[i-1][j]`.
  - File: `distinct_subsequences.go` + `distinct_subsequences_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

- [ ] **Burst Balloons** — Hard (LC 312)
  - Why: Interval DP. Think about which balloon to pop LAST in an interval, not first.
  - Focus: `dp[l][r]` = max coins from bursting all balloons in `(l, r)`. Try each k as last balloon.
  - File: `burst_balloons.go` + `burst_balloons_test.go`
  - Levels: [B] [ ] [O] [ ] [M] [ ]
  - First solved: _____ | Last reviewed: _____

---

## Spaced Repetition Review Tracker

> Schedule: review at 1 day, 3 days, 1 week, 2 weeks, 1 month after first solving. After each review, write a short note on what you forgot or found tricky.

| Problem | First Solved | Day 1 | Day 3 | Week 1 | Week 2 | Month 1 | Notes |
|---------|-------------|-------|-------|--------|--------|---------|-------|
| Two Sum | | | | | | | |
| Contains Duplicate | | | | | | | |
| Valid Palindrome | | | | | | | |
| Binary Search | | | | | | | |
| Reverse Linked List | | | | | | | |
| Climbing Stairs | | | | | | | |
| Number of Islands | | | | | | | |
| Valid Parentheses | | | | | | | |
| House Robber | | | | | | | |
| Longest Common Subsequence | | | | | | | |

> Add more rows as you progress. The goal is to solve the problem from scratch, not re-read your solution.

---

## Active Recall Prompts

> Before looking at your previous solution, answer these from memory:

- What is the brute force approach and why does it work?
- What is the key insight that leads to the optimal solution?
- Which data structure is central to the optimal solution, and why that one?
- What is the time complexity and space complexity of each approach?
- What are the edge cases? (empty input, single element, all duplicates, negatives, overflow)
- Can you code the optimal solution in under 15 minutes without looking?

---

## Visual Resources

| Resource | What It's For |
|----------|--------------|
| [NeetCode Roadmap](https://neetcode.io/roadmap) | Interactive visual of all topics and their dependencies |
| [NeetCode DSA for Beginners](https://neetcode.io/courses/dsa-for-beginners/0) | Start here before LeetCode |
| [Big-O Cheat Sheet](https://www.bigocheatsheet.com/) | Visual complexity chart + common algorithms |
| [roadmap.sh/datastructures-and-algorithms](https://roadmap.sh/datastructures-and-algorithms) | Step-by-step visual learning path |

---

## Progress Summary

| Phase | Total | Attempted | Brute Only | Optimal | Most Optimized |
|-------|-------|-----------|------------|---------|----------------|
| Phase 0 — Foundations | 14 | | | | |
| Phase 1 — Arrays & Hashing | 8 | | | | |
| Phase 2 — Two Pointers | 5 | | | | |
| Phase 3 — Sliding Window | 6 | | | | |
| Phase 4 — Stack | 6 | | | | |
| Phase 5 — Binary Search | 6 | | | | |
| Phase 6 — Linked Lists | 7 | | | | |
| Phase 7 — Trees | 12 | | | | |
| Phase 8 — Heap | 7 | | | | |
| Phase 9 — Backtracking | 7 | | | | |
| Phase 10 — Graphs | 9 | | | | |
| Phase 11 — DP 1D | 11 | | | | |
| Phase 12 — DP 2D | 8 | | | | |
| **Total** | **101** | | | | |

---

## Milestones

- [ ] Phase 0 complete — Mental models established, Big O is intuitive, Go basics solid
- [ ] Phase 1 complete — Comfortable reaching for a map as first instinct on lookup problems
- [ ] Phase 2 complete — Recognize sorted-input + pair problems and reach for two pointers
- [ ] Phase 3 complete — Can identify a sliding window problem by its structure within seconds
- [ ] Phase 4 complete — Know when LIFO order is the key insight to a problem
- [ ] Phase 5 complete — Can apply binary search to answer-space problems, not just arrays
- [ ] Phase 6 complete — Pointer manipulation feels natural; no index arithmetic needed
- [ ] Phase 7 complete — DFS recursion is fluent; BFS queue pattern is automatic
- [ ] Phase 8 complete — Can implement Go heap interface from memory; know when to use it
- [ ] Phase 9 complete — Backtracking template is internalized; can identify pruning opportunities
- [ ] Phase 10 complete — Graph traversal in any representation is comfortable
- [ ] Phase 11 complete — Can define a 1D DP recurrence independently from scratch
- [ ] Phase 12 complete — Ready for mock interviews across all major patterns
