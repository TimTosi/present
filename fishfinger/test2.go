// START OMIT

func TestNewRedis(t *testing.T) {
	testCases := []struct {
		name              string
		mockServerAddr    string
		expectedErrorFunc func(ensure.Fataler, interface{}, ...interface{})
	}{
		{"regular", "0.0.0.0:5101", ensure.Nil},
		{"invalid", "256.256.256.256", ensure.NotNil},
	}

	// SETUP HERE
	c := setUp("$GOPATH/$PROJECT/docker-compose.yml", "redis")
	// TEAR DOWN HERE
	defer func() { ensure.Nil(t, tearDown(c)) }()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := NewRedis(tc.mockServerAddr, 0*time.Second)
			tc.expectedErrorFunc(t, err)
		})
	}
}

// END OMIT
