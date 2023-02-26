package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <qscore> <input.fastq> <output.fastq>\n", os.Args[0])
		os.Exit(1)
	}

	qscore, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid Q score: %s\n", os.Args[1])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	out, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %s\n", err)
		os.Exit(1)
	}
	defer out.Close()

	scanner := bufio.NewScanner(f)
	writer := bufio.NewWriter(out)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) == 4 {
			q := lines[3]
			if qBelowThreshold(q, qscore) {
				sequence := []rune(lines[1])
				quality := []rune(lines[3])
				for i, c := range quality {
					if int(c)-33 <= qscore || c == '{' || c == 'Z' {
						sequence[i] = 'N'
					}
				}
				lines[1] = string(sequence)
			}
			for _, line := range lines {
				writer.WriteString(line + "\n")
			}
			lines = lines[:0]
		}
	}

	if len(lines) != 0 {
		fmt.Fprintf(os.Stderr, "Error: incomplete record at the end of the file\n")
		os.Exit(1)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %s\n", err)
		os.Exit(1)
	}

	if err := writer.Flush(); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Done!")
}

func qBelowThreshold(q string, threshold int) bool {
	for _, c := range q {
		if c == '{' || int(c)-33 <= threshold {
			return true
		}
	}
	return false
}
