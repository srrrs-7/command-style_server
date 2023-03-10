package db

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateCodes(t *testing.T) {
	username := "srrrs"
	code := utils.RandomString(10)
	description := utils.RandomString(20)
	access := int64(1)

	// create code
	arg1 := CreateCodeParams{
		Username:    username,
		Code:        code,
		Img:         []byte{10},
		Description: description,
		Performance: "",
		Star:        []int64{1, 2},
		Tags:        []string{"go"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Access:      access,
	}
	err := testQueries.CreateCode(context.Background(), arg1)
	if err != nil {
		require.True(t, strings.Contains(fmt.Sprintf("%s", err), "violates foreign key constraint"))
	} else {
		require.NoError(t, err)
	}
}

func TestGetCode(t *testing.T) {
	id := int64(0)
	code, err := testQueries.GetCode(context.Background(), id)
	if err != nil {
		require.True(t, strings.Contains(fmt.Sprintf("%s", err), "no rows in result set"))
	} else {
		require.NotEmpty(t, code)
	}
}

func TestGetAllCodes(t *testing.T) {
	arg := GetAllCodesParams{
		Limit:  30,
		Offset: 0,
	}
	codes, err := testQueries.GetAllCodes(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestGetAllCodesByKeyword(t *testing.T) {
	keyword := "search"
	arg := GetAllCodesByKeywordParams{
		Username:    keyword,
		Code:        keyword,
		Description: keyword,
		Limit:       30,
		Offset:      0,
	}
	codes, err := testQueries.GetAllCodesByKeyword(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestGetCodesByTag(t *testing.T) {
	tags := make([]string, 10)
	tags[0] = "go"
	arg := GetAllCodesByTagParams{
		Column1:  tags[0],
		Column2:  tags[1],
		Column3:  tags[2],
		Column4:  tags[3],
		Column5:  tags[4],
		Column6:  tags[5],
		Column7:  tags[6],
		Column8:  tags[7],
		Column9:  tags[8],
		Column10: tags[9],
		Limit:    30,
		Offset:   0,
	}
	codes, err := testQueries.GetAllCodesByTag(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestGetAllCodesSortAccess(t *testing.T) {
	arg := GetAllCodesSortedAccessParams{
		Limit:  30,
		Offset: 0,
	}
	codes, err := testQueries.GetAllCodesSortedAccess(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestGetAllCodesSortStar(t *testing.T) {
	arg := GetAllCodesSortedStarParams{
		Limit:  30,
		Offset: 0,
	}
	codes, err := testQueries.GetAllCodesSortedStar(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestGetAllOwnCodes(t *testing.T) {
	username := utils.RandomUser(5)
	arg := GetAllOwnCodesParams{
		Username: username,
		Limit:    30,
		Offset:   0,
	}
	codes, err := testQueries.GetAllOwnCodes(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(codes) >= 0, true)
}

func TestUpdateAccess(t *testing.T) {
	arg := UpdateAccessParams{
		ID:     3,
		Access: 1,
	}
	err := testQueries.UpdateAccess(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdateStar(t *testing.T) {
	arg := UpdateStarParams{
		ID:   3,
		Star: []int64{1, 2, 3},
	}
	err := testQueries.UpdateStar(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdateCode(t *testing.T) {
	code := utils.RandomString(10)
	description := utils.RandomString(20)
	arg := UpdateCodeParams{
		ID:          3,
		Code:        code,
		Img:         []byte{10},
		Description: description,
		Performance: "",
		Tags:        []string{"go"},
		UpdatedAt:   time.Now(),
	}
	err := testQueries.UpdateCode(context.Background(), arg)
	require.NoError(t, err)
}
