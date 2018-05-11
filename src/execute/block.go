package execute

import (
	"blockchain/console/command"
	"blockchain/core"
	"blockchain/log"
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

func BlockCommands() {
	_ = command.AddCommand("", command.Command{
		Name:        "block",
		ShortName:   "b",
		Description: "manage blocks",
		Commands: []command.Command{
			command.Command{
				Name:          "new",
				Description:   "create a candidate block",
				Commands:      make([]command.Command, 0),
				Flags:         nil,
				DefaultParams: []interface{}{uint64(1)},
				Run:           NewCandidateBlock,
			},
			command.Command{
				Name:          "attach",
				Description:   "Attach candidate block to blockchain",
				Commands:      make([]command.Command, 0),
				Flags:         nil,
				DefaultParams: []interface{}{uint64(1)},
				Run:           AttachCandidateBlockToBlockchain,
			},
			command.Command{
				Name:        "list",
				ShortName:   "ls",
				Description: "show list of blocks",
				Commands:    make([]command.Command, 0),
				Flags:       nil,
				Run:         ShowBlocksList,
			},
			command.Command{
				Name:          "info",
				Description:   "show information of block",
				Commands:      make([]command.Command, 0),
				Flags:         nil,
				DefaultParams: []interface{}{uint64(1), uint64(0)},
				Run:           ShowBlockInformation,
			},
		},
		Flags: nil,
		Run:   nil,
	})
}

func ShowBlocksList(bcidx uint64) error {
	log.Debug("Show Blocks List")
	log.Info(perforatedLine)

	bc, err := getBlockchain(bcidx)

	if err != nil {
		return err
	}

	log.Info("Blockchain index : " + strconv.FormatUint(bcidx, 10))

	for idx, b := range bc.Blocks {
		i := strconv.Itoa(idx)
		log.Info(blockStringInfo(&b, "Block Index : "+i))
		log.Info(perforatedLine)
	}

	return nil
}

func ShowBlockInformation(bcidx uint64, bidx uint64) error {
	log.Debug("Show Block Information")
	log.Info(perforatedLine)

	bcidxs := strconv.FormatUint(bcidx, 10)
	bidxs := strconv.FormatUint(bidx, 10)
	bc, err := getBlockchain(bcidx)

	if err != nil {
		return err
	}

	if bidx > bc.BlockchainHeight-1 {
		return errors.New("Incorrect block index")
	}

	b := &bc.Blocks[bidx]

	if bidx == 0 {
		bidxs = "Candidate"
		b = bc.CandidateBlock
	}

	log.Info(blockStringInfo(b, "Blockchain index : "+bcidxs+"\tBlock Index : "+bidxs))
	log.Info(perforatedLine)

	return nil
}

func NewCandidateBlock(bcidx uint64) error {
	log.Debug("Create New Candidate Block")
	log.Info(perforatedLine)

	bcidxs := strconv.FormatUint(bcidx, 10)
	bc, err := getBlockchain(bcidx)

	if err != nil {
		return err
	}

	bc.CandidateBlock = core.NewBlock(&bc.Blocks[bc.BlockchainHeight-1])

	log.Info(blockStringInfo(bc.CandidateBlock, "Blockchain index : "+bcidxs+"'s Candidate Block"))
	log.Info(perforatedLine)
	log.Debug("Create completed")
	return nil
}

func AttachCandidateBlockToBlockchain(bcidx uint64) error {
	log.Debug("Attach Candidate Block to Blockchain")
	log.Info(perforatedLine)

	bcidxs := strconv.FormatUint(bcidx, 10)
	log.Debug("Blockchain index : " + bcidxs)
	bc, err := getBlockchain(bcidx)

	if err != nil {
		return err
	}
	//pow.Mining(bc.CandidateBlock)

	_ = bc.AddBlock()
	log.Info(perforatedLine)
	log.Debug("Attach completed")
	return nil
}

func blockStringInfo(b *core.Block, title string) string {
	buffer := bytes.NewBuffer([]byte{})
	if b != nil {
		var ph, mh string
		for _, v := range []struct {
			str  *string
			hash [32]byte
		}{
			{&ph, b.Header.PreviousHash},
			{&mh, b.Header.MerkleRootHash},
			//{&dh, b.Header.Difficulty},
		} {
			for _, h := range v.hash {
				*v.str += fmt.Sprintf("%02x", h)
			}
		}

		fmt.Fprintf(buffer, "PreviousHash     %v\n", ph)
		fmt.Fprintf(buffer, "MerkleRootHash   %v\n", mh)
		//fmt.Fprintf(buffer, "Difficulty       %v\n", dh)
		fmt.Fprintf(buffer, "Timestamp        %v\n", b.Header.Timestamp)
		fmt.Fprintf(buffer, "Index            %v\n", b.Header.Index)
		fmt.Fprintf(buffer, "Transactions     %v\n", len(b.Body.Transactions))
		fmt.Fprintf(buffer, "%v", transactionsString(b, ""))
	}

	res := title + "\n" + buffer.String()
	return res
}
