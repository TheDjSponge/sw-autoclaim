package users

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type mockValidator struct{
	shouldBeValid bool
	returnedUID int
}
func (m *mockValidator) CheckUser(ctx context.Context, id, srv string) (bool, int, error){
	return m.shouldBeValid, m.returnedUID, nil
}

type mockDB struct{
	addDiscordUserErr error

	addUserErr error

	deleteUserErr error
	deleteUserRows int

	users []database.User
	getAllUsersErr error
}

func (m *mockDB) AddDiscordUser(ctx context.Context, params database.AddDiscordUserParams) error{
	return m.addDiscordUserErr
}
func (m *mockDB) AddUser(ctx context.Context, params database.AddUserParams) error{
	return m.addUserErr
}
func (m *mockDB) DeleteUserAndCount(ctx context.Context, params database.DeleteUserParams) (int,error){
	return m.deleteUserRows, m.deleteUserErr
}
func (m *mockDB) GetAllUsers(ctx context.Context) ([]database.User, error){
	return m.users, m.getAllUsersErr
}

func TestRegisterUser_Unit(t *testing.T){
	mockV := &mockValidator{shouldBeValid: true, returnedUID: 999}
	mockDB := &mockDB{}
	svc := NewService(mockDB, mockV)

	err := svc.RegisterUser(context.Background(), "user1", "europe", "bob", 123)

	if err != nil {
		t.Errorf("found error when none is expected when registering user")
	}
}

func TestDeleteUser_Unit(t *testing.T) {
    tests := []struct {
        name          string
        rowsAffected  int
        dbErr         error
        expectedError string
    }{
        {"Success",		 	1, nil, 					""},
        {"User Not Found", 	0, nil, 					"failed to reference any user to delete"},
        {"Database Crash", 	0, fmt.Errorf("db down"), 	"failed to delete user"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockDB := &mockDB{
                // We update the mock to return exactly what we want for this case
                deleteUserRows: tt.rowsAffected, 
                deleteUserErr:  tt.dbErr,
            }
            svc := NewService(mockDB, &mockValidator{})
            
            err := svc.DeleteUser(context.Background(), 123, "id", "srv")
            
            if tt.expectedError == "" && err != nil {
                t.Errorf("expected success, got %v", err)
            }
            if tt.expectedError != "" && (err == nil || !strings.Contains(err.Error(), tt.expectedError)) {
                t.Errorf("expected error containing %q, got %v", tt.expectedError, err)
            }
        })
    }
}
func makeUUID(s string) pgtype.UUID{
	var u pgtype.UUID
	u.UnmarshalJSON([]byte(fmt.Sprintf("%q",s)))
	return u
}

func makePGBool(b string) pgtype.Bool{
	var u pgtype.Bool
	u.UnmarshalJSON([]byte(fmt.Sprintf("%q",b)))
	return u
}

func TestGetAllUsers_Unit(t *testing.T){
	tests := []struct{
		name string
		mockReturn []database.User
		mockErr error
		expectedCount int
		expectErr bool
	}{
		{
			name: "Success with users",
			mockReturn: []database.User{
				{ID: makeUUID("550e8400-e29b-41d4-a716-446655440000"),DiscordID: 12,HiveID: "user_a", Server: "europe", Active: makePGBool("true"), GameUid: 012},
				{ID: makeUUID("234e8400-e29b-41d4-a716-446655440000"),DiscordID: 54,HiveID: "user_b", Server: "europe", Active: makePGBool("false"), GameUid: 344},
			},
			expectedCount: 2,
			expectErr: false,
		},
		{
			name: "Success but empty",
			mockReturn: []database.User{},
			expectedCount: 0,
			expectErr: false,
		},
		{
			name: "Database Error",
			mockErr: fmt.Errorf("connection lost"),
			expectErr: true,
		},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			db := &mockDB{users: tt.mockReturn, getAllUsersErr: tt.mockErr}
			svc := NewService(db, nil)

			users, err := svc.GetAllUsers(context.Background())

			if (err != nil) != tt.expectErr {
				t.Errorf("GetAllUsers() error = %v, expectErr %v", err, tt.expectErr)
			}
			if len(users) != tt.expectedCount {
				t.Errorf("Expected %d users, got %d", tt.expectedCount, len(users))
			}
		})
	}
}

