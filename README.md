# ChessProblem

This repository holds a solution to a given chess challenge program.

The relevant board logic is in an internal package separated from the main
application to enforce some boundary between them, demonstrating that the
chess application can modestly stand alone.

There are only two pieces, a White Bishop at c3 and a Black Rook at h1.

In the problem, a coin flip determines whether the Rook moves up or right.  Dice determine how far it moves in the chosen direction.

## Assumptions

*    This board wraps around at the edges for **both** pieces, though the problem only refers to the Rook's wrapping behaviour.  I assume the Bishop can attack the Rook through an edge.
*    The directions only refer to the Rook wrapping off the right edge and top edge of the board because it only moves right and up.  I assume that the bottom and left edges wrap for computing Bishop moves. 
*    If the Rook happens to hit the Bishop, it immediately wins.
*    If the Rook would land on itself by wrapping around, this is fine.
*    If the Rook moves **through** the Bishop this is also fine, as it could have wrapped back the other direction.  The destination square is a destination, not a path.

## Known issues

*    Only Rook and Bishop movements are implemented
*    No turn order enforcement is performed.  Pieces can make any legal move.