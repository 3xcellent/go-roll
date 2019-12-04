# go-roll
My attempt at a cli dnd dice roller

## Install
```
go get github.com/3xcellent/go-roll
```

## Syntax
`go-roll d20` - roll a 20-sided die

`go-roll d20h` - roll twice and keep the highest (i.e. w/ advantage) 

`go-roll d20l` - roll twice and keep the lowest (i.e. w/ disadvantage)

`go-roll d20+2` - add modifier to roll

`go-roll 4d8` - roll 4 8-sided die 

syntax can be combined, but doesn't always make sense `go-roll 100d100h+99`

## Verbose Output
Sometimes you like to see what the system rolled for you:
```go
go-roll -v 2d20+9
```

will output something similar to:
```                                                                                                                                                                                                                                                                       ✔  713  08:37:11
Rolls: 18, 19
Modifier: +9
Total: 46 (min/max 11/49)
```

## Known Issues
* `go-roll` is only parsing modifiers up to two characters currently.  If there's a good reason to add modifiers for numbers higher than 99, please create an issue.