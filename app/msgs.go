package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	cbd "github.com/cybercongress/cyberd/app/types"
	"github.com/ipfs/go-cid"
)

type MsgLink struct {
	Address sdk.AccAddress `json:"address"`
	From    cbd.Cid        `json:"cid1"`
	To      cbd.Cid        `json:"cid2"`
}

var _ sdk.Msg = MsgLink{}

func NewMsgLink(address sdk.AccAddress, fromCid cbd.Cid, toCid cbd.Cid) MsgLink {
	return MsgLink{Address: address, From: fromCid, To: toCid}
}

func (msg MsgLink) Name() string { return "link" }

func (MsgLink) Route() string { return "link" }
func (MsgLink) Type() string  { return "link" }

func (msg MsgLink) ValidateBasic() sdk.Error {

	if len(msg.Address) == 0 {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}

	if _, err := cid.Decode(string(msg.From)); err != nil {
		return cbd.ErrInvalidCid()
	}

	if _, err := cid.Decode(string(msg.To)); err != nil {
		return cbd.ErrInvalidCid()
	}

	return nil
}

func (msg MsgLink) GetSignBytes() []byte {
	b, err := msgCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgLink) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}