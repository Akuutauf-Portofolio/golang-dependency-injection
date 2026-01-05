package test

import (
	"belajar-go-lang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mengimeplementasi pengujian dengan hasil generate kode dari google wire untuk initialized service
func TestSimpleServiceError(t *testing.T) {
	// membuat variabel untuk menymipan service baru
	simpleService, err := simple.InitializedService(true)
	
	// memebandingkan dengan assert
	assert.Nil(t, simpleService) // artinya errornya ada
	assert.NotNil(t, err) // datanya kosong, karena terjadi erorr
}

func TestSimpleServiceSuccess(t *testing.T) {
	// membuat variabel untuk menymipan service baru
	simpleService, err := simple.InitializedService(false)
	
	// memebandingkan dengan assert
	assert.Nil(t, err) // artinya errornya ada
	assert.NotNil(t, simpleService) // datanya kosong, karena terjadi erorr
}