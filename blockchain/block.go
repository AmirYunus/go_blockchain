package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//to represent a block, we can create a struct which has a hash, data field and a previous hash field
type Block struct {
	Hash		 []byte		//the hash field represents the hash of this block
					//we derive the hash inside our block from the data inside of the block and the previous hash that is being passed to the block
					//we will also add timestamp and other fields which will also go into the hash calculation

	Transactions     []*Transaction	//the data field represents the data inside of this block and this can be anything from ledgers to documents and images (update for transactions)

	PrevHash	 []byte		//the previous hash represents the last block's hash
					//having this previous hash allows us to link the blocks together sort of like a back linked list
					//(i.e. each block inside our blockchain references the last block that was created inside of the blockchain)

	Nonce		 int
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

//this function takes in a string of data (update for transaction) and then it takes in the previous hash from the last block and it outputs a pointer (*) to a block
func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}	//create a new block using a block constructor which will be a reference (&) to a block
							//for the hash field, we will just put in an empty slice of bytes
							//for the data field, we will take in the data string and convert it into a slice of bytes
							//we will take the previous hash from input
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block	//return block from the CreateBlock function
}

//create the first block in the blockchain
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})	//return a new CreateBlock call with data that we want
								//we get an array of pointers (*) to transaction and an empty previous hash which is a slice of bytes
}

func (b *Block) Serialise() []byte {
	var res bytes.Buffer

	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialise(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
