package handlers

import (
	//"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/entryBlock"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/state"
)

var CheckStatus = func() *sevCmd {
	cmd := new(sevCmd)
	cmd.helpMsg = "serveridentity check 'ROOT_ID_CHAIN"
	cmd.description = "Trys to build the identity associated with the root id chain"
	cmd.execFunc = func(args []string) {
		if len(args) == 0 {
			fmt.Println("Not enough arguments")
			return
		}
		os.Args = args[1:]
		host := flag.String("host", "localhost:8088", "Change factomd location")
		details := flag.Bool("v", false, "Printout the full identity and keys")
		flag.Parse()

		factom.SetFactomdServer(*host)
		cid, err := primitives.HexToHash(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		err = checkStatus(cid, *details)
		if err != nil {
			if *details {
				fmt.Println(err.Error())
			}
			fmt.Println(false)
			return
		}
	}
	Help.Add("Trys to build the identity associated with the root id chain", cmd)
	return cmd
}()

type BetterIEBEntry struct {
	IEB    interfaces.IEBEntry
	Height uint32
}

func getAllBetterIEBsFromChainID(chainID string) []BetterIEBEntry {
	rEnts := make([]BetterIEBEntry, 0)

	head, err := factom.GetChainHead(chainID)
	if err != nil {
		return rEnts
	}

	if head == primitives.NewZeroHash().String() {
		return rEnts
	}

	eblock, err := factom.GetEBlock(head)
	if err != nil {
		return rEnts
	}

	for eblock != nil {
		ents := eblock.EntryList
		height := eblock.Header.DBHeight
		for _, e := range ents {
			entry, err := factom.GetEntry(e.EntryHash)
			if err != nil {
				continue
			}
			ieb := factomEntryToIEBEntry(entry)
			t := new(BetterIEBEntry)
			t.IEB = ieb
			t.Height = uint32(height)
			rEnts = append(rEnts, *t)
		}
		if eblock.Header.PrevKeyMR == primitives.NewZeroHash().String() {
			break
		}
		eblock, err = factom.GetEBlock(eblock.Header.PrevKeyMR)
		if err != nil {
			break
		}
	}

	return rEnts
}

func checkStatus(cid interfaces.IHash, verbose bool) error {
	st := new(state.State)
	st.LoadConfig("", "")
	st.InitMapDB()

	index := st.CreateBlankFactomIdentity(cid)

	managementChain, _ := primitives.HexToHash(state.MAIN_FACTOM_IDENTITY_LIST)

	fEnts, err := factom.GetAllChainEntries(managementChain.String())
	if err != nil {
		return err
	}

	// Main ID Chain
	ents := factomEntrysToIEBEntrys(fEnts)

	if len(ents) == 0 {
		return errors.New("Identity Error: No main Main Factom Identity Chain chain created")
	}

	// Root ID chain
	rEnts := getAllBetterIEBsFromChainID(cid.String())
	if len(rEnts) == 0 {
		return errors.New("No root identity chain found")
	}

	for _, e := range rEnts {
		state.LoadIdentityByEntry(e.IEB, st, e.Height, false)
	}

	mEnts := getAllBetterIEBsFromChainID(managementChain.String())

	for _, ent := range mEnts {
		if len(ent.IEB.ExternalIDs()) > 3 {
			// This is the Register Factom Identity Message
			if len(ent.IEB.ExternalIDs()[2]) == 32 {
				idChain := primitives.NewHash(ent.IEB.ExternalIDs()[2][:32])
				if string(ent.IEB.ExternalIDs()[1]) == "Register Factom Identity" && cid.IsSameAs(idChain) {
					state.RegisterFactomIdentity(ent.IEB, cid, ent.Height, st)
					break // Found the registration
				}
			}
		}
	}

	// Management chain
	index = findIdentity(st, cid)
	if index == -1 {
		return errors.New("Identity error")
	}

	if st.Identities[index].ManagementChainID == nil {
		return errors.New("Identity Error: No management chain found")
	}

	sEnts := getAllBetterIEBsFromChainID(st.Identities[index].ManagementChainID.String())
	for i := len(sEnts) - 1; i >= 0; i-- {
		e := sEnts[i]
		state.LoadIdentityByEntry(e.IEB, st, e.Height, false)
	}

	if verbose {
		fmt.Println("Identity is full:", st.Identities[index].IsFull())
		PrintIdentity(st.Identities[index])
	} else {
		fmt.Println(st.Identities[index].IsFull())
	}

	return nil
}

func findIdentity(st *state.State, chainID interfaces.IHash) int {
	for i := range st.Identities {
		if st.Identities[i].ManagementChainID.IsSameAs(chainID) {
			return i
		} else if st.Identities[i].IdentityChainID.IsSameAs(chainID) {
			return i
		}
	}
	return -1
}

func factomEntrysToIEBEntrys(ents []*factom.Entry) []interfaces.IEBEntry {
	x := make([]interfaces.IEBEntry, 0)
	for _, e := range ents {
		x = append(x, factomEntryToIEBEntry(e))
	}
	return x
}

func factomEntryToIEBEntry(ent *factom.Entry) interfaces.IEBEntry {
	iEnt := entryBlock.NewEntry()
	// Content
	bs := new(primitives.ByteSlice)
	bs.Bytes = ent.Content
	iEnt.Content = *bs
	// ChainID
	iEnt.ChainID, _ = primitives.HexToHash(ent.ChainID)
	// ExtIDs
	for _, ex := range ent.ExtIDs {
		exB := new(primitives.ByteSlice)
		exB.Bytes = ex
		iEnt.ExtIDs = append(iEnt.ExtIDs, *exB)
	}

	return iEnt
}

func PrintIdentity(i *state.Identity) {
	fmt.Print("Identity Chain: ", i.IdentityChainID, "\n")
	fmt.Print("Management Chain: ", i.ManagementChainID, "\n")
	fmt.Print("Matryoshka Hash: ", i.MatryoshkaHash, "\n")
	fmt.Print("Key 1: ", i.Key1, "\n")
	fmt.Print("Key 2: ", i.Key2, "\n")
	fmt.Print("Key 3: ", i.Key3, "\n")
	fmt.Print("Key 4: ", i.Key4, "\n")
	fmt.Print("Signing Key: ", i.SigningKey, "\n")
}
