q := sb.Select("id", "name", "age", "create_time").From("users")

// var createdSince time.Time
if !createdSince.IsZero() {
	q = q.Where(squirrel.Gt{
		"create_time": createdSince, // use time.Time directly
	})
}

// var olderThan time.Duration
if olderThan != 0 {
	i := &pgtype.Interval{}
	_ = i.Set(olderThan)
	q = q.Where( // Automatically ANDs to existing WHERE statements
		squirrel.Expr(
			"CURRENT_TIMESTAMP - create_time > ?", i, // use pgtype.Interval
		),
	)
}
