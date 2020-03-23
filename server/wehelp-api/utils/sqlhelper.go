package utils

import "database/sql"

func SetNullString(input string) sql.NullString {
	if input == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: input, Valid: true}
}

func SetNullInt32(input int32) sql.NullInt32 {
	if input == 0 {
		return sql.NullInt32{Int32: 0, Valid: false}
	}
	return sql.NullInt32{Int32: input, Valid: true}
}

func SetNullInt64(input int64) sql.NullInt64 {
	if input == 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: input, Valid: true}
}
