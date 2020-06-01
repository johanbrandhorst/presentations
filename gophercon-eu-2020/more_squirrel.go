q := sb.Insert(
	"users",
).SetMap(map[string]interface{}{
	// Column: Value
	"name": "Johan",
	"age":  28,
}).Suffix(
	"RETURNING id, name, age, create_time",
)

// Use with *sql.DB directly
row := q.RunWith(db).QueryContext(ctx)
