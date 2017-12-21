## AI Design
Need a way to "traverse a board"
* Can visit all board states for a given depth (without accounting for new tile placement)
* Returns an action for the best move, determined by heuristics
* Each node contains a board state, a "score", and a pointers to the parent node, and up/down/left/right nodes.



Score from board state:
    8 => score of 12 = 8 + 4
    8 and 4 => score of 16 = (8 + 4) + (4)

    4, 4, 64, 16 => score of 356 = ? 

    to get to a 64
    32 and 32 (once)
    16 and 16 (twice)
    8 and 8 (thrice)
    4 and 4 (fourice)

