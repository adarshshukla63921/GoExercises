# BankAccount Methods with Pointer Receivers (Go Struct Exercise)

## 📌 Problem Overview

This exercise focuses on **using methods with pointer receivers in Go** to safely mutate struct data.

You are given a `BankAccount` struct.  
Your task is to implement **business logic as methods** on the struct instead of standalone functions.

This problem builds directly on:
- pointers
- struct mutation
- method receivers
- basic API design

---

## 🧱 Struct Definition

```go
type BankAccount struct {
	Owner   string
	Balance int
}
