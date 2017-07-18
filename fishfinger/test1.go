// START OMIT

// setUp is an helper function executed before any test.
func setUp(composeFile string, services ...string) *fishfinger.Compose {
	c, err := fishfinger.NewCompose(composeFile)
	if err != nil {
		log.Fatal(err)
	}
	if err = c.Start(services...); err != nil {
		log.Fatal(err)
	}
	return c
}

// tearDown is an helper function executed after all tests.
func tearDown(c *fishfinger.Compose) error { return c.Stop() }

// END OMIT
