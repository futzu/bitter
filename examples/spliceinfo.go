package main

import "fmt"
import "github.com/futzu/gobit"

// SpInfo is the splice info section of the SCTE 35 cue.
type SpInfo struct {
	Name                   string
	TableId                string
	SectionSyntaxIndicator bool
	Private                bool
	Reserved               string
	SectionLength          uint16
	ProtocolVersion        uint8
	EncryptedPacket        bool
	EncryptionAlgorithm    uint8
	PtsAdjustment          float64
	CwIndex                string
	Tier                   string
	SpliceCommandLength    uint16
	SpliceCommandType      uint8
}

// Decode extracts bits for the splice info section values
func (spi *SpInfo) Decode(bitn *gobit.Bitn) {
	spi.Name = "Splice Info Section"
	spi.TableId = bitn.AsHex(8)
	spi.SectionSyntaxIndicator = bitn.AsBool()
	spi.Private = bitn.AsBool()
	spi.Reserved = bitn.AsHex(2)
	spi.SectionLength = bitn.AsUInt16(12)
	spi.ProtocolVersion = bitn.AsUInt8(8)
	spi.EncryptedPacket = bitn.AsBool()
	spi.EncryptionAlgorithm = bitn.AsUInt8(6)
	spi.PtsAdjustment = bitn.As90k(33)
	spi.CwIndex = bitn.AsHex(8)
	spi.Tier = bitn.AsHex(12)
	spi.SpliceCommandLength = bitn.AsUInt16(12)
	spi.SpliceCommandType = bitn.AsUInt8(8)
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
