package bitter

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

// Bitn converts bytes to a list of bits.
type Bitn struct {
	idx  uint
	bits string
}

// Load raw bytes and convert to bits
func (b *Bitn) Load(bites []byte) {
	i := new(big.Int)
	i.SetBytes(bites)
	b.bits = fmt.Sprintf("%b", i)
	b.idx = 0
}

// Chunk slices bitcount of bits and returns it as a uint64
func (b *Bitn) Chunk(bitcount uint) uint64 {
	j := new(big.Int)
	d := b.idx + bitcount
	j.SetString(b.bits[b.idx:d], 2)
	b.idx = d
	return j.Uint64()
}

// AsUInt64 is a wrapper for Chunk
func (b *Bitn) AsUInt64(bitcount uint) uint64 {
	return b.Chunk(bitcount)

}

// AsBool slices 1 bit and returns true for 1 , false for 0
func (b *Bitn) AsBool() bool {
	var bitcount uint
	bitcount = 1
	return (b.Chunk(bitcount) == 1)

}

// AsFloat slices bitcount of bits and returns as float64
func (b *Bitn) AsFloat(bitcount uint) float64 {
	asfloat := float64(b.Chunk(bitcount))
	return asfloat
}

// As90k is AsFloat / 90000.00 rounded to six decimal places.
func (b *Bitn) As90k(bitcount uint) float64 {
	as90k := b.AsFloat(bitcount) / 90000.00
	return float64(uint64(as90k*1000000)) / 1000000
}

// AsHex slices bitcount of bits and returns as hex string
func (b *Bitn) AsHex(bitcount uint) string {
	ashex := fmt.Sprintf("%#x", b.Chunk(bitcount))
	return ashex
}

// AsDeHex slices bitcount of bits and returns as hex string
func (b *Bitn) AsDeHex(bitcount uint) []byte {
	ashex := fmt.Sprintf("%x", b.Chunk(bitcount))
	asdehex, err := hex.DecodeString(ashex)
	if err != nil {
		panic(err)
	}
	return asdehex
}

// Forward advances b.idx by bitcount
func (b *Bitn) Forward(bitcount uint) {
	b.idx += bitcount
}
