package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRandString
func TestRandString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want int
	}{
		{
			name: "TestRandString",
			want: 25,
		},
	}
	//
	for _, c := range tests {
		t.Run("", func(t *testing.T) {
			output := RandString(c.want)

			if !assert.Len(t, output, c.want) {
				assert.Fail(t, "RandString(%d) failed", c.want)
			}
		})
	}
}

// BenchmarkRandString
func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(25)
	}
}

//

// TestRandInt
func TestRandInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		min  int
		max  int
	}{
		{
			name: "TestRandInt",
			min:  12,
			max:  34,
		},
	}
	//
	for _, c := range tests {
		t.Run("", func(t *testing.T) {
			output := RandInt(c.min, c.max)

			// c.min<=output<=c.max
			assert.True(t, (c.min <= output && output <= c.max))
		})
	}
}

// BenchmarkRandInt
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt(0, 99999)
	}
}
