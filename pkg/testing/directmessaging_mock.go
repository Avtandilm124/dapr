// Code generated by mockery v1.0.0.

package testing

import (
	"regexp"

	messaging "github.com/actionscore/actions/pkg/messaging"
	mock "github.com/stretchr/testify/mock"
)

// MockDirectMessaging is an autogenerated mock type for the MockDirectMessaging type
type MockDirectMessaging struct {
	mock.Mock
}

// Invoke provides a mock function with given fields: req
func (_m *MockDirectMessaging) Invoke(req *messaging.DirectMessageRequest) (*messaging.DirectMessageResponse, error) {
	re := regexp.MustCompile(`&__header_delim__&X-Correlation-Id&__header_equals__&[0-9a-fA-F]+;[0-9a-fA-F]+;[0-9a-fA-F]+`)
	req.Metadata["headers"] = re.ReplaceAllString(req.Metadata["headers"], "")
	ret := _m.Called(req)

	var r0 *messaging.DirectMessageResponse
	if rf, ok := ret.Get(0).(func(*messaging.DirectMessageRequest) *messaging.DirectMessageResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*messaging.DirectMessageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*messaging.DirectMessageRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
