# Go Blackboard

Go Blackboard is a simple key/value store inspired by go's flag package.

## Installation

```
go get -u github.com/SirMetathyst/go-blackboard
```

## Supported Types

 * ```string```
 * ```[]string```
 * ```bool```
 * ```int```

## Example

```go
func main() {
    bb := blackboard.NewBlackboard()

    s := "my_string_value"
    bb.SetStringP("my_string", &s)
    bb.SetString("another_string", "value")

    fmt.Println(*bb.StringP("my_string"))
    fmt.Println(*bb.StringP("another_string"))
}
```

```go
func main() {
    s := "my_string_value"
    blackboard.SetStringP("my_string", &s)
    blackboard.SetString("another_string", "value")

    fmt.Println(*blackboard.StringP("my_string"))
    fmt.Println(*blackboard.StringP("another_string"))
}
```

### []String

```go
SetStringSliceP(key string, value *[]string)
SetStringSlice(key string, value []string)
StringSliceP(key string) *[]string
AllStringSlice() []KSS
```

### String

```go
SetStringP(key string, value *string)
SetString(key string, value string)
StringP(key string) *string
AllString() []KS
```

### Bool

```go
SetBoolP(key string, value *bool)
SetBool(key string, value bool)
BoolP(key string) *bool
AllBool() []KB
```

### Int

```go
SetIntP(key string, value *int)
SetInt(key string, value int)
IntP(key string) *int
AllInt() []KI
```