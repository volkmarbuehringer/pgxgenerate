package db

import (
	"encoding/json"

	"github.com/jackc/pgx/pgtype"
)

type Bool struct {
	pgtype.Bool
}

func (v Bool) MarshalJSON() ([]byte, error) {
	if v.Status == pgtype.Present {

		return json.Marshal(v.Bool.Bool)

	} else {
		return json.Marshal(nil)
	}
}

type BoolArray = pgtype.BoolArray

type Bytea = pgtype.Bytea

type ByteaArray = pgtype.ByteaArray

type Date struct {
	pgtype.Date
}

type DateArray = pgtype.DateArray

type Decimal = pgtype.Decimal

type Float4 struct {
	pgtype.Float4
}

type Float4Array = pgtype.Float4Array

type Float8 struct {
	pgtype.Float8
}

type Float8Array = pgtype.Float8Array

type Int2 struct {
	pgtype.Int2
}

type Int2Array = pgtype.Int2Array

type Int4 struct {
	pgtype.Int4
}

func (v Int4) MarshalJSON() ([]byte, error) {
	if v.Status == pgtype.Present {
		return json.Marshal(v.Int)
	} else {
		return json.Marshal(nil)
	}
}

type Int4Array = pgtype.Int4Array

type Int8 struct {
	pgtype.Int8
}

type Int8Array = pgtype.Int8Array

type Interval struct {
	pgtype.Interval
}
type JSON = pgtype.JSON

type JSONB = pgtype.JSONB
type Numeric struct {
	pgtype.Numeric
}
type NumericArray = pgtype.NumericArray

type OID = pgtype.OID

type Text struct {
	pgtype.Text
}

func (v Text) MarshalJSON() ([]byte, error) {
	if v.Status == pgtype.Present {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

type TextArray = pgtype.TextArray

type Timestamp struct {
	pgtype.Timestamp
}

func (v Timestamp) MarshalJSON() ([]byte, error) {
	if v.Status == pgtype.Present {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

type TimestampArray = pgtype.TimestampArray

type Timestamptz struct {
	pgtype.Timestamptz
}

func (v Timestamptz) MarshalJSON() ([]byte, error) {
	if v.Status == pgtype.Present {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

type TimestamptzArray = pgtype.TimestamptzArray

type UUID struct {
	pgtype.UUID
}
type UUIDArray = pgtype.UUIDArray

type Varchar struct {
	pgtype.Varchar
}

func (v Varchar) MarshalJSON() ([]byte, error) {
	if v.Status == pgtype.Present {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

type VarcharArray = pgtype.VarcharArray
