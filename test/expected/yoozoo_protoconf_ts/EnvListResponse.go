// Code generated by protoapi; DO NOT EDIT.

package yoozoo_protoconf_ts

// EnvListResponse
type EnvListResponse struct {
    Envs []*Env `json:"envs"`
}

func (r EnvListResponse) Validate() *ValidateError {
    errs := []*FieldError{}
    if len(errs) > 0 {
		return &ValidateError{Errors: errs}
	}
	return nil
}