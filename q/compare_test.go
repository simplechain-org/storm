package q

import (
	"go/token"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Compare(t *testing.T) {
	require.False(t, compare(big.NewInt(1011), big.NewInt(110), token.LEQ))
	require.True(t, compare(big.NewInt(1011), big.NewInt(110), token.GTR))
	require.True(t, compare(big.NewInt(11), big.NewInt(110), token.LEQ))
	require.False(t, compare(big.NewInt(11), big.NewInt(110), token.GTR))
	require.True(t, compare(big.NewInt(110), big.NewInt(110), token.EQL))

	require.True(t, compare(big.NewFloat(12.33), big.NewFloat(12.33), token.EQL))
	require.True(t, compare(big.NewFloat(12.33), big.NewFloat(12.33), token.GEQ))
	require.True(t, compare(big.NewFloat(12.33), big.NewFloat(12.33), token.LEQ))
	require.True(t, compare(big.NewFloat(12.43), big.NewFloat(12.33), token.GTR))
	require.False(t, compare(big.NewFloat(12.13), big.NewFloat(12.33), token.GTR))
	require.True(t, compare(big.NewFloat(12.13), big.NewFloat(12.33), token.LEQ))
	require.False(t, compare(big.NewFloat(12.53), big.NewFloat(12.33), token.LEQ))
}
