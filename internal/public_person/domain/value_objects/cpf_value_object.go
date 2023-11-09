package valueobjects

import (
	"encoding/json"

	"github.com/Nhanderu/brdoc"
	"github.com/intwone/eda-arch-golang/internal/private_shared/err"
)

type Cpf struct {
	Value string
}

func NewCpf(value string) (*Cpf, error) {
	if !brdoc.IsCPF(value) {
		return nil, err.NewInvalidValueError("cpf")
	}
	cpf := Cpf{Value: value}
	return &cpf, nil
}

func (e *Cpf) UnmarshalJSON(data []byte) error {
	var cpf string
	if err := json.Unmarshal(data, &cpf); err != nil {
		return err
	}
	if !brdoc.IsCPF(cpf) {
		return err.NewInvalidValueError("cpf")
	}
	e.Value = cpf
	return nil
}
