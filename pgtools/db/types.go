package db

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/jackc/pgx/pgtype"
)

type Bool struct {
	pgtype.Bool
}

func (v *Bool) SetBool(b bool) {
	v.Status = pgtype.Present
	v.Bool.Bool = b

}

func (v *Bool) Stringer() string {

	return strconv.FormatBool(v.Bool.Bool)

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

func (v *Date) Stringer() string {

	return v.Time.Format(time.RFC3339)

}

type DateArray = pgtype.DateArray

type Decimal = pgtype.Decimal

type Float4 struct {
	pgtype.Float4
}

func (v *Float4) Stringer() string {

	return strconv.FormatFloat(float64(v.Float), 'f', 10, 32)

}

type Float4Array = pgtype.Float4Array

type Float8 struct {
	pgtype.Float8
}

func (v *Float8) SetFloat(b float64) {
	v.Status = pgtype.Present
	v.Float = b

}

func (v *Float8) Stringer() string {

	return strconv.FormatFloat(v.Float, 'f', 10, 64)

}

type Float8Array = pgtype.Float8Array

type Int2 struct {
	pgtype.Int2
}

func (v *Int2) Stringer() string {

	return strconv.FormatInt(int64(v.Int), 10)

}

type Int2Array = pgtype.Int2Array

type Int4 struct {
	pgtype.Int4
}

func (v *Int4) Stringer() string {

	return strconv.FormatInt(int64(v.Int), 10)

}

func (v *Int4) SetInt(b int32) {
	v.Status = pgtype.Present
	v.Int = b

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

func (v *Int8) SetInt(b int) {
	v.Status = pgtype.Present
	v.Int = int64(b)

}

func (v *Int8) Stringer() string {

	return strconv.FormatInt(v.Int, 10)

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

func (v *Numeric) Stringer() string {

	return v.Numeric.Int.String()

}

type NumericArray = pgtype.NumericArray

type OID = pgtype.OID

type Text struct {
	pgtype.Text
}

func (v *Text) SetText(b string) {
	v.Status = pgtype.Present
	v.Text.String = b

}

func (v *Text) Stringer() string {

	return v.Text.String

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

func (v *Timestamp) SetTimestamp(b time.Time) {
	v.Status = pgtype.Present
	v.Time = b

}

func (v *Timestamp) Stringer() string {

	return v.Time.Format(time.RFC3339)

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

func (v *Timestamptz) SetTimestamp(b time.Time) {
	v.Status = pgtype.Present
	v.Time = b

}

func (v *Timestamptz) Stringer() string {

	return v.Time.Format(time.RFC3339)

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

func (v *Varchar) SetVarchar(b string) {
	v.Status = pgtype.Present
	v.Varchar.String = b

}

func (v *Varchar) Stringer() string {

	return v.Varchar.String

}

func (v Varchar) MarshalJSON() ([]byte, error) {
	if v.Status == pgtype.Present {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

type VarcharArray = pgtype.VarcharArray
