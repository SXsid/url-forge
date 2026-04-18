package repository

type Transaction interface {
	Commit() error
	Rollback() error
}
