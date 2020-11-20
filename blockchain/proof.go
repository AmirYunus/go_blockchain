//define as package blockchain which is the name of the folder
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

//our difficulty will be static for this implementation (consider creating an algorithm that will make the difficulty flexible)
const Difficulty = 12

//create a proof of work struct
type ProofOfWork struct {
	Block  *Block	//pointer (*) to a block
	Target *big.Int	//target is a number that represents that nonce requirement
			//derived from our difficulty
}

//create a new proof of work by taking a pointer (*) to a block and produce a pointer (*) to a proof of work
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)	//create target by casting the number 1 as a new big int
	target.Lsh(target, uint(256-Difficulty))	//take 256 (i.e. number of bytes in each hash) and subtract the difficulty from it
							//take the target to shift the number of bytes over by this number
							//Lsh function is short hand for left shift

	pow := &ProofOfWork{b, target}	//take the target which have been left shifted and put it into an instance of a proof of work and add the block to it

	return pow	//return proof of work
}

//replace the original DeriveHash function with this method
//the method is a pointer (*) to a proof of work struct and will take in the nonce which is an int and output a slice of bytes
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,		//get the block's previous hash
			pow.Block.HashTransactions(),	//get the block's data
			ToHex(int64(nonce)),		//cast our nounce into an int64 and then call the ToHex function
			ToHex(int64(Difficulty)),	//do the same for our difficulty constant
		},
		[]byte{},	//combine it with another byte struct using the byte join function to create a cohesive set of bytes
	)

	return data	//return the joined byte struct as the data
}

//to check if the hash meets a set of requirements
//this main computational function will be called Run
//have it run on the pointer (*) of the proof of work
//the function itself will output an int and a slice of bytes inside of a tuple
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0	//define nonce as 0

	for nonce < math.MaxInt64 {			//virtually an infinite loop
		data := pow.InitData(nonce)		//assign data by calling proof of work InitData function with our nonce
		hash = sha256.Sum256(data)		//take the data and hash it into a SHA-256 format using Sum256

		fmt.Printf("\r%x", hash)		//see the process happening
							//hash is changing inline because of the \r

		intHash.SetBytes(hash[:])		//convert the hash into big int by calling the SetBytes function and passing our hash

		if intHash.Cmp(pow.Target) == -1 {	//compare our proof of work target with this bew big int version of our hash
							//if result is -1 then we want to break out of loop
							//-1 means our hash is less than our target value
							//this means we have already signed the block
			break
		} else {				//otherwise increase our nonce by 1 and repeat our process again
			nonce++
		}

	}
	fmt.Println()	//create a new line for new hashes

	return nonce, hash[:]	//return the nonce and hash from this function
}

//create validate method which goes off the proof of work pointer (*) and returns a boolean
//the idea is that after we have run our proof of work Run function, we will have the nonce which will allow us to derive the hash which met the target we wanted
//run that cycle one more time to show that this hash is valid
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int	//set up big int of our hash

	data := pow.InitData(pow.Block.Nonce)	//call InitData with our proof of work's nonce
	hash := sha256.Sum256(data)		//take the data and convert it to a hash

	intHash.SetBytes(hash[:])	//convert that hash into a big int and put it into our intHash variable

	return intHash.Cmp(pow.Target) == -1	//compare if the result is -1 and return true if the block is valid
}

//create a utility function which will take in an int64 and output a slice of bytes
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)	//this will create a new bytes buffer
	err := binary.Write(buff, binary.BigEndian, num)	//use a binary write to take our number and decode it into bytes
								//this binary big endian signififes how we want our bytes to be organised

	Handle(err)

	return buff.Bytes()	//return the bytes portion of our buffer to hex function
}
