package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/luique16/blockchain/internal/blockchain"
	"github.com/luique16/blockchain/internal/cli"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Dificuldade: ")
	if !scanner.Scan() {
		return
	}
	difficulty, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || difficulty < 1 {
		fmt.Println("Dificuldade invalida. Usando 4.")
		difficulty = 4
	}

	fmt.Print("Caractere do inicio do hash: ")
	if !scanner.Scan() {
		return
	}
	proofStr := strings.TrimSpace(scanner.Text())
	if len(proofStr) != 1 || !strings.ContainsAny(proofStr, "0123456789abcdefABCDEF") {
		fmt.Println("Caractere invalido. Use 0-9 ou a-f. Usando '0'.")
		proofStr = "0"
	}
	proofChar := []rune(proofStr)[0]

	bc := blockchain.NewBlockchain(difficulty, proofChar)
	fmt.Printf("Blockchain criada! Dificuldade=%d, ProofChar='%c'\n", difficulty, proofChar)
	fmt.Println()

	cli.Run(bc)
}
