package domain

type AccountCreateVM struct {
	Errors []interface{}
}

func NewAccountCreateVM(errors ...interface{}) *AccountCreateVM {
	return &AccountCreateVM{Errors: errors}
}
