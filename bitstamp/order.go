package bitstamp

import (
	"encoding/json"
	"github.com/mattomatic/go-bitcoin/common"
)

type Order struct {
	price  string
	volume string
}

func (o *Order) GetExchange() common.Exchange { return common.Exchange(ExchangeId) }
func (o *Order) GetSymbol() common.Symbol     { return common.Symbol("BTC") }
func (o *Order) GetPrice() common.Price       { return getPrice(o.price) }
func (o *Order) GetVolume() common.Volume     { return getVolume(o.volume) }

func (o *Order) UnmarshalJSON(bytes []byte) error {
	values := &[2]string{}
	err := json.Unmarshal(bytes, values)

	if err != nil {
		panic(err.Error())
	}

	o.price = values[0]
	o.volume = values[1]

	return err
}
