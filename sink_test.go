package storm

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/asdine/storm/v3/codec/json"
	"github.com/stretchr/testify/require"
)

func Test_compareValue(t *testing.T) {
	s := &sorter{
		node: &node{
			codec: json.Codec,
		},
	}
	getValue := func(l, r interface{}) (reflect.Value, reflect.Value) {
		return reflect.ValueOf(l), reflect.ValueOf(r)
	}

	require.Equal(t, 1, s.compareValue(getValue(100, 25)))
	require.Equal(t, 1, s.compareValue(getValue(25.137, 25.0)))
	require.Equal(t, -1, s.compareValue(getValue(25.137, 27.0)))
	require.Equal(t, 1, s.compareValue(getValue(big.NewInt(100), big.NewInt(100))))
	require.Equal(t, -1, s.compareValue(getValue(big.NewInt(10), big.NewInt(100))))
	require.Equal(t, 1, s.compareValue(getValue(big.NewFloat(10.11), big.NewFloat(10.11))))
	require.Equal(t, -1, s.compareValue(getValue(big.NewFloat(10.00), big.NewFloat(10.12))))
}
