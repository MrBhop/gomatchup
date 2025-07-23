# gomatchup
gomatchup is a CLI Tool for generating matchups and teams. 
It implements REPL functionality, which allows the user to add participants to a list, from which team compositions and matchups can be generated.

Example:
```
add Player1
add Player2
add Player3
add Player4
add Player5
add constraint Player1 Not Player2
generate teams
> team 1: Player1, Player3
> team 2: Player2, Player4, Player5
```
