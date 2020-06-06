
package gobit

import (
	"fmt"
	"math/big"
)

type Bitn struct {
	idx  uint
	bits string
}

// load raw bytes and convert to bits
func (b *Bitn) Load(bites []byte) {
	i := new(big.Int)
	i.SetBytes(bites)
	b.bits = fmt.Sprintf("%b", i)
	b.idx = 0
}

// slices bitcount of bits and returns it as a uint64
func (b *Bitn) Chunk(bitcount uint) uint64 {
	j := new(big.Int)
	d := b.idx + bitcount
	j.SetString(b.bits[b.idx:d], 2)
	b.idx = d
	fmt.Printf("bitidx: %v\n", b.idx)
	return j.Uint64()
}

// wrapper for Chunk
func (b *Bitn) AsInt(bitcount uint) uint64 {
	asint := b.Chunk(bitcount)
	fmt.Println(asint)
	return asint
}

// slices 1 bit and returns true for 1 , false for 0
func (b *Bitn) AsBool() bool {
	var bitcount uint
	bitcount = 1
	boo := (b.Chunk(bitcount) == 1)
	fmt.Println(boo)
	return boo
}

// slices bitcount of bits and returns as float64
func (b *Bitn) AsFloat(bitcount uint) float64 {
	asfloat := float64(b.Chunk(bitcount))
	fmt.Printf("%.6f\n", asfloat)
	return asfloat
}

// slices bitcount of bits and returns as hex string
func (b *Bitn) AsHex(bitcount uint) string {
	ashex := fmt.Sprintf("%#x", b.Chunk(bitcount))
	fmt.Println(ashex)
	return ashex
}
