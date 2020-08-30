connConfig, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
if err != nil {
	// handle error
}

connConfig.Logger = myLogger // supports logrus, zap, zerolog, log15

db := stdlib.OpenDB(*connConfig)

db.QueryRow("select * from users where id=$1", userID)
