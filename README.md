# gobit
[![Go Report Card](https://goreportcard.com/badge/github.com/futzu/gobiT)](https://goreportcard.com/report/github.com/futzu/gobiT)

### Bitslicer for bytes 

package gobit // import "github.com/futzu/gobit"
```
type Bitn struct {
        // Has unexported fields.
}
```

* ```func (b *Bitn) AsBool() bool```

 * ```func (b *Bitn) AsFloat(bitcount uint) float64```

* ```func (b *Bitn) AsHex(bitcount uint) string```

* ```func (b *Bitn) AsInt(bitcount uint) uint64```

* ```func (b *Bitn) Chunk(bitcount uint) uint64```

* ```func (b *Bitn) Load(bites []byte)```
