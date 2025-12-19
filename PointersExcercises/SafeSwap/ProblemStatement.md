# Safe Swap Using Pointers (Go Pointer Exercise)

## 📌 Problem Overview

This exercise is designed to practice **basic pointer usage in Go**, specifically:

- Passing variables by reference
- Safely dereferencing pointers
- Mutating values in place
- Preventing runtime panics using `nil` checks

You will implement a function that swaps two integers **only if it is safe to do so**.

---

## 🎯 Task Requirements

Write a function that swaps the values of two integers **in place**.

### Rules:
- The function must accept **pointers to integers**
- If **either pointer is `nil`**, the function must:
  - Do nothing
  - Not panic
- The function should not return anything

---

## 🔧 Function Signature

```go
func SafeSwap(a *int, b *int)
