import _ "github.com/yourdatabase/driver" // Blank import forces driver to use init()
//                                ,-- careful not to let user control any part of this
rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age = $1", age)
if err != nil { //                                                         ^-- interface{} type
	log.Fatal(err)
}
defer rows.Close() // Don't forget to check this error or you might miss a failure
for rows.Next() {
	var name string
	if err := rows.Scan(&name); err != nil {
		log.Fatal(err)
	}
}
err = rows.Err() // Don't forget to call this!
// handle err
