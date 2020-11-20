//Blockchain

//a blockchain use essentially a public database that is distributed across multiple different peers
//what makes this revolutionary is that it doesn't rely on trust

//up until blockchain technology was invented, there were a lot of decent distributed solutions for technology
//however, they kind of hinged on the fact that every node would be trustworthy
//that is to say that every piece of data coming from each of your nodes has to have the correct data inside of it

//with blockchain, however, one, or let's say 49% of your nodes could be producing incorrect data, the database would be able to fix itself
//since the database doesn't need to rely on trusting the nodes, you can do a lot of really cool things with your blockchain
//for instance, create a cryptocurreny or create some smart contracts
//as the name blockchain implies, a blockchain is composed of multiple different blocks
//each block contains the data that we want to pass around inside of our database as well as a hash which is associated with the block itself

package main

import (
	"os"

	"github.com/AmirYunus/go_blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
