package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"board-games/internal/domain/model"
	"board-games/internal/domain/process"
	"board-games/internal/repository/mock"
	uModel "board-games/internal/usecases/model"
	uService "board-games/internal/usecases/service"
)

var pics = map[int]string{
	0: "🔵",
	1: "🔴",
	2: "🎓",
	3: "🎻",
}

func PrintTicTacToeBoard(game *model.Game[model.TicTacToeType]) {
	for i, row := range game.Board.Cells {
		for _, fig := range row {
			pic := "⬜"
			if fig.Type == model.TicTacToeMark {
				pic = pics[fig.Player]
			}
			fmt.Print(pic)
		}
		fmt.Printf(" %d\n", i)
	}
	for j := range game.Board.Width() {
		fmt.Printf(" %d", j)
	}
	fmt.Println()
}

func ReadMove(sc *bufio.Scanner, playerID uModel.PlayerID) (uModel.MoveInfo, bool) {
	sc.Scan()
	move := sc.Text()
	if move == "u" || move == "undo" {
		return uModel.MoveInfo{}, true
	}
	return uModel.MoveInfo{
		PlayerID: playerID,
		MoveRepr: move,
	}, false
}

func main() {
	proc := process.TicTacToeProcessor{}
	repo := mock.MockRepository{}
	master := uService.NewGameMaster(repo)
	interactor := uService.NewDefaultGameInteractor(repo, proc)
	var width, height int
	fmt.Println("Enter height and width of the board:")
	fmt.Scan(&height, &width)
	var minToWin string
	fmt.Println("Enter min marks in a row to win:")
	fmt.Scan(&minToWin)
	scanner := bufio.NewScanner(os.Stdin)
	var players []uModel.PlayerID
OUTER:
	for len(players) == 0 {
		fmt.Println("Enter player ids in one line (space separated):")
		scanner.Scan()
		idsStr := strings.Split(scanner.Text(), " ")
		ids := make([]int, 0, len(idsStr))
		for _, idStr := range idsStr {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Bad player id format:", err)
				continue OUTER
			}
			ids = append(ids, id)
		}
		for _, id := range ids {
			players = append(players, uModel.PlayerID(id))
		}
	}
	newGameInfo := uModel.NewGameInfo{
		Type:           uModel.TicTacToeType,
		BoardWidth:     height,
		BoardHeight:    width,
		Players:        players,
		AdditionalInfo: minToWin,
	}
	gameInfo, _ := master.CreateGame(newGameInfo)
	interactor.StartGame(gameInfo)
	fmt.Println("Game started!")
	game := gameInfo.Game.(*model.Game[model.TicTacToeType])
	PrintTicTacToeBoard(game)
	for gameInfo.Status != uModel.Finished {
		game = gameInfo.Game.(*model.Game[model.TicTacToeType])
		undo := false
		var move uModel.MoveInfo
		if gameInfo.Turn < 0 {
			move, _ = interactor.GetHint(gameInfo)
			time.Sleep(time.Second)
		} else {
			fmt.Println("Enter move: PlayerID: ", gameInfo.Players[game.Turn])
			move, undo = ReadMove(scanner, gameInfo.Players[game.Turn])
		}
		if undo {
			fmt.Println("Undo")
			interactor.UndoMove(gameInfo)
			undo = false
		} else {
			fmt.Println("Move")
			err := interactor.MakeMove(gameInfo, move)
			if err != nil {
				continue
			}
		}
		PrintTicTacToeBoard(game)
	}

	if gameInfo.Winner == uModel.NoWinner {
		fmt.Println("Draw")
	} else if gameInfo.Winner > 0 {
		fmt.Println("You win. Congraturalions!")
	} else {
		fmt.Println("You are loser")
	}
	fmt.Println("total moves: ", len(gameInfo.Moves))
}
