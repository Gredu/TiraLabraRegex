package trex

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		inputs := os.Args[1]
		regexp := os.Args[2]

		machine := generateMachine(parseRegexp(regexp))

		matchLine(inputs, machine, true)

	} else {
		// TODO
		fmt.Println("hi pipe!")
	}

}

// matchLine check each line of a file and checks if regular expression match on that line using match() function.
func matchLine(path string, machine State, verbose bool) {
	inFile, _ := os.Open(path)
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if match(scanner.Text(), machine) && verbose {
			fmt.Println(scanner.Text())
		}
	}

	err := inFile.Close()
	if err != nil {
		fmt.Println(err)
	}
}
