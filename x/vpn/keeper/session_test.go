package keeper

import (
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestKeeper_SetSessionDetails(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.SetSessionDetails(ctx, &TestSessionEmpty)
	require.Nil(t, err)

	err = keeper.SetSessionDetails(ctx, &TestSessionValid)
	require.Nil(t, err)
	result1, err := keeper.GetSessionDetails(ctx, TestSessionValid.ID)
	require.Nil(t, err)
	require.Equal(t, &TestSessionValid, result1)
}

func TestKeeper_GetSessionDetails(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	result1, err := keeper.GetSessionDetails(ctx, TestSessionValid.ID)
	require.Nil(t, err)
	require.Nil(t, result1)

	err = keeper.SetSessionDetails(ctx, &TestSessionValid)
	require.Nil(t, err)
	result2, err := keeper.GetSessionDetails(ctx, TestSessionValid.ID)
	require.Nil(t, err)
	require.Equal(t, &TestSessionValid, result2)
}

func TestKeeper_SetActiveSessionIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.SetActiveSessionIDsAtHeight(ctx, 0, TestSessionIDsEmpty)
	require.Nil(t, err)
	result1, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, TestSessionIDsEmpty, result1)

	err = keeper.SetActiveSessionIDsAtHeight(ctx, 0, TestSessionIDsValid)
	require.Nil(t, err)
	result2, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, TestSessionIDsValid, result2)
}

func TestKeeper_GetActiveSessionIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	result1, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Nil(t, result1)

	err = keeper.SetActiveSessionIDsAtHeight(ctx, 0, TestSessionIDsValid)
	require.Nil(t, err)
	result2, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, TestSessionIDsValid, result2)
}

func TestKeeper_SetSessionsCount(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.SetSessionsCount(ctx, types.TestAddress, 0)
	require.Nil(t, err)
	result1, err := keeper.GetSessionsCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(0), result1)

	err = keeper.SetSessionsCount(ctx, types.TestAddress, 1)
	require.Nil(t, err)
	result2, err := keeper.GetSessionsCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(1), result2)
}

func TestKeeper_GetSessionsCount(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	result1, err := keeper.GetSessionsCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(0), result1)

	err = keeper.SetSessionsCount(ctx, types.TestAddress, 1)
	require.Nil(t, err)
	result2, err := keeper.GetSessionsCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(1), result2)
}

func TestKeeper_AddSession(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	tags, err := keeper.AddSession(ctx, &TestSessionValid)
	require.Nil(t, err)
	require.Equal(t, TestSessionTagsValid, tags)
	result1, err := keeper.GetSessionsCount(ctx, types.TestAddress)
	require.Nil(t, err)
	require.Equal(t, uint64(1), result1)
	result2, err := keeper.GetSessionDetails(ctx, TestSessionValid.ID)
	require.Nil(t, err)
	require.Equal(t, &TestSessionValid, result2)
}

func TestKeeper_AddActiveSessionIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.AddActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)
	result1, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, types.SessionIDs{types.TestSessionIDValid}, result1)
	err = keeper.AddActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)
	result2, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, types.SessionIDs{types.TestSessionIDValid}, result2)
}

func TestKeeper_RemoveActiveSessionIDsAtHeight(t *testing.T) {
	var err csdkTypes.Error
	ctx, keeper, _ := TestCreateInput()

	err = keeper.RemoveActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)

	err = keeper.AddActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)
	result1, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, types.SessionIDs{types.TestSessionIDValid}, result1)

	err = keeper.RemoveActiveSessionIDsAtHeight(ctx, 0, types.TestSessionIDValid)
	require.Nil(t, err)
	require.Equal(t, types.SessionIDs{types.TestSessionIDValid}, result1)
	result2, err := keeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Nil(t, result2)
}
