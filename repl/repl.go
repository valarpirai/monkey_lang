package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey_lang/lexer"
	"monkey_lang/token"

	"github.com/fatih/color"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	defer color.Unset()
	scanner := bufio.NewScanner(in)
	for {
		color.Set(color.FgYellow)
		fmt.Printf(PROMPT)
		color.Set(color.FgWhite)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
