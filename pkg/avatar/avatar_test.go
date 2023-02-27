package avatar

import (
	"fmt"
	"testing"
)

func TestUrl(t *testing.T) {
	fmt.Println(Url())
	fmt.Println(Url())
	fmt.Println(Url())
	fmt.Println(Url())
	fmt.Println(Url())
}

func BenchmarkUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Url()
	}
}
