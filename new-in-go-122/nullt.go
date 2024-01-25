package main

import "database/sql"

func main() {
	oldInt := new(int)         // *int
	nullInt := sql.NullInt64{} // sql.NullInt64
	newInt := sql.Null[int]{}  // Works for any integer/string/byte type
	_ = rows.Scan(&oldInt, &nullInt, &newInt)
}
