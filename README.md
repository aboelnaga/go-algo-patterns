# go-algo-patterns

Structured DSA learning in Go — phase-based algorithmic patterns practiced through a Google interview-style coaching flow with spaced repetition.

---

## What this is

A personal learning repo for mastering data structures and algorithms in Go. Problems are solved in progression: brute force first, then optimal, then most optimized. Every session follows a coaching flow that mirrors how problems are approached in technical interviews.

---

## Structure

```
.
├── requirements.md          # Roadmap — source of truth for what to work on next
├── notes/                   # Concept notes per phase (written during learning)
├── *.go                     # One file per problem
└── *_test.go                # Table-driven tests per problem
```

Each problem has three levels:

| Level | Label | Goal |
|-------|-------|------|
| Brute Force | `[B]` | Correct logic, O(n²) or worse is fine |
| Optimal | `[O]` | Best known approach, pattern-based |
| Most Optimized | `[M]` | Clean, micro-optimized, edge cases covered |

---

## Phases

| Phase | Topic | Problems |
|-------|-------|----------|
| 0 | Foundations (Big O, Go fundamentals, recursion) | 14 |
| 1 | Arrays & Hashing | 8 |
| 2 | Two Pointers | 5 |
| 3 | Sliding Window | 6 |
| 4 | Stack | 6 |
| 5 | Binary Search | 6 |
| 6 | Linked Lists | 7 |
| 7 | Trees | 12 |
| 8 | Heap / Priority Queue | 7 |
| 9 | Backtracking | 7 |
| 10 | Graphs | 9 |
| 11 | Dynamic Programming 1D | 11 |
| 12 | Dynamic Programming 2D | 8 |
| **Total** | | **101** |

---

## Coaching Flow (Per Problem)

Each problem follows this sequence before any code is written:

1. Restate the problem in your own words
2. Ask clarifying questions — constraints, edge cases, input types
3. Describe the brute force approach verbally
4. Write pseudocode
5. Implement brute force with comments
6. Analyze time and space complexity
7. Identify the bottleneck
8. Propose the optimal approach verbally
9. Implement the optimal solution
10. Cover edge cases and discuss follow-ups

---

## Running Tests

```bash
# Run all tests
go test ./...

# Run a specific test
go test -v -run TestFunctionName
```

---

## Spaced Repetition

After solving a problem, reviews are scheduled at:
**1 day → 3 days → 1 week → 2 weeks → 1 month**

The tracker lives in `requirements.md`. Each review is done from scratch — no peeking at the previous solution.

---

## Progress

Track in [`requirements.md`](./requirements.md).
