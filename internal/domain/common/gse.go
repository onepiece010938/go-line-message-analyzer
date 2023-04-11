package common

import (
	"github.com/go-ego/gse"
)


type Gse struct {
	Seg gse.Segmenter
}

func NewGse() *Gse {
	
	new(Gse).Seg.LoadDict()
	seg.LoadDict()
	return &Gse{
		Seg: ,
	}
}
