package test

import (
	"belajar-go-lang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

// membuat pengujian untuk file dan connection (dependency injection)
func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")

	// melakukan perbandingan dengan assert
	assert.NotNil(t, connection)
	cleanup()
}