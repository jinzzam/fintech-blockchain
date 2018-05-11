package execute

import (
	"blockchain/console/command"
	"blockchain/core"
	"blockchain/log"
	"bytes"
	"fmt"
	"strconv"
)

func TransactionCommands() {
	_ = command.AddCommand("", command.Command{
		Name:        "transaction",
		ShortName:   "t",
		Description: "manage transaction of candidate block",
		Commands: []command.Command{
			command.Command{
				Name:        "new",
				Description: "create a transaction of candidate block",
				Commands:    make([]command.Command, 0),
				Flags:       nil,
				Run:         NewTransactionInCandidateBlock,
			},
			command.Command{
				Name:        "list",
				ShortName:   "ls",
				Description: "show transactions list of candidate block",
				Commands:    make([]command.Command, 0),
				Flags:       nil,
				Run:         ShowCandidateBlockTransactionsList,
			},
		},
		Flags: nil,
		Run:   nil,
	})
}

func NewTransactionInCandidateBlock(bcidx uint64, amount uint64, from []byte, to []byte) error {
	log.Info("Create New Transaction in the Candidate block")
	log.Info(perforatedLine)
	bcidxs := strconv.FormatUint(bcidx, 10)
	amounts := strconv.FormatUint(amount, 10)

	tl, fl := len(to), len(from)
	if len(to) > 20 {
		tl = 20
	}
	if len(from) > 20 {
		fl = 20
	}
	var toa, froma core.Address

	copy(toa[:], to[:tl])
	copy(froma[:], from[:fl])

	log.Debug("Blockchain index : " + bcidxs)

	bc, err := getBlockchain(bcidx)
	if err != nil {
		return err
	}

	t := core.NewTransaction(amount, froma, toa)
	_ = bc.CandidateBlock.AddTransaction(t)
	log.Info("Amount : " + amounts + "\tFrom : " + froma.ToString() + "\tTo : " + toa.ToString())
	log.Debug("Create completed")
	log.Info(perforatedLine)

	return nil
}

func ShowCandidateBlockTransactionsList(bcidx uint64) error {
	log.Info("Show Candidate Block Transactions list")
	log.Info(perforatedLine)

	bcidxs := strconv.FormatUint(bcidx, 10)
	bc, err := getBlockchain(bcidx)

	if err != nil {
		return err
	}

	b := bc.CandidateBlock

	log.Info(transactionsString(b, "Blockchain index : "+bcidxs+"'s Candidate Block"))
	log.Info(perforatedLine)

	return nil
}

func transactionsString(b *core.Block, name string) string {
	res := name + "\nNum\tAmount\tFrom\t\t\t\t\t\tTo\n"
	buffer := bytes.NewBuffer([]byte{})

	for idx, t := range b.Body.Transactions {
		fmt.Fprintf(buffer, strconv.Itoa(idx+1))
		fmt.Fprintf(buffer, "\t%v", t.Data.Amount)
		fmt.Fprintf(buffer, "\t%v", t.From.ToString())
		fmt.Fprintf(buffer, "\t%v\n", t.Data.To.ToString())
	}

	res = res + buffer.String()
	return res
}
