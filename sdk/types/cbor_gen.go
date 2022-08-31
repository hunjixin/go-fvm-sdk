// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package types

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufAggregateSealVerifyInfo = []byte{133}

func (t *AggregateSealVerifyInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufAggregateSealVerifyInfo); err != nil {
		return err
	}

	// t.Number (abi.SectorNumber) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Number)); err != nil {
		return err
	}

	// t.Randomness (abi.SealRandomness) (slice)
	if len(t.Randomness) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Randomness was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Randomness))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Randomness[:]); err != nil {
		return err
	}

	// t.InteractiveRandomness (abi.InteractiveSealRandomness) (slice)
	if len(t.InteractiveRandomness) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.InteractiveRandomness was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.InteractiveRandomness))); err != nil {
		return err
	}

	if _, err := cw.Write(t.InteractiveRandomness[:]); err != nil {
		return err
	}

	// t.SealedCID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.SealedCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.SealedCID: %w", err)
	}

	// t.UnsealedCID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.UnsealedCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.UnsealedCID: %w", err)
	}

	return nil
}

func (t *AggregateSealVerifyInfo) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AggregateSealVerifyInfo{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Number (abi.SectorNumber) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Number = abi.SectorNumber(extra)

	}
	// t.Randomness (abi.SealRandomness) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Randomness: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Randomness = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Randomness[:]); err != nil {
		return err
	}
	// t.InteractiveRandomness (abi.InteractiveSealRandomness) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.InteractiveRandomness: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.InteractiveRandomness = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.InteractiveRandomness[:]); err != nil {
		return err
	}
	// t.SealedCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.SealedCID: %w", err)
		}

		t.SealedCID = c

	}
	// t.UnsealedCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.UnsealedCID: %w", err)
		}

		t.UnsealedCID = c

	}
	return nil
}

var lengthBufAggregateSealVerifyProofAndInfos = []byte{133}

func (t *AggregateSealVerifyProofAndInfos) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufAggregateSealVerifyProofAndInfos); err != nil {
		return err
	}

	// t.Miner (abi.ActorID) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Miner)); err != nil {
		return err
	}

	// t.SealProof (abi.RegisteredSealProof) (int64)
	if t.SealProof >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.SealProof)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.SealProof-1)); err != nil {
			return err
		}
	}

	// t.AggregateProof (abi.RegisteredAggregationProof) (int64)
	if t.AggregateProof >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.AggregateProof)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.AggregateProof-1)); err != nil {
			return err
		}
	}

	// t.Proof ([]uint8) (slice)
	if len(t.Proof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Proof))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Proof[:]); err != nil {
		return err
	}

	// t.Infos ([]types.AggregateSealVerifyInfo) (slice)
	if len(t.Infos) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Infos was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Infos))); err != nil {
		return err
	}
	for _, v := range t.Infos {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}
	return nil
}

func (t *AggregateSealVerifyProofAndInfos) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AggregateSealVerifyProofAndInfos{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (abi.ActorID) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Miner = abi.ActorID(extra)

	}
	// t.SealProof (abi.RegisteredSealProof) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SealProof = abi.RegisteredSealProof(extraI)
	}
	// t.AggregateProof (abi.RegisteredAggregationProof) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.AggregateProof = abi.RegisteredAggregationProof(extraI)
	}
	// t.Proof ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Proof = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Proof[:]); err != nil {
		return err
	}
	// t.Infos ([]types.AggregateSealVerifyInfo) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Infos: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Infos = make([]AggregateSealVerifyInfo, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v AggregateSealVerifyInfo
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.Infos[i] = v
	}

	return nil
}

var lengthBufReplicaUpdateInfo = []byte{133}

func (t *ReplicaUpdateInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufReplicaUpdateInfo); err != nil {
		return err
	}

	// t.UpdateProofType (abi.RegisteredUpdateProof) (int64)
	if t.UpdateProofType >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.UpdateProofType)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.UpdateProofType-1)); err != nil {
			return err
		}
	}

	// t.OldSealedSectorCID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.OldSealedSectorCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.OldSealedSectorCID: %w", err)
	}

	// t.NewSealedSectorCID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.NewSealedSectorCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.NewSealedSectorCID: %w", err)
	}

	// t.NewUnsealedSectorCID (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.NewUnsealedSectorCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.NewUnsealedSectorCID: %w", err)
	}

	// t.Proof ([]uint8) (slice)
	if len(t.Proof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Proof))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Proof[:]); err != nil {
		return err
	}
	return nil
}

func (t *ReplicaUpdateInfo) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ReplicaUpdateInfo{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.UpdateProofType (abi.RegisteredUpdateProof) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.UpdateProofType = abi.RegisteredUpdateProof(extraI)
	}
	// t.OldSealedSectorCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.OldSealedSectorCID: %w", err)
		}

		t.OldSealedSectorCID = c

	}
	// t.NewSealedSectorCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.NewSealedSectorCID: %w", err)
		}

		t.NewSealedSectorCID = c

	}
	// t.NewUnsealedSectorCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.NewUnsealedSectorCID: %w", err)
		}

		t.NewUnsealedSectorCID = c

	}
	// t.Proof ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Proof = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Proof[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufInstallParams = []byte{129}

func (t *InstallParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufInstallParams); err != nil {
		return err
	}

	// t.Code ([]uint8) (slice)
	if len(t.Code) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Code was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Code))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Code[:]); err != nil {
		return err
	}
	return nil
}

func (t *InstallParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = InstallParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Code ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Code: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Code = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Code[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufInstallReturn = []byte{130}

func (t *InstallReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufInstallReturn); err != nil {
		return err
	}

	// t.CodeCid (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.CodeCid); err != nil {
		return xerrors.Errorf("failed to write cid field t.CodeCid: %w", err)
	}

	// t.Installed (bool) (bool)
	if err := cbg.WriteBool(w, t.Installed); err != nil {
		return err
	}
	return nil
}

func (t *InstallReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = InstallReturn{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.CodeCid (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.CodeCid: %w", err)
		}

		t.CodeCid = c

	}
	// t.Installed (bool) (bool)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.Installed = false
	case 21:
		t.Installed = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	return nil
}
