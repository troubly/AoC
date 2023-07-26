# Input assumptions
- Each crate is identified by a single char -> this implies each crate representation (be it empty or not) takes 3 chars space in the line.  
- Each crate is separated by 1 space.  
- No physically impossible input, e.g.: 
```
[A] [D]
    [C]
[B]
 1   2
```
- Stack ids are ordered, begin at '1', and don't leave any gaps. So we wouldn't want something like:
```
[A] [B] [C]
 1   3   2
```
or
```
[A] [B] [C]
 2   3   4
```
or
```
[A] [B] [C]
 1   3   4
```