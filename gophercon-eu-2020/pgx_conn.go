// Get a connection from the *sql.DB connection pool
conn, err := db.Conn(context.Background())
if err != nil {
	// handle err
}
defer func() {
	cerr := conn.Close()
	if err == nil {
		err = cerr
	}
}()

// Get the underlying driver connection
// Go 1.13 and above
err = conn.Raw(func(driverConn interface{}) error {
	conn := driverConn.(*stdlib.Conn).Conn() // conn is a *pgx.Conn
	// Do pgx specific stuff with conn
	conn.CopyFrom(...)
	return nil
})
if err != nil {
	// handle err
}
