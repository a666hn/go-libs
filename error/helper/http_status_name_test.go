package error_helper

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetHttpStatusName_Success(t *testing.T) {
	s := http.StatusOK
	hsn := GetHttpStatusName(s)
	assert.Equal(t, http.StatusText(http.StatusOK), hsn, "http status name should be equal")
}

func TestGetHttpStatusName_ReturnDefault(t *testing.T) {
	s := 666
	hsn := GetHttpStatusName(s)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), hsn, "http status name should be return default")
}

func BenchmarkGetHttpStatusName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetHttpStatusName(http.StatusUnprocessableEntity)
	}
}
