![](https://github.com/saleem/peg_game/workflows/peg_game/badge.svg)

# Peg game 

The peg game is a conceptually simple yet tactically intriguiing single-player board game. It was popularized by a restaurant chain. The board is in the shape of an equilateral triangle with 15 holes. At the start of the game, 14 holes are occuped by individual pegs, and one hole is empty. The locaton of this empty hole is arbitrary chosen by the game player, although the specific position greatly affects the strategy the player _should_ take to win the game

On each move, the player removes one peg from the board. The peg to be removed must:
1. Have an adjacent hole on one side, and
2. Have an adjacent peg on the linearly opposite side.

E.g. `.  a  b` ==> `b  .  .` (`a` and `b` are letters representing two pegs, however, there's nothing special about either a or b). If you've played checkers, each move is identical to capturing an opponents piece by "jumping" over it.

The game is over when either:
1. There is exactly one peg left on the board. The player **wins**; or
2. There are multiple pegs on the board; however, there are no possible legal moves because the pegs are separated by multiple holes. The player **loses**.


## Starting positions

### Vertex (3 symmetric variations)
```
    .
   x x
  x x x
 x x x x
x x x x x
```

### Side-center (3 symmetric variations)
```
    x
   x x
  . x x
 x x x x
x x x x x
```

### Inner-triangle (3 symmetric variations)
```
    x
   x x
  x x x
 x . x x
x x x x x

```

### Vertex-adjacent (6 symmetric variations)
```
    x
   . x
  x x x
 x x x x
x x x x x
```

# How to run

Clone the repo. You can run unit tests by running `make go-test` in your shell.

# References

http://peggamesolutions.com/
