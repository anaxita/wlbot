package entity

type Mikrotik struct {
	ID        int64
	Address   string
	Login     string
	Password  string
	DefaultWL string
	IsDefault bool
}
