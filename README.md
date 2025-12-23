# MyLang

> A tiny, dynamically typed language with a friendly REPL. Built in Go.

```
 /$$      /$$ /$$     /$$ /$$        /$$$$$$  /$$   /$$  /$$$$$$ 
| $$$    /$$$|  $$   /$$/| $$       /$$__  $$| $$$ | $$ /$$__  $$
| $$$$  /$$$$ \  $$ /$$/ | $$      | $$  \ $$| $$$$| $$| $$  \__/
| $$ $$/$$ $$  \  $$$$/  | $$      | $$$$$$$$| $$ $$ $$| $$ /$$$$
| $$  $$$| $$   \  $$/   | $$      | $$__  $$| $$  $$$$| $$|_  $$
| $$\  $ | $$    | $$    | $$      | $$  | $$| $$\  $$$| $$  \ $$
| $$ \/  | $$    | $$    | $$$$$$$$| $$  | $$| $$ \  $$|  $$$$$$/
|__/     |__/    |__/    |________/|__/  |__/|__/  \__/ \______/ 
```

## Overview

MyLang is a Monkey-inspired interpreter with a clean lexer, Pratt parser, and tree-walking evaluator. It supports integers, booleans, strings, arrays, hashes, functions with closures, and a handful of builtins. Extra operators like `%`, `<=`, and `>=` are included out of the box.

## Quick start

```bash
go run ./...
```

You’ll see the banner and land in the REPL prompt `>> `. Type `exit` or `Ctrl+C` to quit.

## Run tests

```bash
go test ./...
```

## Language cheatsheet

**Values**: integers, booleans, strings, arrays, hashes, functions, `null`  
**Vars**: `let answer = 42;`  
**Arithmetic**: `+ - * / %`  
**Comparisons**: `== != < > <= >=`  
**Truthiness**: `false`, `null`, and `0` are falsey; everything else is truthy  
**Conditionals**: `if (x < 10) { ... } else { ... }`  
**Functions/closures**: `let add = function(a, b) { a + b; }; add(1, 2);`  
**Arrays**: `[1, 2, 3][0]  // => 1`  
**Hashes**: `{"name": "Ada"}["name"]  // => Ada`

### Builtins

| Function        | Purpose                               |
| --------------- | ------------------------------------- |
| `len(x)`        | Length of string or array             |
| `first(arr)`    | First element or `null`               |
| `last(arr)`     | Last element or `null`                |
| `rest(arr)`     | Copy without the first element        |
| `push(arr, v)`  | New array with `v` appended           |
| `puts(x, ...)`  | Print inspected values, return `null` |

## Sample REPL session

```
>> let twice = function(f, x) { f(f(x)); };
>> let inc = function(n) { n + 1; };
>> twice(inc, 5);
7
>> let nums = [1, 2, 3];
>> len(nums);
3
>> {"greeting": "hi"}["greeting"];
hi
```

## Project layout

- `main.go` — CLI entrypoint; prints the banner and starts the REPL.
- `repl/` — prompt loop connecting lexer, parser, evaluator, and environment.
- `lexer/`, `token/` — lexical analysis and token definitions.
- `parser/`, `ast/` — Pratt parser and AST node types.
- `evaluator/` — interpreter, builtins, truthiness, indexing, hash/array ops.
- `object/` — runtime objects (values, functions, environments, hash keys).

## Implementation notes

- Functions are first-class and close over the environment where they are defined.
- Hash keys must be hashable types (integers, booleans, strings); non-hashable keys raise an error.
- Out-of-bounds array access returns `null` to keep execution safe in the REPL.
