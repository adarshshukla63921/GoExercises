# Day 2 — Task Manager (Struct + Pointer Exercise)

## 📌 Problem Overview

This exercise is designed to strengthen your understanding of **structs, pointers, and safe in-place mutation in Go**.

You will build a small in-memory **Task Manager** where tasks can be updated, escalated, and marked as completed using **pointer-based functions**.

The focus is on:
- Correct pointer handling
- Nil safety
- Preserving business invariants
- Writing clean, reusable logic functions

---

## 🧱 Struct Definition

```go
type Task struct {
	ID       int
	Name     string
	Priority int  // Range: 1 (low) to 5 (high)
	Done     bool
}
