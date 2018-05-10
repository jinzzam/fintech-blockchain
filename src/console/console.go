package console

import (
	"blockchain/console/command"
	"blockchain/log"
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

var ctx Context

func start() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			go start()
		}
	}()

	for reader := bufio.NewReader(os.Stdin); ; {

		fmt.Printf("> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.Replace(cmd, "\n", "", 1)
		_ = command.StartCommand(cmd)
	}
}

// Start run a console deamon.
func Start() {
	ctx.quitChannel = make(chan int, 1)

	go start()

	<-ctx.quitChannel
}

// GetContext returns a global context for console program.
func GetContext() *Context {
	return &ctx
}

// RegisterBlockchain registers Blockchainer interface for console interaction.
// The first argument is a structure satisfying Blockchainer interface. And
// The second argument is a method for generating a new blockchain.
func RegisterBlockchain(bc Blockchainer, nbc interface{}) {
	ctx.Blockchain = bc
	ctx.newBlockchain = reflect.ValueOf(nbc)
}

// RegisterBlock registers Block interface for console interaction.
// The composition of parameters is similar to RegisterBlockchain.
func RegisterBlock(nb interface{}) {
	ctx.newBlock = reflect.ValueOf(nb)
}
