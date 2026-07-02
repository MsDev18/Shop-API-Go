package richerror

func (r *RichError) SetOp(op string) {
	r.op = op
}

func (r *RichError) SetMsg(message string) {
	r.message = message
}

func (r *RichError) SetKind(kind Kind) {
	r.kind = kind
}

func (r *RichError) SetErr(err error) {
	r.err = err
}

func (r *RichError) SetMeta(meta map[string]string) {
	r.meta = meta
}