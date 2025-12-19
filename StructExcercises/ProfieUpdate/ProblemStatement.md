# In-Place User Profile Update (Go Struct Pointer Exercise)

## 📌 Problem Overview

This exercise is designed to practice **using pointers with structs in Go** and understanding how **in-place mutation** works.

You are given a `User` struct representing a user profile.  
Your task is to write a function that **updates the user profile directly**, without returning a new struct.

---

## 🧱 Struct Definition

```go
type User struct {
	Name  string
	Age   int
	Email string
}
