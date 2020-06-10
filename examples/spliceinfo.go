package main

import "fmt"
import "github.com/futzu/gobit"

// SpInfo is the splice info section of the SCTE 35 cue.
type SpInfo struct {
	TableId                string
	SectionSyntaxIndicator bool
	Private                bool
	Reserved               string
	SectionLength          uint64
	ProtocolVersion        uint64
	EncryptedPacket        bool
	EncryptionAlgorithm    uint64
	PtsAdjustment          float64
	CwIndex                string
	Tier                   string
	SpliceCommandLength    uint64
	SpliceCommandType      uint64
}

// Decode extracts bits for the splice info section values
func (spi *SpInfo) Decode(bitn *gobit.Bitn) {
	// gobit is always Big endian
	// bitslice 8 bits and return the hex value
	spi.TableId = bitn.AsHex(8)
	// bit slice off 1 bit and return a bool
	spi.SectionSyntaxIndicator = bitn.AsBool()
	spi.Private = bitn.AsBool()
	spi.Reserved = bitn.AsHex(2)
	// bitslice off 12 bits and return the value as a uint64
	spi.SectionLength = bitn.AsInt(12)
	spi.ProtocolVersion = bitn.AsInt(8)
	spi.EncryptedPacket = bitn.AsBool()
	spi.EncryptionAlgorithm = bitn.AsInt(6)
	// bitslice off 33 bits and return the 90k clock as float64
	spi.PtsAdjustment = bitn.As90k(33)
	spi.CwIndex = bitn.AsHex(8)
	spi.Tier = bitn.AsHex(12)
	spi.SpliceCommandLength = bitn.AsInt(12)
	spi.SpliceCommandType = bitn.AsInt(8)
}

func main() {
	// bites is the payload of a SCTE 35 packet
	bites := []byte("\xfc0/\x00\x00\x00\x00\x00\x00\xff\xff\xf0\x14\x05H\x00\x00\x8f\x7f\xef\xfesi\xc0.\xfe\x00R\xcc\xf5\x00\x00\x00\x00\x00\n\x00\x08CUEI\x00\x00\x015b\xdb\xa3")
	var bitn gobit.Bitn
	bitn.Load(bites)
	//SpInfo is the splice info section
	var spi SpInfo
	// extract the values
	spi.Decode(&bitn)
	fmt.Printf("%+v\n", spi)
}
