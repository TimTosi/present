// START OMIT

func main() {
	c, err := fishfinger.NewCompose("./docker-compose.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err := c.StartBackoff(fishfinger.SocketBackoff, "datastore"); err != nil {
		log.Fatal(err)
	}
}

// END OMIT
