// Code generated by protoapi:go; DO NOT EDIT.

package todolistsvr

// AddReq
type AddReq struct {
	Item *Todo `json:"item"`
}

func (r *AddReq) GetItem() *Todo {
	if r == nil {
		var zeroVal *Todo
		return zeroVal
	}
	return r.Item
}
