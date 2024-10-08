# gothreadsafe: Thread-Safe Data Structures in Go

**gothreadsafe** is a Go package designed to provide thread-safe versions of Go's built-in data structures such as slices, maps, and other types. This package ensures that operations on these data structures are safe against race conditions, making it ideal for concurrent applications.

## Features

- **Thread-Safe Slices:** Perform append, remove, and many other operations on slices without worrying about race conditions.
- **Thread-Safe Maps:** Use maps with concurrent read/write operations safely.
- **Thread-Safe Sets:** Perform various operations with sets.
- **Additional Thread-Safe Types:** Includes thread-safe implementations of other commonly used types.
- **Easy to Use:** Designed to be a drop-in replacement for non-thread-safe versions with minimal changes to your code.

## Installation

To install SafeMap, use the following command:

```sh
go get -u github.com/sebastiankristof/gothreadsafe
```

## Usage

Import the package in your Go code:

```go
import "github.com/sebastiankristof/gothreadsafe"
```

Use the thread-safe data structures in your code:

### Maps:

```go
// Create a new thread-safe map
m := gothreadsafe.NewSafeMap()

// Set a key-value pair
m.Set("key", "value")

// Get the value for a key
value, ok := m.Get("key")
if ok {
    fmt.Println(value)
}

// Delete a key-value pair
m.Delete("key")
```

### Slices:

```go
// Create a new thread-safe slice
s := gothreadsafe.NewSafeSlice()

// Append an element to the slice
s.Append("element")

// Get the element at an index
element, ok := s.Get(0)
if ok {
    fmt.Println(element)
}

// Remove an element from the slice
s.Remove(0)
```

### Sets:

```go
// Create a new thread-safe set
s1 := gothreadsafe.NewSet()
s2 := gothreadsafe.NewSet()

// Append an element to the set
s1.Add("element1")
s2.Add("element2")

// Check if set contains element 1
isInSet := s1.Contains("element1")
fmt.Printf("Set contains the element: %t", isInSet)

// Perform a union operation
setUnion := s1.Union(s2)

// Check that union set contains element 2
isInSet := setUnion.Contains("element2")
fmt.Printf("Set contains the element: %t", isInSet)

// Clear the base sets
s1.Clear()
s2.Clear()
```

### Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any bugs, features, or improvements.

### License
This project is licensed under the MIT License - see the LICENSE file for details.
