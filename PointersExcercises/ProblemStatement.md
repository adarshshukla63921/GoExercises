# In-Place Normalization (Go Pointer Exercise)

## 📌 Problem Overview

This exercise is designed to practice **basic pointer usage in Go**, specifically focusing on **in-place mutation of data**.

You are given a slice of integers.  
Your task is to **normalize the slice in place** such that:

- The **minimum value becomes 0**
- Every other element becomes `(value - min)`

⚠️ You must modify the original slice.  
⚠️ No new slice should be returned.

---

## 🧠 Learning Goals

- Understand how **pointers work in Go**
- Practice **dereferencing pointers**
- Mutate data **in place**
- Reason about when pointers are **necessary vs unnecessary**
- Avoid runtime panics and unintended copies

---

## 🧪 Example

### Input
```go
[]int{10, 5, 7, 20}
