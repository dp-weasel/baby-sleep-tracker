package domain

// DomainError represents a semantic error in the domain.
type DomainError string

func (e DomainError) Error() string { return string(e) }

var (
	ErrInvalidFirstEvent   = DomainError("first event must be SUEÑO_FIN")
	ErrSameTimestamp       = DomainError("event with same timestamp already exists")
	ErrOutOfOrder          = DomainError("event timestamp is out of order")
	ErrConsecutiveSameType = DomainError("consecutive events of same type are not allowed")
	ErrInsufficientEvents  = DomainError("not enough events to derive periods")
)
