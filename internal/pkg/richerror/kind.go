package richerror

type Kind uint

const (
	KindUnexpectedErr Kind = iota + 1
	KindConflictErr
	KindNotFoundErr
	KindInternalServerErr
	KindForbiddenErr
	KindUnauthorizeErr
	KindBadRequestErr
)
