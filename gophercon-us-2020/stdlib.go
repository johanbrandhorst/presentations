import _ "github.com/yourdatabase/driver"

rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age = $1", age)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	var name string
	if err := rows.Scan(&name); err != nil {
		log.Fatal(err)
	}
}
err = rows.Err()
// handle err
