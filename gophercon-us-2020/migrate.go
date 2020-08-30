sourceInstance, err := bindata.WithInstance(bindata.Resource(migrations.AssetNames(), migrations.Asset))
if err != nil {
	// handler err
}
targetInstance, err := postgres.WithInstance(db /* *sql.DB */, new(postgres.Config))
if err != nil {
	// handler err
}
m, err := migrate.NewWithInstance("go-bindata", sourceInstance, "postgres", targetInstance)
if err != nil {
	// handler err
}
err = m.Migrate(1) // Or whatever is the current version
