// START OMIT
package queue

// SqsInterface : what a sqs queue looks like
type SqsInterface interface {
	ReceiveMessages() ([]*sqs.Message, error)
	DeleteMessage(m *sqs.Message) error
	Send(subject string, data []byte, delay int64, traceID string) (*sqs.SendMessageOutput, error)
}
// END OMIT