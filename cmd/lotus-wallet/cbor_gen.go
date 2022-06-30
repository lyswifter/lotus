package main

// import (
// 	"fmt"
// 	"io"

// 	"github.com/filecoin-project/go-state-types/abi"
// 	"golang.org/x/xerrors"
// )

// var lengthBufProposeParams = []byte{132}

// func (t *ProposeParams) MarshalCBOR(w io.Writer) error {
// 	if t == nil {
// 		_, err := w.Write(cbg.CborNull)
// 		return err
// 	}
// 	if _, err := w.Write(lengthBufProposeParams); err != nil {
// 		return err
// 	}

// 	scratch := make([]byte, 9)

// 	// t.To (address.Address) (struct)
// 	if err := t.To.MarshalCBOR(w); err != nil {
// 		return err
// 	}

// 	// t.Value (big.Int) (struct)
// 	if err := t.Value.MarshalCBOR(w); err != nil {
// 		return err
// 	}

// 	// t.Method (abi.MethodNum) (uint64)

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Method)); err != nil {
// 		return err
// 	}

// 	// t.Params ([]uint8) (slice)
// 	if len(t.Params) > cbg.ByteArrayMaxLen {
// 		return xerrors.Errorf("Byte array in field t.Params was too long")
// 	}

// 	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Params))); err != nil {
// 		return err
// 	}

// 	if _, err := w.Write(t.Params[:]); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (t *ProposeParams) UnmarshalCBOR(r io.Reader) error {
// 	*t = ProposeParams{}

// 	br := cbg.GetPeeker(r)
// 	scratch := make([]byte, 8)

// 	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
// 	if err != nil {
// 		return err
// 	}
// 	if maj != cbg.MajArray {
// 		return fmt.Errorf("cbor input should be of type array")
// 	}

// 	if extra != 4 {
// 		return fmt.Errorf("cbor input had wrong number of fields")
// 	}

// 	// t.To (address.Address) (struct)

// 	{

// 		if err := t.To.UnmarshalCBOR(br); err != nil {
// 			return xerrors.Errorf("unmarshaling t.To: %w", err)
// 		}

// 	}
// 	// t.Value (big.Int) (struct)

// 	{

// 		if err := t.Value.UnmarshalCBOR(br); err != nil {
// 			return xerrors.Errorf("unmarshaling t.Value: %w", err)
// 		}

// 	}
// 	// t.Method (abi.MethodNum) (uint64)

// 	{

// 		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
// 		if err != nil {
// 			return err
// 		}
// 		if maj != cbg.MajUnsignedInt {
// 			return fmt.Errorf("wrong type for uint64 field")
// 		}
// 		t.Method = abi.MethodNum(extra)

// 	}
// 	// t.Params ([]uint8) (slice)

// 	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
// 	if err != nil {
// 		return err
// 	}

// 	if extra > cbg.ByteArrayMaxLen {
// 		return fmt.Errorf("t.Params: byte array too large (%d)", extra)
// 	}
// 	if maj != cbg.MajByteString {
// 		return fmt.Errorf("expected byte array")
// 	}

// 	if extra > 0 {
// 		t.Params = make([]uint8, extra)
// 	}

// 	if _, err := io.ReadFull(br, t.Params[:]); err != nil {
// 		return err
// 	}
// 	return nil
// }
