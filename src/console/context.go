package console

import (
	"errors"
	"reflect"
)

// Context contains interfaces for interaction with blockchain.
type Context struct {
	Blockchain    Blockchainer
	Block         Blocker
	newBlockchain reflect.Value
	newBlock      reflect.Value
	quitChannel   chan int
}

func callF(f reflect.Value, v ...interface{}) (res []reflect.Value, err error) {
	in := make([]reflect.Value, 0)
	for i := 0; i < len(v); i++ {
		in = append(in, reflect.ValueOf(v[i]))
	}

	ch := make(chan int)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				v, _ := e.(string)
				err = errors.New(v)
			}
			ch <- 1
		}()
		res = f.Call(in)
	}()
	<-ch

	return
}

// Quit exits console program.
func (c *Context) Quit() {
	c.quitChannel <- 1
}

// NewBlockchain create a blockchain invoking a registered blockchain
// function.
func (c *Context) NewBlockchain(v ...interface{}) interface{} {
	out, err := callF(c.newBlockchain, v...)

	if err != nil || len(out) <= 0 {
		return nil
	}

	return out[0].Interface()
}

// NewBlock create a block invoking a registered block function.
func (c *Context) NewBlock(v ...interface{}) interface{} {
	out, err := callF(c.newBlock, v...)

	if err != nil || len(out) <= 0 {
		return nil
	}

	return out[0].Interface()
}
