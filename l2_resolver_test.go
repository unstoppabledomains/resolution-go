package resolution

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockedFunction struct {
	called       bool
	impl         func() (interface{}, error)
	returns      interface{}
	returnsError error
}

func (f *mockedFunction) mock() (interface{}, error) {
	f.called = true
	if f.impl != nil {
		return f.impl()
	}
	return f.returns, f.returnsError
}

func makeMock(returns interface{}, returnsError error) mockedFunction {
	return mockedFunction{called: false, returns: returns, returnsError: returnsError}
}

func TestCallsBothMethods(t *testing.T) {
	t.Parallel()
	l1Test := makeMock(nil, &DomainNotRegisteredError{DomainName: "test"})
	l2Test := makeMock(nil, &DomainNotRegisteredError{DomainName: "test"})

	_, _ = resolveGeneric(genericFunctions{l1Test.mock, l2Test.mock})

	assert.True(t, l1Test.called)
	assert.True(t, l2Test.called)
}

func TestReturnsResultFromL2(t *testing.T) {
	t.Parallel()
	l1Test := makeMock("L1 result", nil)
	l2Test := makeMock("L2 result", nil)

	result, _ := resolveGeneric(genericFunctions{l1Test.mock, l2Test.mock})
	stringResult, ok := result.(string)

	assert.True(t, l2Test.called)
	assert.True(t, ok)
	assert.Equal(t, stringResult, "L2 result")
}

func TestThrowsNetErrorsFromL2(t *testing.T) {
	t.Parallel()
	var expectedError error = errors.New("unexpected network error")
	l1Test := makeMock(nil, nil)
	l2Test := makeMock(nil, errors.New("unexpected network error"))

	_, err := resolveGeneric(genericFunctions{l1Test.mock, l2Test.mock})

	assert.True(t, l2Test.called)
	assert.Equal(t, err.Error(), expectedError.Error())
}

func TestReturnsResultFromL1(t *testing.T) {
	t.Parallel()
	l1Test := makeMock("L1 result", nil)
	l2Test := makeMock(nil, &DomainNotRegisteredError{DomainName: "test"})

	result, _ := resolveGeneric(genericFunctions{l1Test.mock, l2Test.mock})
	stringResult, ok := result.(string)

	assert.True(t, l2Test.called)
	assert.True(t, l1Test.called)
	assert.True(t, ok)
	assert.Equal(t, stringResult, "L1 result")
}

func TestThrowsNetErrorsFromL1(t *testing.T) {
	t.Parallel()
	var expectedError error = errors.New("unexpected network error")
	l1Test := makeMock(nil, errors.New("unexpected network error"))
	l2Test := makeMock(nil, &DomainNotRegisteredError{DomainName: "test"})

	_, err := resolveGeneric(genericFunctions{l1Test.mock, l2Test.mock})

	assert.True(t, l2Test.called)
	assert.Equal(t, err.Error(), expectedError.Error())
}
