# parser

## Usage

```bash
N/A
```

## Example

```go

input := "id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}"
parser := parser.NewParser(input)

placeholders := parser.ParsePlaceholders()

parser.ReplaceWithTokens(input, placeholders, map[string]string{
    "ID": "1",
    "NAME": "John",
    "AGE": "25",
    "SALARY": "50000",
})

result := parser.ReplaceWithTokens(input, placeholders, func(key string) string {
    switch key {
    case "ID":
        return "1"
    case "NAME":
        return "John"
    case "AGE":
        return "25"
    case "SALARY":
        return "50000"
    }
})

fmt.Println(result)

// output:
// id:1/name:John/age:25/salary:50000

```