package v1

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// MockUsersClient is a mock implementation of the usersClient interface
type MockUsersClient struct {
	mock.Mock
}

func (m *MockUsersClient) CreateUser(ctx context.Context, req *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.User, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.User), args.Error(1)
}

func (m *MockUsersClient) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.Empty), args.Error(1)
}

// TestPostUserHandler tests the POST user handler
func TestPostUserHandler(t *testing.T) {
	// Create a new mock users client
	mockUsersClient := new(MockUsersClient)

	// Create a new handler with the mock users client
	handler := New(mockUsersClient, nil, nil)

	// Define a sample request body
	requestBody := []byte(`{"id": "1", "username": "test_user", "password": "test_password"}`)

	// Create a new HTTP request with the sample request body
	req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP test recorder
	rr := httptest.NewRecorder()

	// Serve the HTTP request using the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check if the CreateUser method on the mock users client was called
	mockUsersClient.AssertCalled(t, "CreateUser", mock.Anything, mock.Anything)
}

// TestDeleteUserHandler tests the DELETE user handler
func TestDeleteUserHandler(t *testing.T) {
	// Create a new mock users client
	mockUsersClient := new(MockUsersClient)

	// Create a new handler with the mock users client
	handler := New(mockUsersClient, nil, nil)

	// Create a new HTTP request
	req, err := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP test recorder
	rr := httptest.NewRecorder()

	// Serve the HTTP request using the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// Check if the DeleteUser method on the mock users client was called
	mockUsersClient.AssertCalled(t, "DeleteUser", mock.Anything, mock.Anything)
}
