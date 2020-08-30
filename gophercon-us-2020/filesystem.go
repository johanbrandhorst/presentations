marshalled := fmt.Sprintf("%s, %d", users[0].Name, users[0].Age)

err := ioutil.WriteFile("./users.csv", []byte(marshalled), 0644)
if err != nil {
	// handle err
}
