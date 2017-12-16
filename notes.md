## AI Design
Need a way to "traverse a board"
* Can visit all board states for a given depth (without accounting for new tile placement)
* Returns an action for the best move, determined by heuristics
* Each node contains a board state, a "score", and a pointers to the parent node, and up/down/left/right nodes.

