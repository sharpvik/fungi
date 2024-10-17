# Fungi

[![Go Reference](https://pkg.go.dev/badge/pkg.go.dev/github.com/sharpvik/fungi.svg)](https://pkg.go.dev/github.com/sharpvik/fungi@v1.0.0)

**Fungi** provides a great suite of functional stream processing primitives that
can be used for a wide range of purposes. Use this library to describe your
intent declaratively and produce elegant code that is easy to read and refactor.

## Import

```bash
go get github.com/sharpvik/fungi@latest
```

## Spread The Spores

Very soon `fungi` will start popping up all over your codebase! And that is a
good thing. Here are some things it can help you with:

1. [`Filter`][filter] out irrelevant items using a custom validation function.
2. Apply custom transformations to stream items
   ([`Map`][map], [`TryMap`][try_map]).
3. Select a [`Range`][range] of items you're interested in.
4. [`Sort`][sort] items with a generic comparator.
5. [`Page`][page] items efficiently based on page number and size.
6. [`Loop`][loop] through every item (see also [`ForEach`][for_each]).
7. Collect items into a Go builtin `slice` or `map`
   ([`CollectSlice`][slice], [`CollectMap`][hmap]).
8. [`Find`][find] an item that fits a description.

[filter]: ./filter.go
[map]: ./map.go
[try_map]: ./try_map.go
[range]: ./range.go
[sort]: ./sort.go
[page]: ./page.go
[loop]: ./loop.go
[for_each]: ./for_each.go
[slice]: ./slice.go
[hmap]: ./hmap.go
[find]: ./find.go

## Don't Sweat

Fungi is very well-tested with a consistent test coverage of **over 95%**.
See for yourself:

```bash
git clone git@github.com:sharpvik/fungi.git
cd fungi
go test -cover
```

```text
[6th of December 2022 checked out at v1.1.0]
PASS
coverage: 98.1% of statements
ok      github.com/sharpvik/fungi       0.193s
```

Moreover, our tests can and _should_ be used as examples: they are written with
clarity and readability in mind.

> Test files have the `_test.go` suffix. Browse through, don't be shy!

## Make It Stream

Written with generics, `fungi` gives you the flexibility to apply it to any
iterator that implements the very simple [`fungi.Stream`](stream.go) interface.

Suppose you already have multiple iterable `type`s that fetch elements using a
method called `Recv`. Here's how you can write a converter function to make them
all comply with the [`fungi.Stream`](stream.go) interface:

```go
// Every one of your iterable receivers follows this generic interface.
type Receiver[T any] interface {
    Recv() (T, error)
}

// receiverStream implements fungi.Stream interface.
type receiverStream[T any] struct {
    Receiver[T]
}

// Next wraps Recv method of the origincal Receiver.
func (rs receiverStream[T]) Next() (T, error) {
    return rs.Recv()
}

// ReceiverStream converts any Receiver into a fungi.Stream.
func ReceiverStream[T any](r Receiver[T]) fungi.Stream[T] {
    return receiverStream[T]{r}
}
```

## Declare Elegance

Here's how your code is going to look soon:

```go
func GetNamesOfMyDrinkingBuddies() ([]string, error) {
    users := ReceiverStream[*User](GetMyFriends())
    over18 := fungi.FilterMap(func(u *User) (name string, isLegal bool) {
        return u.Name, u.Age >= 18
    })
    sortAlphabetically := fungi.Sort(func(a, b string) bool { return a < b })
    return fungi.CollectSlice(sortAlphabetically(over18(users)))
}
```
