# Templa3

Templa3 is a CLI tool that applies a json to a go template (text/template). It implements all sprig functions and some made in-house.

## Custom functions

### isdef 

```go
{{isdef .dict "key1" "key2" "key3" ...}}
```
Returns `true` if the given keys chain exists in the given map (dict).

#### Example
Template:
```go.tmpl
{{/* template */ -}}
{{ isdef . "a" "b" "c" }}
```
JSON file

```jsonc
// vars
{
    "a": {
        "b": {
            "c": [1, 2, 3]
        }
    }
}
```

Output:
```
true
```

## Build source

```bash
go build
```

## Install from source

```bash
go install github.com/jmwielandt/templa3@latest
```

No precompiled binaries are currently provided.

## Use
```
Usage of templa3.exe:
  -template string
        Path to the template file
  -vars string
        Path to the vars json file
  -verbose
        Enables stdout prints after and during template execution
```

## How to run it?

```bash
templa3 -template examples/template.go.tmpl -vars examples/vars.json
```