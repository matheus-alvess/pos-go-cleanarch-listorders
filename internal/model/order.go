package model

type Order struct {
	IDField    int32   `json:"id"`
	PriceField float64 `json:"price"`
	TaxField   float64 `json:"tax"`
}

func (o *Order) ID() int32 {
	return o.IDField
}

func (o *Order) Price() float64 {
	return o.PriceField
}

func (o *Order) Tax() float64 {
	return o.TaxField
}
