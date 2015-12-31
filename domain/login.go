package domain

type LoginVM struct {
	Errors []interface{}
}

func NewLoginVM(errors ...interface{}) *LoginVM {
	return &LoginVM{Errors: errors}
}
