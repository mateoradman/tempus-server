package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/mateoradman/tempus/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeam(t *testing.T) Team {
	arg := CreateTeamParams{
		Name: util.RandomString(100),
		ManagerID: sql.NullInt64{
			Valid: true,
			Int64: createRandomUser(t).ID,
		},
	}

	team, err := testQueries.CreateTeam(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, team)
	require.Equal(t, arg.Name, team.Name)
	require.Equal(t, arg.ManagerID, team.ManagerID)
	require.NotZero(t, team.ID)
	require.WithinDuration(t, time.Now(), team.CreatedAt, 2*time.Second)
	require.False(t, team.UpdatedAt.Valid)

	return team
}

func TestCreateTeam(t *testing.T) {
	createRandomTeam(t)
}

func TestGetTeam(t *testing.T) {
	team := createRandomTeam(t)
	gotTeam, err := testQueries.GetTeam(context.Background(), team.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotTeam)
	require.Equal(t, team.ID, gotTeam.ID)
	require.Equal(t, team.Name, gotTeam.Name)
	require.Equal(t, team.ManagerID, gotTeam.ManagerID)
	require.Equal(t, team.CreatedAt, gotTeam.CreatedAt)
	require.Equal(t, team.UpdatedAt, gotTeam.UpdatedAt)
}

func TestUpdateTeam(t *testing.T) {
	team := createRandomTeam(t)
	manager := createRandomUser(t)
	expectedLen := 25
	name := util.RandomString(expectedLen)
	arg := UpdateTeamParams{
		ID:   team.ID,
		Name: name,
		ManagerID: sql.NullInt64{
			Int64: manager.ID,
			Valid: true,
		},
	}

	updatedTeam, err := testQueries.UpdateTeam(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTeam)
	require.Equal(t, team.ID, updatedTeam.ID)
	require.Equal(t, name, updatedTeam.Name)
	require.Len(t, name, expectedLen)
	require.Equal(t, arg.ManagerID, updatedTeam.ManagerID)
	require.True(t, updatedTeam.UpdatedAt.Valid)
	require.WithinDuration(t, time.Now(), updatedTeam.UpdatedAt.Time, time.Second)
	require.Equal(t, team.CreatedAt, updatedTeam.CreatedAt)
}

func TestDeleteTeam(t *testing.T) {
	team := createRandomTeam(t)
	deletedTeam, err := testQueries.DeleteTeam(context.Background(), team.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedTeam)
	require.Equal(t, team.ID, deletedTeam.ID)
	require.Equal(t, team.Name, deletedTeam.Name)
	require.False(t, deletedTeam.UpdatedAt.Valid)
	require.Equal(t, team.CreatedAt, deletedTeam.CreatedAt)
	require.Equal(t, team.ManagerID, deletedTeam.ManagerID)
}

func TestListTeams(t *testing.T) {
	for i := 0; i < 10; i++ {
		// Create 10 teams
		createRandomTeam(t)
	}

	arg := ListTeamsParams{
		Limit:  5,
		Offset: 0,
	}

	teams, err := testQueries.ListTeams(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, teams, int(arg.Limit))

	for _, team := range teams {
		require.NotEmpty(t, team)
	}
}

func TestListTeamMembers(t *testing.T) {
	users := []User{}
	team := createRandomTeam(t)
	for i := 0; i < 10; i++ {
		user := createRandomUser(t)
		arg := UpdateUserParams{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Name:     user.Name,
			Surname:  user.Surname,
			TeamID: sql.NullInt64{
				Valid: true,
				Int64: team.ID,
			},
			Gender:    user.Gender,
			BirthDate: user.BirthDate,
			Language:  user.Language,
			Country:   user.Country,
			Timezone:  user.Timezone,
			ManagerID: user.ManagerID,
			CompanyID: user.CompanyID,
		}
		updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
		require.NoError(t, err)
		users = append(users, updatedUser)
	}
	arg := ListTeamMembersParams{
		TeamID: sql.NullInt64{
			Int64: team.ID,
			Valid: true,
		},
		Limit:  100,
		Offset: 0,
	}
	employees, err := testQueries.ListTeamMembers(context.Background(), arg)
	require.NoError(t, err)
	require.Subset(t, employees, users)
}
