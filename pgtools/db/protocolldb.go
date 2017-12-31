package db

import (
	"encoding/binary"
	"fmt"

	"github.com/jackc/pgx/pgtype"
)

func Helper(ci *pgtype.ConnInfo, src []byte, helper func() []pgtype.BinaryDecoder) error {
	if src == nil {
		//fmt.Println("hier null")
		return nil
	}

	var arrayHeader pgtype.ArrayHeader
	rp, err := arrayHeader.DecodeBinary(ci, src)
	if err != nil {
		return err
	}
	var elementCount int32
	if len(arrayHeader.Dimensions) == 0 {

		elementCount = 0
	} else {
		elementCount = arrayHeader.Dimensions[0].Length
		for _, d := range arrayHeader.Dimensions[1:] {
			elementCount *= d.Length
		}
	}

	//fmt.Println("testarrrr", len(src), arrayHeader, elementCount, rp)

	var i int32
	for i = 0; i < elementCount; i++ {
		d := helper()

		elemLen := int(int32(binary.BigEndian.Uint32(src[rp:])))
		rp += 4
		var elemSrc []byte
		if elemLen >= 0 {
			elemSrc = src[rp : rp+elemLen]
			rp += elemLen
		}

		if err = DecodeBinary(ci, d, elemSrc); err != nil {
			return err
		}

	}

	return nil
}

func DecodeBinary(ci *pgtype.ConnInfo, fields []pgtype.BinaryDecoder, src []byte) error {

	//fmt.Println("testtext1111", len(src))

	rp := 0

	if len(src[rp:]) < 4 {
		return fmt.Errorf("Record incomplete %v", src)
	}
	fieldCount := int(int32(binary.BigEndian.Uint32(src[rp:])))
	rp += 4

	if fieldCount != len(fields) {
		return fmt.Errorf("Mismatch number of fields %d %d", fieldCount, len(fields))
	}

	for i := 0; i < fieldCount; i++ {
		if len(src[rp:]) < 8 {
			return fmt.Errorf("Record incomplete %v", src)
		}
		//fieldOid := pgtype.OID(binary.BigEndian.Uint32(src[rp:]))
		rp += 4

		fieldLen := int(int32(binary.BigEndian.Uint32(src[rp:])))
		rp += 4
		/*
			if dt, ok := ci.DataTypeForOID(fieldOid); ok {
				//	var binaryDecoder pgtype.BinaryDecoder
				if _, ok := dt.Value.(pgtype.BinaryDecoder); !ok {
					//	return fmt.Errorf("unknown oid while decoding record: %v", fieldOid)
					fmt.Printf("unknown oid while decoding record: %v", fieldOid)
				}

			}*/

		var fieldBytes []byte
		if fieldLen >= 0 {
			if len(src[rp:]) < fieldLen {
				return fmt.Errorf("Record incomplete %v", src)
			}
			fieldBytes = src[rp : rp+fieldLen]
			rp += fieldLen
		}

		if err := (fields)[i].DecodeBinary(ci, fieldBytes); err != nil {
			return err
		}
		//	fmt.Println("gut bis hier", i, (fields)[i])

		/*
			if err := binaryDecoder.DecodeBinary(ci, fieldBytes); err != nil {
				return err
			}

			ga := binaryDecoder.(pgtype.Value)
			la := ga.Get()
			if la != nil {
				fmt.Printf("bis hier %T %T %v %v\n", la, (fields)[i], la, (fields)[i])
				err := ga.AssignTo((fields)[i])

				if err != nil {
					fmt.Println("hier", err)
					return err
				}
				//	fmt.Println("feld", i, ga, (*fields)[i])
			}
		*/
	}

	//	fmt.Println("url", fields[1].Get(), fields[1])
	//fmt.Printf("str %s %s\n", *stru.Url, *stru.W_cr_uid)
	//*dst = *stru
	return nil
}
