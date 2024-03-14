package port

type RepositoryAtomicCallback func(r Repository) error

type Repository interface {
	Atomic(RepositoryAtomicCallback) error
	// Add other repository interface
}
