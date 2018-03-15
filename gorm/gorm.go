//go:generate protoc -I../proto -I ${GOPATH}/src -I ${GOPATH}/src/github.com/golang --go_out=plugins=grpc:. ../proto/fsme.proto
// 
//if following fails, (and because "go get github.com/infobloxopen/protoc-gen-gorm" fails to import), you need to:
//   1. git clone https://github.com/infobloxopen/protoc-gen-gorm.git
//   2. cd protoc-gen-gorm && go install
//   3. which protoc-gen-gorm (to ensure it's on your path)
//   4. go generate ./...  (or go generate ./gorm/...)

package fsm

import (
	"fmt"
)

type MutableData interface {
        GetStep() string
        SetStep(string)
        GetParameter(string) (string, error)
        SetParameter(string, string)
}

// GetParameter checks for the named parameter in the Parameter List and if found, returns it; if
// not, returns an error.
func (m *Mutable) GetParameter(name string) (val string, err error) {
        if v, ok := m.GetVariables()[name]; ok {
                return v, err
        }
        return val, fmt.Errorf(`Parameter "%s" not found`, name)
}

// SetParameter is parallel to GetParameter(): it stores the parameter into the Parameter List,
// replacing any pre-existing parameter of the same name/key
func (m *Mutable) SetParameter(name string, val string) {
        m.Variables[name] = val
}

// SetStep is parallel to GetStep(): it stores the value into the mu.Step
func (m *Mutable) SetStep(name string) {
        m.Step = name
}


