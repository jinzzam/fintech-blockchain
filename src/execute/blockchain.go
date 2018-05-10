package execute

import (
	"blockchain/console/command"
	"blockchain/core"
	"blockchain/log"
	"bytes"
	"errors"
	"fmt"
)

const perforatedLine string = "-----------------------------------------------------"

func BlockchainCommands() {
	_ = command.AddCommand("", command.Command{
		Name:        "blockchain",
		ShortName:   "bc",
		Description: "manage blockchains",
		Commands: []command.Command{
			command.Command{
				Name:        "new",
				Description: "create a blockchain",
				Commands:    make([]command.Command, 0),
				Flags:       nil,
				Run:         NewBlockchain,
			},
			command.Command{
				Name:        "number",
				ShortName:   "num",
				Description: "show the number of blockchains",
				Commands:    make([]command.Command, 0),
				Flags:       nil,
				Run:         ShowNumberofBlockchains,
			},
			command.Command{
				Name:          "info",
				Description:   "show blockchain Information",
				Commands:      nil,
				Flags:         nil,
				DefaultParams: []interface{}{uint64(1)},
				Run:           ShowBlockchainInformation,
			},
		},
		Flags: nil,
		Run:   nil,
	})
}

func NewBlockchain() error {
	log.Debug("Create New Blockchain")
	log.Info(perforatedLine)

	bc := core.NewBlockchain()
	core.AppendBlockchain(bc)

	log.Debug("Create completed")

	return nil
}

func ShowNumberofBlockchains() error {
	log.Debug("Show Number of Blockchains")
	log.Info(perforatedLine)

	result := ""

	result += fmt.Sprintf("Discovered blockchains: %v ", len(core.GlobalBlockchains))

	log.Info(result)
	log.Info(perforatedLine)

	return nil
}

func ShowBlockchainInformation(bcidx uint64) error {
	log.Debug("Show Blockchain Information")
	log.Info(perforatedLine)
	bc, err := getBlockchain(bcidx)

	if err != nil {
		return err
	}

	log.Info(blockchainStringInfo(bc, ""))
	log.Info(perforatedLine)

	return nil
}

func getBlockchain(bcidx uint64) (*core.Blockchain, error) {
	if bcidx < 0 && bcidx > uint64(len(core.GlobalBlockchains)-1) {
		return nil, errors.New("Invalid Select Blockchain")
	}

	return core.GlobalBlockchains[bcidx], nil
}

func blockchainStringInfo(bc *core.Blockchain, title string) string {
	buffer := bytes.NewBuffer([]byte{})
	fmt.Fprintf(buffer, "Height %v\n\n", bc.BlockchainHeight)
	fmt.Fprintf(buffer, "%v\n", blockStringInfo(bc.GenesisBlock, "Genesis Block"))
	fmt.Fprintf(buffer, "%v", blockStringInfo(bc.CandidateBlock, "Candidate Block"))

	res := title + "\n" + buffer.String()
	return res
}
