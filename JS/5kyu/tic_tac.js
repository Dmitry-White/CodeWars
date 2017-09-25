/**
 * Created on Mon Sep 25 23:26:13 2017
 * @author: Dmitry White
 */

 /**
  * TODO: If we were to set up a Tic-Tac-Toe game,
  * we would want to know whether the board's current state is solved,
  * wouldn't we? Our goal is to create a function that will check that for us!
  * Assume that the board comes in the form of a 3x3 array,
  * where the value is 0 if a spot is empty,
  * 1 if it is an X, or 2 if it is an O.
  * We want our function to return -1 if the board is not solved yet,
  * 1 if X won, 2 if O won, or 0 if it's a cat's game (i.e. a draw).
  * You may assume that the board passed in is valid in the context
  * of a game of Tic-Tac-Toe.
  */

function isSolved(board) {
        var result = -1;
        var leftDiag = board[0][0] == board[1][1]
                        && board[1][1] == board[2][2];
        var rightDiag = board[0][2] == board[1][1]
                        && board[1][1] == board[2][0];
        if (leftDiag) {
          return board[0][0] === 0 ? -1 : board[0][0];
        }
        else if (rightDiag) {
          return board[0][2] === 0 ? -1 : board[0][2];
        }

        for (var i = 0; i < 3; i++) {
            var hori = board[i][0] == board[i][1]
                        && board[i][0] == board[i][2]
                        && board[i][1]!== 0;
            var vert = board[0][i] == board[1][i]
                        && board[1][i] == board[2][i]
                        && board[1][i]!== 0;
            if (hori) {
              result = board[i][1] === 0 ? -1 : board[i][1];
            }
            else if (vert) {
              result = board[1][i] === 0 ? -1 : board[1][i];
            }
        }
        for (var i = 0; i < 3; i++) {
          if (board[i].includes(0)) return result;
        }
        return 0;
}
