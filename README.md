# High Order Function

High-order functions are a fundamental concept in functional programming. These functions accept other functions as arguments, enabling the automation of common data transformation operations. In everyday coding, we often find ourselves performing tasks like filtering, finding specific elements, reducing lists, sorting data, and more. High-order functions provide a powerful way to abstract and reuse these transformation operations.

This repository serves as a demonstration of high-order function usage, showcasing the capabilities of the Go generic library. With the Go generic library, you can effortlessly apply high-order functions to various input arrays, eliminating the need to manually write these transformation operations repeatedly. Whether you're working with different data structures or use cases, this repository simplifies and enhances your development workflow. Explore the power of high-order functions and streamline your code with ease.

## Filter

```go
func main() {
    numbers := []int{1, 2, 3, 4, 5}
    getNumberGtTwo := func(n number) bool {
        return n > 2
    }
    numbersGtTwo := Filter(numbers, getNumberGreaterThanTwo)
}
```

## Find

```go
func main() {
    names := []string{"john", "doe", "robert", "john dorry"}
    findNameJohn := func(s string) bool {
        return s == "john"
    }
    johnName := Find(names, findNameJohn)
}
```

## Mapper

```go
func main() {
    names := []string{"john", "doe", "robert", "john dorry"}
    mapNameToUser := func(index int, name string) User {
        return User{
            Id:   index,
            Name: name,
        }
    }
    user := Map(names, mapNameToUser)
}
```

## Reduce

```go
func main() {
    numbers := []int{1, 2, 3, 4, 5}
    sumNumbers := func(total, number int) int {
        return total + number
    }
    total := Reduce(numbers, sumNumbers)
}
```

## Sorts

```go
func main() {
    users := []User{
        {
            Name: "John",
            Age:  20,
        },
        {
            Name: "Doe",
            Age:  19,
        },
        {
            Name: "Rocky",
            Age:  18,
        },
    }
    sortAgeAsc := func(before, after User) bool {
        return before.Age < after.Age
    }
    userSortedByAge := Reduce(users, sortAgeAsc)
}
```
