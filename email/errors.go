package email

// ErrNoRCPT indicates that the list of email recipients is
// empty, which is a violation of RFC 822.
type ErrNoRCPT struct {
}

func (e *ErrNoRCPT) Error() string {
	return "Need RCPT (recipient)"
}
