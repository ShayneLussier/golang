# Go study track

Each lesson is a folder with one `_test.go` file. Notes live in `Example_*` functions;
the `// Output:` comment is checked, so a wrong note fails the test.

```
go test ./...                                         # check every lesson
go test ./01_types                                    # check one lesson
go test -v -run Example_constants ./01_types          # run one example
```

## Track

1. `01_types` - basic types, zero values, operators, conversion, declaration, constants
2. `02_composite_types` - arrays, slices, maps, structs
