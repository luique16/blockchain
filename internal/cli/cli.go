package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/luique16/blockchain/internal/blockchain"
)

func Run(bc *blockchain.Blockchain) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		fmt.Print("Escolha uma opcao: ")

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			cmdSend(scanner, bc)
		case "2":
			cmdViewChain(bc)
		case "3":
			cmdViewBlock(scanner, bc)
		case "4":
			cmdValidate(bc)
		case "5":
			cmdInfo(bc)
		case "0":
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opcao invalida!")
		}

		if choice != "0" {
			fmt.Println()
			fmt.Print("Pressione Enter para continuar...")
			scanner.Scan()
		}
	}
}

func printMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("===================================")
	fmt.Println("       BLOCKCHAIN CLI v1.0")
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println("  1. Mandar mensagem")
	fmt.Println("  2. Ver blockchain")
	fmt.Println("  3. Ver detalhes de um bloco")
	fmt.Println("  4. Validar blockchain")
	fmt.Println("  5. Informacoes")
	fmt.Println("  0. Sair")
	fmt.Println()
	fmt.Println("-----------------------------------")
}

func cmdSend(scanner *bufio.Scanner, bc *blockchain.Blockchain) {
	fmt.Print("Usuario: ")
	if !scanner.Scan() {
		return
	}
	user := strings.TrimSpace(scanner.Text())

	fmt.Print("Mensagem: ")
	if !scanner.Scan() {
		return
	}
	message := strings.TrimSpace(scanner.Text())

	if user == "" || message == "" {
		fmt.Println("Erro: usuario e mensagem sao obrigatorios!")
		return
	}

	payload := blockchain.Payload{
		User:         user,
		Message:      message,
		Timestamp:    time.Now().String(),
		PreviousHash: bc.LastBlockHash(),
	}

	block := bc.CreateBlock(payload)
	fmt.Print("\nMinerando bloco... ")

	bc.MineBlock(block)

	last := bc.Chain[len(bc.Chain)-1]
	fmt.Println("OK!")
	fmt.Printf("Bloco #%d minerado com sucesso!\n", last.Header.Index)
	fmt.Printf("Hash: %x\n", last.Header.Hash)
	fmt.Printf("Nonce: %d\n", last.Header.Nonce)
}

func cmdViewChain(bc *blockchain.Blockchain) {
	if len(bc.Chain) == 0 {
		fmt.Println("Blockchain vazia.")
		return
	}

	fmt.Println()
	fmt.Println("  BLOCO | NONCE        | USUARIO          | MENSAGEM")
	fmt.Println("  ------+--------------+------------------+----------------------------------------")

	for i, b := range bc.Chain {
		user := b.Payload.User
		if user == "" {
			user = "(genesis)"
		}
		msg := b.Payload.Message
		if msg == "" {
			msg = "(genesis)"
		}
		if len(msg) > 36 {
			msg = msg[:33] + "..."
		}
		if len(user) > 16 {
			user = user[:13] + "..."
		}

		fmt.Printf("  %5d | %12d | %-16s | %s\n", i, b.Header.Nonce, user, msg)
		fmt.Printf("        hash: %x\n", b.Header.Hash)
		fmt.Println("  ------+--------------+------------------+----------------------------------------")
	}

	fmt.Printf("\nTotal de blocos: %d\n", len(bc.Chain))
}

func cmdViewBlock(scanner *bufio.Scanner, bc *blockchain.Blockchain) {
	fmt.Print("Indice do bloco: ")
	if !scanner.Scan() {
		return
	}

	idxStr := strings.TrimSpace(scanner.Text())
	idx, err := strconv.Atoi(idxStr)
	if err != nil || idx < 0 || idx >= len(bc.Chain) {
		fmt.Printf("Erro: indice invalido (0-%d)\n", len(bc.Chain)-1)
		return
	}

	b := bc.Chain[idx]

	fmt.Println()
	fmt.Println("  ========================================")
	fmt.Println("  BLOCO #", idx)
	fmt.Println("  ========================================")
	fmt.Printf("  Indice:      %d\n", b.Header.Index)
	fmt.Printf("  Nonce:       %d\n", b.Header.Nonce)
	fmt.Printf("  Hash:        %x\n", b.Header.Hash)
	fmt.Println("  ----------------------------------------")
	fmt.Printf("  Usuario:     %s\n", b.Payload.User)
	fmt.Printf("  Mensagem:    %s\n", b.Payload.Message)
	fmt.Printf("  Timestamp:   %s\n", b.Payload.Timestamp)
	fmt.Printf("  Previous:    %x\n", b.Payload.PreviousHash)
	fmt.Println("  ========================================")
}

func cmdValidate(bc *blockchain.Blockchain) {
	fmt.Println()

	if len(bc.Chain) <= 1 {
		fmt.Println("Blockchain valida (apenas bloco genesis).")
		return
	}

	valid := true
	for i := 1; i < len(bc.Chain); i++ {
		prev := bc.Chain[i-1]
		curr := bc.Chain[i]

		hashOk := curr.Header.Hash == curr.Payload.Hash(curr.Header.Nonce)
		linkOk := curr.Payload.PreviousHash == prev.Header.Hash

		if hashOk && linkOk {
			fmt.Printf("  Bloco #%d: VALIDO\n", i)
		} else {
			fmt.Printf("  Bloco #%d: INVALIDO (hash=%v, link=%v)\n", i, hashOk, linkOk)
			valid = false
		}
	}

	if valid {
		fmt.Println("\n  Blockchain VALIDA!")
	} else {
		fmt.Println("\n  Blockchain INVALIDA!")
	}
}

func cmdInfo(bc *blockchain.Blockchain) {
	fmt.Println()
	fmt.Println("  ========================================")
	fmt.Println("  INFORMACOES DA BLOCKCHAIN")
	fmt.Println("  ========================================")
	fmt.Printf("  Dificuldade:  %d\n", bc.Difficulty)
	fmt.Printf("  Proof char:   '%c'\n", bc.ProofChar)
	fmt.Printf("  Total blocos: %d\n", len(bc.Chain))
	fmt.Printf("  Genesis:      %x\n", bc.Chain[0].Header.Hash)
	fmt.Printf("  Ultimo bloco: %x\n", bc.Chain[len(bc.Chain)-1].Header.Hash)
	fmt.Println("  ========================================")
}
