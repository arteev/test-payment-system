package contextkey

// Key type from context key
type Key int

// Known keys
const (
	RequestIDKey Key = iota + 1
)

var (
	KeyNames = map[Key]string{
		RequestIDKey: "request_id",
	}
)