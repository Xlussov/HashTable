# Hash Table in Go

This repository demonstrates a **generic Hash Table** implementation in **Go** with support for **any key/value types** and **open addressing with linear probing**.

---

## 📘 Implemented Features

* **Put** — insert or update a key-value pair
* **Get** — retrieve a value by key
* **Remove** — delete a key-value pair
* **Print** — display all key-value pairs
* **Size** — get the current number of elements
* **Cap** — get the current capacity of the table
* **Automatic Rehashing** — doubles capacity when load factor > 0.75

---

## 🧠 How It Works

Each operation works on generic types `K` and `V`. Keys must be **comparable** (`==` supported). Values can be any type. The table uses **open addressing** with **double hashing** to resolve collisions.

Example:

```go
ht := hashtable.New
ht.Put("Alice", 25)
ht.Put("Bob", 30)
value, ok := ht.Get("Alice")
fmt.Println(value, ok) // 25 true
```

Remove a key:

```go
ht.Remove("Bob")
```

Check size and capacity:

```go
fmt.Println(ht.Size()) // 1
fmt.Println(ht.Cap())  // 8
```

---

## 🚀 Run Example

```bash
go run ./cmd/main.go
```

Example of printing the hash table:

```go
ht.Print()
```

Output:

```
Index 0: Key = Alice, Value = 25
Index 1: Key = , Value = 0
Index 2: Key = , Value = 0
...
```

---

## 🧪 Run Tests

```bash
go test ./...
```

Tests cover:

* Inserting new key-value pairs
* Updating existing keys
* Retrieving values
* Removing keys
* Automatic rehashing when load factor exceeds threshold

---

## 🧩 Examples

### 1. Strings to Integers

```go
ht := hashtable.New
ht.Put("Alice", 25)
ht.Put("Bob", 30)
fmt.Println(ht.Get("Bob")) // 30 true
```

### 2. Integers to Strings

```go
ht := hashtable.New
ht.Put(1, "One")
ht.Put(2, "Two")
fmt.Println(ht.Get(1)) // One true
```

### 3. Custom Struct as Value

```go
type Person struct {
    Name string
    Age  int
}

ht := hashtable.New
ht.Put("Alice", Person{"Alice", 25})
fmt.Println(ht.Get("Alice")) // {Alice 25} true
```

---

© 2025 Danil Xlussov — minimalistic generic Hash Table library with open addressing and automatic rehashing support.
