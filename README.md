# The Myers Diff algorithm
The Myers Diff algorithm is a linear space algorithm that calculates the shortest edit script between two sequences.
This Algo is used in the `diff` command in Unix systems and in the `git diff` command.

## How it works
The algorithm works by creating a graph of the shortest path between two sequences.
The graph is created by moving diagonally in the graph (k = x - y) and finding the shortest path to the end of the graph.
The algorithm then backtracks the graph to find the shortest path between the two sequences.

## Move types
1. → horizontal moves (x + 1, y) - delete 
2. ↓ vertical moves (x, y + 1)  - insert
3. ↘ diagonal moves (x + 1, y + 1) - keep

## Example
Let's say we have two sequences:
- Sequence 1: ABC
- Sequence 2: ADC

Diagonals:

    (0,0), (1,1) - match `A..` in `A` and `B`
    (2,2), (3,3) - match `..C` in `A` and `B`

The graph would look like this:
```
         A     B     C
      0     1     2     3
  0 (0,0) (1,0) (2,0) (3,0)
A        \ 
  1 (0,1) (1,1) (2,1) (3,1)
D 
  2 (0,2) (1,2) (2,2) (3,2)
C                    \ 
  3 (0,3) (1,3) (2,3) (3,3)
```

The shortest path would be:

    (0,0) -> (1,1) => diagonal move (keep) => keep A
    (1,1) -> (2,1) => horizontal move (delete) => delete B
    (2,1) -> (2,2) => vertical move (insert) => insert D
    (2,2) -> (3,3) => diagonal move (keep) => keep C

The edit script would be:
```
- Keep A:   ABC -> (1, 1) -> ABC
- Delete B: ABC -> (2, 1) -> AC
- Insert D: AC  -> (3, 2) -> ADC
- Keep C:   ADC -> (3, 3) -> ADC
```

## Examples
Use `-d` flag to debug the algorithm
```shell
go run ./cmd/.  examples/1a examples/1b 
```

```shell
go run ./cmd/.  examples/2a examples/2b 
```

```shell
go run ./cmd/.  examples/3a examples/3b 
```

```shell
go run ./cmd/.  examples/4a examples/4b 
```

```shell
go run ./cmd/.  examples/5a examples/5b 
```

```shell
go run ./cmd/.  examples/6a examples/6b 
```

```shell
go run ./cmd/.  examples/7a examples/7b 
```

```shell
go run ./cmd/.  examples/8a examples/8b 
```

```shell
go run ./cmd/.  examples/9a examples/9b 
```

```shell
go run ./cmd/. -l  examples/10a examples/10b 
```


## Helpful links
- [Myers Diff algorithm](https://blog.robertelder.org/diff-algorithm/)
- [James Coglan - Myers Diff algorithm](https://blog.jcoglan.com/2017/02/12/the-myers-diff-algorithm-part-1/)
