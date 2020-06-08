# gobit
[![Go Report Card](https://goreportcard.com/badge/github.com/futzu/gobiT)](https://goreportcard.com/report/github.com/futzu/gobiT)

### Bitslicer for bytes 
```go
$ go doc gobit.Bitn

package gobit // import "github.com/futzu/gobit"

type Bitn struct {
        // Has unexported fields.
}
    Bitn conversts bytes to a list of bits.

func (b *Bitn) As90k(bitcount uint) float64
func (b *Bitn) AsBool() bool
func (b *Bitn) AsFloat(bitcount uint) float64
func (b *Bitn) AsHex(bitcount uint) string
func (b *Bitn) AsInt(bitcount uint) uint64
func (b *Bitn) Chunk(bitcount uint) uint64
func (b *Bitn) Forward(bitcount uint)
func (b *Bitn) Load(bites []byte)
```
