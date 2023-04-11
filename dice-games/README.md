# DICE GAMES
Create a dice game script that accepts input N number of players and M number dice, with the following rules:
1. At the beginning of the game, each player gets M units of dice.
2. All players will roll their respective dice simultaneously
3. Each player will check the results of their dice roll and evaluate as follows:
    a. The number 6 dice will be removed from the game and added as a point for that player.
    b. Dice number 1 will be given to the player sitting next to it.
    For example, the first player will give the dice number 1 to the second player.
    c. Dice numbers 2,3,4 and 5 will still be played by players.
4. After evaluation, players who still have dice will repeat the 2nd step until only 1 player remains.
    a. Players who don't have any more dice are considered to have finished playing.
5. The player who has the most points wins.

### RUNNING THE APPS
To run the code, you can use
```
    go run src/main.go
```

### HOW TO USE
1. Input the number of player first
2. Then input the number of dice
3. The project will show the result, the winner of the dice game who has the most points and the loser of the dice game who has remaining dice

### BUILT WITH

[Go](https://go.dev/)- An open source programming language that makes it simple to build secure, scalable systems.