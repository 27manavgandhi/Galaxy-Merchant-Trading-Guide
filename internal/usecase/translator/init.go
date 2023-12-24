package translator

import "github.com/27manavgandhi/galaxy-merchant-trading/internal/model"

type Translator struct {
	dictionary model.Dictionary
	priceList  model.PriceList
}

func New() *Translator {
	return &Translator{
		dictionary: make(model.Dictionary),
		priceList:  make(model.PriceList),
	}
}
