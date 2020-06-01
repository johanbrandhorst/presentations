sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // Use postgres placeholders
name := "Johan" // External variable
q := sb.Select(
	"*",
).From(
	"users",
).Where(
	squirrel.Eq{"name": name}, // Becomes positional parameter
)
query, args, err := q.ToSql()
if err != nil {
	// handle err
}
db.Query(query, args...) // SELECT * FROM users WHERE name = $1, [Johan]
