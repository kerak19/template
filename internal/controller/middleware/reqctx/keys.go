package reqctx

type key int

const (
	keyClaims key = iota
	keyUser
	keyLogger
)
