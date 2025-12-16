package repl

import (
	"bufio"
	"fmt"
	"io"
	"mylang/evaluator"
	"mylang/lexer"
	"mylang/parser"
)

const PROMPT = ">> "

const MYLANG_LOGO = `
 /$$      /$$ /$$     /$$ /$$        /$$$$$$  /$$   /$$  /$$$$$$ 
| $$$    /$$$|  $$   /$$/| $$       /$$__  $$| $$$ | $$ /$$__  $$
| $$$$  /$$$$ \  $$ /$$/ | $$      | $$  \ $$| $$$$| $$| $$  \__/
| $$ $$/$$ $$  \  $$$$/  | $$      | $$$$$$$$| $$ $$ $$| $$ /$$$$
| $$  $$$| $$   \  $$/   | $$      | $$__  $$| $$  $$$$| $$|_  $$
| $$\  $ | $$    | $$    | $$      | $$  | $$| $$\  $$$| $$  \ $$
| $$ \/  | $$    | $$    | $$$$$$$$| $$  | $$| $$ \  $$|  $$$$$$/
|__/     |__/    |__/    |________/|__/  |__/|__/  \__/ \______/ 
                                                                                                                                                                 
`

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "*********************************************\n")
	io.WriteString(out, "Woops! We ran into an error here!\n")
	io.WriteString(out, " Parser errors:\n")
	for _, message := range errors {
		io.WriteString(out, "\t"+message+"\n")
	}
}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}
