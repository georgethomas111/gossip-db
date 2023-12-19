package node

import "time"

type Row struct {
	Value        []byte
	PutTimestamp time.Time
}

func NewRow(d []byte) *Row {
	return &Row{
		Value:        d,
		PutTimestamp: time.Now(),
	}
}

func MergeRows(row1 *Row, row2 *Row) *Row {
	if row1 == nil {
		return row2
	}

	if row2 == nil {
		return row1
	}

	if row1.PutTimestamp.After(row2.PutTimestamp) {
		return row1
	}

	return row2
}
