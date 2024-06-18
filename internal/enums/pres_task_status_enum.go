// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package enums

import (
	"fmt"
	"strings"
)

const (
	// PreservationTaskStatusUnspecified is a PreservationTaskStatus of type Unspecified.
	PreservationTaskStatusUnspecified PreservationTaskStatus = iota
	// PreservationTaskStatusInProgress is a PreservationTaskStatus of type InProgress.
	PreservationTaskStatusInProgress
	// PreservationTaskStatusDone is a PreservationTaskStatus of type Done.
	PreservationTaskStatusDone
	// PreservationTaskStatusError is a PreservationTaskStatus of type Error.
	PreservationTaskStatusError
	// PreservationTaskStatusQueued is a PreservationTaskStatus of type Queued.
	PreservationTaskStatusQueued
	// PreservationTaskStatusPending is a PreservationTaskStatus of type Pending.
	PreservationTaskStatusPending
)

var ErrInvalidPreservationTaskStatus = fmt.Errorf("not a valid PreservationTaskStatus, try [%s]", strings.Join(_PreservationTaskStatusNames, ", "))

const _PreservationTaskStatusName = "UnspecifiedInProgressDoneErrorQueuedPending"

var _PreservationTaskStatusNames = []string{
	_PreservationTaskStatusName[0:11],
	_PreservationTaskStatusName[11:21],
	_PreservationTaskStatusName[21:25],
	_PreservationTaskStatusName[25:30],
	_PreservationTaskStatusName[30:36],
	_PreservationTaskStatusName[36:43],
}

// PreservationTaskStatusNames returns a list of possible string values of PreservationTaskStatus.
func PreservationTaskStatusNames() []string {
	tmp := make([]string, len(_PreservationTaskStatusNames))
	copy(tmp, _PreservationTaskStatusNames)
	return tmp
}

var _PreservationTaskStatusMap = map[PreservationTaskStatus]string{
	PreservationTaskStatusUnspecified: _PreservationTaskStatusName[0:11],
	PreservationTaskStatusInProgress:  _PreservationTaskStatusName[11:21],
	PreservationTaskStatusDone:        _PreservationTaskStatusName[21:25],
	PreservationTaskStatusError:       _PreservationTaskStatusName[25:30],
	PreservationTaskStatusQueued:      _PreservationTaskStatusName[30:36],
	PreservationTaskStatusPending:     _PreservationTaskStatusName[36:43],
}

// String implements the Stringer interface.
func (x PreservationTaskStatus) String() string {
	if str, ok := _PreservationTaskStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PreservationTaskStatus(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PreservationTaskStatus) IsValid() bool {
	_, ok := _PreservationTaskStatusMap[x]
	return ok
}

var _PreservationTaskStatusValue = map[string]PreservationTaskStatus{
	_PreservationTaskStatusName[0:11]:  PreservationTaskStatusUnspecified,
	_PreservationTaskStatusName[11:21]: PreservationTaskStatusInProgress,
	_PreservationTaskStatusName[21:25]: PreservationTaskStatusDone,
	_PreservationTaskStatusName[25:30]: PreservationTaskStatusError,
	_PreservationTaskStatusName[30:36]: PreservationTaskStatusQueued,
	_PreservationTaskStatusName[36:43]: PreservationTaskStatusPending,
}

// ParsePreservationTaskStatus attempts to convert a string to a PreservationTaskStatus.
func ParsePreservationTaskStatus(name string) (PreservationTaskStatus, error) {
	if x, ok := _PreservationTaskStatusValue[name]; ok {
		return x, nil
	}
	return PreservationTaskStatus(0), fmt.Errorf("%s is %w", name, ErrInvalidPreservationTaskStatus)
}

// Values implements the entgo.io/ent/schema/field EnumValues interface.
func (x PreservationTaskStatus) Values() []string {
	return PreservationTaskStatusNames()
}

// PreservationTaskStatusInterfaces returns an interface list of possible values of PreservationTaskStatus.
func PreservationTaskStatusInterfaces() []interface{} {
	var tmp []interface{}
	for _, v := range _PreservationTaskStatusNames {
		tmp = append(tmp, v)
	}
	return tmp
}

// ParsePreservationTaskStatusWithDefault attempts to convert a string to a ContentType.
// It returns the default value if name is empty.
func ParsePreservationTaskStatusWithDefault(name string) (PreservationTaskStatus, error) {
	if name == "" {
		return _PreservationTaskStatusValue[_PreservationTaskStatusNames[0]], nil
	}
	if x, ok := _PreservationTaskStatusValue[name]; ok {
		return x, nil
	}
	var e PreservationTaskStatus
	return e, fmt.Errorf("%s is not a valid PreservationTaskStatus, try [%s]", name, strings.Join(_PreservationTaskStatusNames, ", "))
}

// NormalizePreservationTaskStatus attempts to parse a and normalize string as content type.
// It returns the input untouched if name fails to be parsed.
// Example:
//
//	"enUM" will be normalized (if possible) to "Enum"
func NormalizePreservationTaskStatus(name string) string {
	res, err := ParsePreservationTaskStatus(name)
	if err != nil {
		return name
	}
	return res.String()
}
