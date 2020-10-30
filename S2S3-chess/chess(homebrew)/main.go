package main
import (
	"fmt"
	"os"
	"strings"

	. "./Model"
)

func main() {
	cb := new(ChessBoard)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		fmt.Print("Instruction recieved.\n")
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		err = runCommand(*cb, cmdString)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(cb ChessBoard, commandStr string) (e error) {

	commandStr = strings.TrimSuffix(commandStr, "\n")
	args := strings.Fields(commandStr)
	switch args[0] {
	case "new":
		// TODO Create a new game on a classic 8x8 board.
		// TODO Display the board on console.

		cb.Init()
		k := new(Knight)
		cb[4][4] = k
		cb.Print()
		break
	case "move":
		// TODO Move a piece. (syntax: move <from> <to>)
		// TODO The command line should be in the form of move A2 A4.
		// TODO     => meaning move piece from position A2 to A4
		// TODO Display the board on console.
		start := new(ChessCoord)
		end := new(ChessCoord)
		start.ReadFrom(args[1])
		end.ReadFrom(args[2])
		if i, e := cb.MovePiece(*start, *end); i != 0 {
			fmt.Print(e)
		} else {
			cb.Print()
		}
		break
	default:
		e = fmt.Errorf("unknown command %s", args[0])