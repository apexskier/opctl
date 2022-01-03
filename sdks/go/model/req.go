package model

// GetDataReq deprecated
type GetDataReq struct {
	DataRef string
}

type ListDescendantsReq struct {
	DataRef string `json:"dataRef"`
}

type StartOpReq struct {
	// map of args keyed by input name
	Args map[string]*Value `json:"args,omitempty"`
	// Op details the op to start
	Op StartOpReqOp `json:"op,omitempty"`
}

type StartOpReqOp struct {
	Ref string
}
