package libmyna

import (
	"errors"

	"github.com/jpki/myna/asn1"
)

type CardInfoAP struct {
	reader *Reader
}

type CardFront struct {
	Header []byte `asn1:"private,tag:33"`
	Birth  string `asn1:"private,tag:34"`
	Age    string `asn1:"private,tag:35"`
	Tag36  []byte `asn1:"private,tag:36"`
	Name   []byte `asn1:"private,tag:37"`
	Addr   []byte `asn1:"private,tag:38"`
	Photo  []byte `asn1:"private,tag:39"`
	Tag40  []byte `asn1:"private,tag:40"`
	Expire string `asn1:"private,tag:41"`
	Code   []byte `asn1:"private,tag:42"`
}

func (self *CardInfoAP) LookupPinA() (int, error) {
	err := self.reader.SelectEF("00 13")
	if err != nil {
		return 0, err
	}
	count := self.reader.LookupPin()
	return count, nil
}

func (self *CardInfoAP) VerifyPinA(pin string) error {
	err := self.reader.SelectEF("00 13")
	if err != nil {
		return err
	}
	err = self.reader.Verify(pin)
	return err
}

func (self *CardInfoAP) LookupPinB() (int, error) {
	err := self.reader.SelectEF("00 12")
	if err != nil {
		return 0, err
	}
	count := self.reader.LookupPin()
	return count, nil
}

func (self *CardInfoAP) VerifyPinB(pin string) error {
	err := self.reader.SelectEF("00 12")
	if err != nil {
		return err
	}
	err = self.reader.Verify(pin)
	return err
}

func (self *CardInfoAP) GetCardFront() (*CardFront, error) {
	err := self.reader.SelectEF("00 02")
	if err != nil {
		return nil, err
	}
	data := self.reader.ReadBinary(7)
	if len(data) != 7 {
		return nil, errors.New("Error at ReadBinary()")
	}

	parser := ASN1PartialParser{}
	err = parser.Parse(data)
	if err != nil {
		return nil, err
	}
	data = self.reader.ReadBinary(parser.GetSize())

	var front CardFront
	_, err = asn1.UnmarshalWithParams(data, &front, "private,tag:32")
	if err != nil {
		return nil, err
	}
	return &front, nil
}
