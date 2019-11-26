package api

import "github.com/shopspring/decimal"

type Decimal struct {
	decimal.Decimal
}

func (d *Decimal) SetFromString(data string) error {
	dec, err := decimal.NewFromString(data)
	if err != nil {
		return err
	}
	d.Decimal = dec
	return nil
}
