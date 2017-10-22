package trex

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func benchmarkMatchLine(path string, regexp string, b *testing.B) {
	verbose := false
	machine := generateMachine(parseRegexp(regexp))

	inFile, _ := os.Open(path)
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	b.ResetTimer()
	for scanner.Scan() {
		if match(scanner.Text(), machine) && verbose {
			fmt.Println(scanner.Text())
		}
	}
	b.StopTimer()

	err := inFile.Close()
	if err != nil {
		fmt.Println(err)
	}

}

func BenchmarkMatchLine63M(b *testing.B) {
	benchmarkMatchLine("data/dna/homoSapiens20-63M.fa", "ATATATAGATTAGATGAATTGTAGTCAAACTAGATCTAATTTAGTTTTTAGTATGTTTTA", b)
}

func BenchmarkMatchLineDots63M(b *testing.B) {
	benchmarkMatchLine("data/dna/homoSapiens20-63M.fa", "ATAT.TAGATT.GATGAATTGTAGTCA...TAGAT..AATT.AGTTTTT.GTATGTTTTA", b)
}

func BenchmarkMatchLine81M(b *testing.B) {
	benchmarkMatchLine("data/dna/homoSapiens17-81M.fa", "AGCACAATTTGGAAGAAAACATTTCCATCTGTTAATAAAGAGCAACGGCCTCTGGTCATA", b)
}

func BenchmarkMatchLineDots81M(b *testing.B) {
	benchmarkMatchLine("data/dna/homoSapiens17-81M.fa", "AGCAC.A.TTGGAAGAAAACATTTCCATCT..TAATAAA.AGCAACGGCCTCTGG...TA", b)
}

func BenchmarkMatchLine104M(b *testing.B) {
	benchmarkMatchLine("data/dna/homoSapiens14-104M.fa", "GGAGAGTCTCAGGAACTTCCCAGGTTGTCTCAGGTCCTCTGTGGGCTCTATCAGCTCCAC", b)
}

func BenchmarkMatchLineDots104M(b *testing.B) {
	benchmarkMatchLine("data/dna/homoSapiens14-104M.fa", "GGAGAGTCTCA..AA..TCCCAGGTTGTCT..GGTCCTCTGTGGG.TCT.T.AGCTCCAC", b)
}
