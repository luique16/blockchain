package blockchain

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func NewBlockchain(difficulty int, proofChar rune) *Blockchain {
	blockchain := Blockchain{
		Difficulty:   difficulty,
		ProofChar:    proofChar,
		Chain:        []Block{},
	}

	payload := Payload{
		User:         "",
		Message:      "",
		Timestamp:    time.Now().String(),
		PreviousHash: "",
	}

	block := blockchain.CreateBlock(payload)

	blockchain.MineBlock(block)

	return &blockchain
}


func (p Payload) String() string {
	bytes, _ := json.Marshal(p)

	return string(bytes)
}

func (b *Blockchain) CreateBlock(payload Payload) Block {
	block := Block{
		Header: Header{
			Index: len(b.Chain),
			Nonce: 0,
			Hash:  "",
		},
		Payload: payload,
	}

	return block
}

func (b *Blockchain) LastBlockHash() string {
	if len(b.Chain) == 0 {
		return ""
	}

	return b.Chain[len(b.Chain)-1].Header.Hash
}

func (b *Blockchain) MineBlock(block Block) {
	nonce := 0

	for {
		hash := block.Payload.Hash(nonce)

		if string(hash[:b.Difficulty]) == strings.Repeat(string(b.ProofChar), b.Difficulty) {
			block.Header.Nonce = nonce
			block.Header.Hash = string(hash[:])

			b.Chain = append(b.Chain, block)

			break
		}

		nonce++
	}

	b.Chain = append(b.Chain, block)
}

func (b *Blockchain) ValidateChain(log bool) bool {
	for i := 1; i < len(b.Chain); i++ {
		if b.Chain[i].Payload.PreviousHash != b.Chain[i-1].Header.Hash &&
		   b.Chain[i].Header.Hash == b.Chain[i].Payload.Hash(b.Chain[i].Header.Nonce) {
			if log {
				println("Block " + strconv.Itoa(i) + " invalid")
			}

			return false
		}

		if log {
			println("Block " + strconv.Itoa(i) + " valid")
		}
	}
	
	return true
}

func (b *Blockchain) PrintChain() {
	for i := 0; i < len(b.Chain); i++ {
		println("[" + b.Chain[i].Payload.Timestamp + "] " + b.Chain[i].Payload.User + ": " + b.Chain[i].Payload.Message)
	}
}
