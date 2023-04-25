# qogi - Quality of Golang Improvements

Small things that makes your development in Golang a little bit more comfy.

## Some Examples

### Pointer Stuff

To Pointers:

```go
stringPointer := qogi.ToPointer("abcd123456")

uintPointer := qogi.ToPointer(unit(123456))

...
```

From Pointers:

```go
uintPointerValue := qogi.UintOrZero(uintPointer)
uintZeroValue := qogi.UintOrZero(nil)

stringPointerValue := StringOrEmpty(stringPointer)
emptyString := StringOrEmpty(nil)
```

### String to Numbers

```go
unsignedInteger, err := AtoUi("7")

float, err := AtoF("12345.54321")
```

### Numbers to String

```go
uintAsString := UiToA(123)
```

## Feedback

Feel free to suggests, contribute, request, improve, criticize and so on :)

Have fun,
 someone.earth
