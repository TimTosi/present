// START OMIT
func main() {
	c, err := fishfinger.NewCompose("./docker-compose.yaml")
	if err != nil {
		log.Fatal(err)
	}
    if err := c.Start(); err != nil { // This launches ALL compose file services
		log.Fatal(err)
	}
	if err := c.Start("redis"); err != nil { // This launches ONLY the service called redis
		log.Fatal(err)
	}
	if err := c.Stop(); err != nil {// This stops ALL compose file services
		log.Fatal(err)
	}
}
// END OMIT
