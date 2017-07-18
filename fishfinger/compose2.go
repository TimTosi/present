// START OMIT

#!/bin/sh

# INSERT BILLION OF DATA RIGHT HERE

# Advertise as ready.
echo "Listening to :9090"
while true ; do echo "ready" | nc -l 0.0.0.0 9090 ; done
echo "Test probe received."

# Nohup without nohup.
tail -f /dev/null

// END OMIT
