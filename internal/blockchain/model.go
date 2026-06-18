package blockchain

type Header struct {
	Index 	     int
	Nonce 	     int
	Hash 	     string
}

type Payload struct {
	User         string
	Message	     string
	Timestamp    string

	PreviousHash string
}

type Block struct {
	Header       Header
	Payload  	 Payload
}

type Blockchain struct {
	Chain        []Block
}
