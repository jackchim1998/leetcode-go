package solutionSet

func tictactoe(moves [][]int) string {
	row := [3]int{}
	col := [3]int{}
	diagonal := 0
	reverseDiagonal := 0
	player := 1

	for _, move := range moves {
		row[move[0]] += player
		col[move[1]] += player
		if move[0] == move[1] {
			diagonal += player
		}
		if move[0]+move[1] == 2 {
			reverseDiagonal += player
		}
		if row[move[0]] == 3 || row[move[0]] == -3 ||
			col[move[1]] == 3 || col[move[1]] == -3 ||
			diagonal == 3 || diagonal == -3 ||
			reverseDiagonal == 3 || reverseDiagonal == -3 {
			if player == 1 {
				return "A"
			} else {
				return "B"
			}
		}
		player *= -1
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
