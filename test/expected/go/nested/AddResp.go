// Code generated by protoapi:go; DO NOT EDIT.

package nested

// AddResp
type AddResp struct {
	Result int    `json:"result"`
	Extra  *Extra `json:"extra"`
}

func (r *AddResp) GetResult() int {
	if r == nil {
		var zeroVal int
		return zeroVal
	}
	return r.Result
}

func (r *AddResp) GetExtra() *Extra {
	if r == nil {
		var zeroVal *Extra
		return zeroVal
	}
	return r.Extra
}
