package hasher

type HasherInterface interface {
	Hash(value string) (*string, error)
	Compare(value string, hashed string) bool
}
