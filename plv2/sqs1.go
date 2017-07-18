// START OMIT
package queue

// NewSQSCustom create a new SQS with custom settings
func NewSQSCustom(sess *session.Session, queueName string, opts Options) (SqsInterface, error) {
	if opts.VisibilityTimeout != int64(0) {
	...
	return newSqs(sess, queueName, opts)
}

// NewSQS : Return an instance of SqsInterface
func NewSQS(sess *session.Session, queueName string) (SqsInterface, error) {
	defaultOptions := Options{
		...
	return newSqs(sess, queueName, defaultOptions)
}

// newSqs create the queue
func newSqs(sess *session.Session, queueName string, opts Options) (SqsInterface, error) {...}
// END OMIT
