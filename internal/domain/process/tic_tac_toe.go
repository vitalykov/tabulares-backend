package process

import (
	"board-games/internal/domain/model"
	"cmp"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"slices"
	"strconv"
)

const (
	defaultDepth = 10
)

type tT = model.TicTacToeType

type TicTacToeProcessor struct{}

func (p TicTacToeProcessor) GenerateMove(game model.Game[tT]) model.Move[tT] {
	var pos pos
	if game.NumPlayers == 2 {
		pos = minimaxGetPos(game)
	} else {
		pos = randomGetPos(game)
	}
	return model.Move[tT]{
		Player: game.Turn,
		Actions: []model.Action[tT]{
			{
				Type: model.ActionAdd,
				Figure: model.Figure[tT]{
					Player: game.Turn,
					Type:   model.TicTacToeMark,
				},
				Args: []int{pos.row, pos.col},
			},
		},
	}
}

type pos struct {
	row, col int
}

func getPositions(board model.Board[tT], numMoves int) []pos {
	positions := make([]pos, 0, board.Height()*board.Width()-numMoves)
	for i := range board.Height() {
		for j := range board.Width() {
			if board.Cells[i][j].Type == model.TicTacToeNoFigure {
				positions = append(positions, pos{i, j})
			}
		}
	}
	return positions
}

func randomGetPos(game model.Game[tT]) pos {
	rows := game.Board.Height()
	cols := game.Board.Width()
	var i, j int
	const maxCount = 30
	count := 0

	for count < maxCount {
		i = rand.Intn(rows)
		j = rand.Intn(cols)
		fig := game.Board.Cells[i][j]
		if fig.Type == model.TicTacToeNoFigure {
			break
		}
		count++
	}
	return pos{i, j}
}

// TODO: remove it
var totalCount int

type eval struct {
	move  pos
	score int
}

func printEvalsReport(player int, numMoves int, evals []eval) {
	s1, s2 := "Wins", "Loses"
	if player == 0 {
		s1, s2 = s2, s1
	}
	for _, e := range evals {
		moveStr := fmt.Sprintf("Move: %d %d.", e.move.row, e.move.col)
		if e.score > 0 {
			log.Printf("%s %s in %d moves.", moveStr, s1, (evalInf-e.score-numMoves)/2)
		} else if e.score < 0 {
			log.Printf("%s %s in %d moves.", moveStr, s2, (e.score+evalInf-numMoves)/2)
		} else {
			log.Printf("%s Maybe draw.", moveStr)
		}
	}
}

// Calculate appropriate depth for minimax algorithm based on the estimation:
//
//	n = p ^ d
//
// where:
//
// n - approximate total number of calculated moves,
//
// p - number of possible moves in current position,
//
// d - depth of minimax algorithm
func calculateDepth(numMoves int) int {
	const maxMoveCount = 5_000_000
	return int(math.Floor(math.Log(maxMoveCount) / math.Log(float64(max(numMoves, 2)))))
}

// TODO: make valid n-player minimax algorithm
func minimaxGetPos(game model.Game[tT]) pos {
	toWin, _ := strconv.Atoi(game.AdditionalInfo)
	positions := getPositions(game.Board, len(game.Moves))
	if len(positions) == 0 {
		return pos{-1, -1}
	}
	evals := make([]eval, len(positions))
	depth := calculateDepth(len(positions))
	player := game.Turn
	for i := range positions {
		pos := positions[i]
		evals[i].move = pos
		game.Board.Cells[pos.row][pos.col] = model.Figure[tT]{Player: player, Type: model.TicTacToeMark}
		evals[i].score = minimax(game.Board, len(game.Moves)+1, depth, pos, player, toWin, -evalInf, evalInf)
		game.Board.Cells[pos.row][pos.col] = model.Figure[tT]{}
	}
	slices.SortFunc(evals, func(a, b eval) int {
		return cmp.Compare(a.score, b.score)
	})

	// TODO: remove logging number of calculated moves
	log.Printf("tic-tac-toe minimax: Depth: %d. Moves calculated: %d\n", depth, totalCount)
	// printEvalsReport(player, len(game.Moves), evals)
	totalCount = 0

	if player == maxPlayer {
		return chooseMaxMove(evals)
	} else {
		return chooseMinMove(evals)
	}
}

// evals must be sorted in ascending order by the score
func chooseMaxMove(evals []eval) pos {
	l := len(evals)
	if l == 0 {
		return pos{}
	}
	firstWin, _ := slices.BinarySearchFunc(evals, 1, func(e eval, t int) int {
		return cmp.Compare(e.score, t)
	})
	firstDraw, _ := slices.BinarySearchFunc(evals, 0, func(e eval, t int) int {
		return cmp.Compare(e.score, t)
	})
	firstMaxLose := firstDraw - 1
	if firstMaxLose >= 0 {
		firstMaxLose, _ = slices.BinarySearchFunc(evals, evals[firstMaxLose].score, func(e eval, t int) int {
			return cmp.Compare(e.score, t)
		})
	}
	// log.Printf("l: %d. first win: %d. first draw: %d. first max lose: %d", l, firstWin, firstDraw, firstMaxLose)
	var i int
	if firstWin != l {
		i = firstWin + rand.Intn(l-firstWin)
	} else if firstDraw != l {
		i = firstDraw + rand.Intn(l-firstDraw)
	} else {
		i = firstMaxLose + rand.Intn(l-firstMaxLose)
	}
	return evals[i].move
}

// evals must be sorted in ascending order by the score
func chooseMinMove(evals []eval) pos {
	l := len(evals)
	if l == 0 {
		return pos{}
	}
	lastWin, _ := slices.BinarySearchFunc(evals, 0, func(e eval, t int) int {
		return cmp.Compare(e.score, t)
	})
	lastWin--
	lastMinLose, _ := slices.BinarySearchFunc(evals, 1, func(e eval, t int) int {
		return cmp.Compare(e.score, t)
	})
	lastDraw := lastMinLose - 1
	if lastMinLose < l {
		lastMinLose, _ = slices.BinarySearchFunc(evals, evals[lastMinLose].score+1, func(e eval, t int) int {
			return cmp.Compare(e.score, t)
		})
		lastMinLose--
	}
	// log.Printf("l: %d. last win: %d. last draw: %d. last min lose: %d", l, lastWin, lastDraw, lastMinLose)
	var i int
	if lastWin >= 0 {
		i = rand.Intn(lastWin + 1)
	} else if lastDraw >= 0 {
		i = rand.Intn(lastDraw + 1)
	} else {
		i = rand.Intn(lastMinLose + 1)
	}
	return evals[i].move
}

const (
	minPlayer = iota
	maxPlayer
)

const (
	evalInf = 1_000_000
)

// returns -1 if player 0 wins and +1 if player 1 wins. Otherwise returns 0
func evaluate(board model.Board[tT], numMoves int, move pos, player int, toWin int) int {
	if playerWin(board, move.row, move.col, player, toWin) {
		return (player*2 - 1) * (evalInf - numMoves)
	}
	return 0
}

func minimax(board model.Board[tT], numMoves int, depth int, move pos, player int, toWin int, alpha, beta int) int {
	// TODO: remove count
	totalCount++
	eval := evaluate(board, numMoves, move, player, toWin)
	if depth == 0 || eval != 0 {
		return eval
	}
	positions := getPositions(board, numMoves)
	if len(positions) == 0 {
		return 0
	}
	if player == minPlayer {
		maxEval := -evalInf
		for i := range positions {
			pos := positions[i]
			board.Cells[pos.row][pos.col] = model.Figure[tT]{Player: maxPlayer, Type: model.TicTacToeMark}
			eval = minimax(board, numMoves+1, depth-1, pos, maxPlayer, toWin, alpha, beta)
			board.Cells[pos.row][pos.col] = model.Figure[tT]{}
			maxEval = max(eval, maxEval)
			alpha = max(eval, alpha)
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		minEval := evalInf
		for i := range positions {
			pos := positions[i]
			board.Cells[pos.row][pos.col] = model.Figure[tT]{Player: minPlayer, Type: model.TicTacToeMark}
			eval = minimax(board, numMoves+1, depth-1, pos, minPlayer, toWin, alpha, beta)
			board.Cells[pos.row][pos.col] = model.Figure[tT]{}
			minEval = min(eval, minEval)
			beta = min(eval, beta)
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}

func (p TicTacToeProcessor) NextTurn(game model.Game[tT]) int {
	return (game.Turn + 1) % game.NumPlayers
}

func (p TicTacToeProcessor) FirstTurn(game model.Game[tT]) int {
	return 0
}

func (p TicTacToeProcessor) ValidateMove(game model.Game[tT], move model.Move[tT]) error {
	row := move.Actions[0].Args[0]
	col := move.Actions[0].Args[1]
	if row >= game.Board.Height() || col >= game.Board.Width() || row < 0 || col < 0 {
		return errors.New("Move outside the board")
	}
	figure := game.Board.Cells[row][col]
	if move.Player != game.Turn || figure.Type != model.TicTacToeNoFigure {
		return errors.New("Not the player turn")
	}
	return nil
}

func (p TicTacToeProcessor) GetWinner(game model.Game[tT]) int {
	toWin, _ := strconv.Atoi(game.AdditionalInfo)
	lastMove := game.Moves[len(game.Moves)-1]
	row := lastMove.Actions[0].Args[0]
	col := lastMove.Actions[0].Args[1]
	lastPlayer := lastMove.Player
	if playerWin(game.Board, row, col, lastPlayer, toWin) {
		return lastPlayer
	}

	if len(game.Moves) == game.Board.Height()*game.Board.Width() {
		return model.Draw
	}
	return model.NoWinner
}

// returns true if the move with coordinates (row, col) gives a win to player
func playerWin(board model.Board[tT], row, col int, player int, toWin int) bool {
	i1 := max(0, row-(toWin-1))
	j1 := max(0, col-(toWin-1))
	i2 := min(board.Height()-1, row+(toWin-1))
	j2 := min(board.Width()-1, col+(toWin-1))
	return winHorizontal(board, row, j1, j2, player, toWin) ||
		winVertical(board, col, i1, i2, player, toWin) ||
		winDiagonal(board, row, col, player, toWin)
}

func winHorizontal(board model.Board[tT], i, j1, j2 int, player int, toWin int) bool {
	count := 0
	for j := j1; j <= j2; j++ {
		fig := board.Cells[i][j]
		if fig.Type == model.TicTacToeMark && fig.Player == player {
			count++
			if count == toWin {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

func winVertical(board model.Board[tT], j, i1, i2 int, player int, toWin int) bool {
	count := 0
	for i := i1; i <= i2; i++ {
		fig := board.Cells[i][j]
		if fig.Type == model.TicTacToeMark && fig.Player == player {
			count++
			if count == toWin {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

func winDiagonal(board model.Board[tT], row, col int, player int, toWin int) bool {
	delta := min(row, col, toWin)
	i1, j1 := row-delta, col-delta
	delta = min(board.Height()-1-row, board.Width()-1-col, toWin)
	i2, j2 := row+delta, col+delta
	count := 0
	for i, j := i1, j1; i <= i2 && j <= j2; i, j = i+1, j+1 {
		fig := board.Cells[i][j]
		if fig.Type == model.TicTacToeMark && fig.Player == player {
			count++
			if count == toWin {
				return true
			}
		} else {
			count = 0
		}
	}
	count = 0
	delta = min(board.Height()-1-row, col, toWin)
	i1, j1 = row+delta, col-delta
	delta = min(row, board.Width()-1-col, toWin)
	i2, j2 = row-delta, col+delta
	for i, j := i1, j1; i >= i2 && j <= j2; i, j = i-1, j+1 {
		fig := board.Cells[i][j]
		if fig.Type == model.TicTacToeMark && fig.Player == player {
			count++
			if count == toWin {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}
