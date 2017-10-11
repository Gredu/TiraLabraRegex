package main

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

		matchLine(inputs, machine)

	} else {
		// TODO
		fmt.Println("hi pipe!")
	}

}

func matchLine(path string, machine State) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if match(scanner.Text(), machine) {
			fmt.Println(scanner.Text())
		}
	}
}
