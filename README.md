# bitter     [![Go Report Card](https://goreportcard.com/badge/github.com/FUTZU/bitter)](https://goreportcard.com/report/github.com/FUTZU/bitter)

### Bitslicer for bytes 
```js

$ package bitter // import "github.com/futzu/bitter"


TYPES

type Bitn struct {
	idx  uint
	bits string
}
    Bitn converts bytes to a list of bits.

func (b *Bitn) As90k(bitcount uint) float64
    As90k is AsFloat / 90000.00 rounded to six decimal places.

func (b *Bitn) AsAscii(bitcount uint) string
    AsAscii returns the Ascii chars of AsBytes

func (b *Bitn) AsBool() bool
    AsBool slices 1 bit and returns true for 1 , false for 0

func (b *Bitn) AsBytes(bitcount uint) []byte
    AsBytes slices bitcount of bits and returns as []bytes

func (b *Bitn) AsFloat(bitcount uint) float64
    AsFloat slices bitcount of bits and returns as float64

func (b *Bitn) AsHex(bitcount uint) string
    AsHex slices bitcount of bits and returns as hex string

func (b *Bitn) AsUInt64(bitcount uint) uint64
    AsUInt64 is a wrapper for Chunk

func (b *Bitn) AsUInt8(bitcount uint) uint8
    AsUInt8 trims AsUInt64 to 8 bits for smaller numbers

func (b *Bitn) Chunk(bitcount uint) *big.Int
    Chunk slices bitcount of bits and returns it as a uint64

func (b *Bitn) Forward(bitcount uint)
    Forward advances b.idx by bitcount

func (b *Bitn) Load(bites []byte)
    Load raw bytes and convert to bits


```
