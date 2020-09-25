package models

import (
	"database/sql"
	"encoding/json"
)

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullString
//overrided to flat the sql NullString map
func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return []byte("null"), nil
}

// MarshalJSON for NullInt64
//overrided to flat the sql NullInt64 map
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return []byte("null"), nil
}
