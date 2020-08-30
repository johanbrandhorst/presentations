// Get a connection from the *sql.DB connection pool
conn, err := db.Conn(context.Background())
// handle error
defer conn.Close()
err = conn.Raw(func(driverConn interface{}) error {
	conn := driverConn.(*stdlib.Conn).Conn() // conn is a *pgx.Conn
	// Do pgx specific stuff with conn
	return nil
})
// handle error
