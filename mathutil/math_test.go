package mathutil

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Positif", func(t *testing.T) {
		if result := Add(3, 5); result != 8 {
			t.Errorf("expedted 8 but got %d", result)
		}
	})
	t.Run("Negatif", func(t *testing.T) {
		if result := Add(-3, -5); result != -8 {
			t.Errorf("expedted -8 but got %d", result)
		}
	})
}
