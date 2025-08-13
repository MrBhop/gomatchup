# gomatchup
gomatchup is a CLI Tool for generating team compositions based on constraints.<br>
It implements REPL functionality, which allows the user to add players to a list, from which team compositions can be generated.

Example:
```
gomatchup >> add-player Player1
Player 'player1' added.

gomatchup >> add-player Player2
Player 'player2' added.

gomatchup >> add-player Player3
Player 'player3' added.

gomatchup >> add-player Player4
Player 'player4' added.

gomatchup >> add-player Player5
Player 'player5' added.

gomatchup >> add-constraint Player1 Player2
Added exclusion constraint, 'player1' X 'player2'.

gomatchup >> generate-teams 2
Team 1 of 2:
        - 'player2'
        - 'player5'
        - 'player3'

Team 2 of 2:
        - 'player1'
        - 'player4'
```
