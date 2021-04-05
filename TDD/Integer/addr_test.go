package Integer

import "testing"
import "fmt"


func TestHello(t *testing.T) {
	t.Run("", func(t *testing.T) {
		sum := Add(2,4)
		expected := 6
		fmt.Println(sum)

		if sum != expected{
			t.Error("Laws of math are broken")
		}

		
	})
}


