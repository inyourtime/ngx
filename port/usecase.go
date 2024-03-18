package port

type Usecase interface {
	Auth() AuthUsecase
	User() UserUsecase
}
