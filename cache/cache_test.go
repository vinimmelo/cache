package cache

import "testing"

func Test_Find(t *testing.T) {
	cache := New()
	cache.Add("test", 25)

	t.Run("when we find the key", func(t *testing.T) {
		result, err := cache.Find("test")

		if err != nil {
			t.Error(err)
		}

		if result.(int) != 25 {
			t.Errorf("The result %d should be 25", result.(int))
		}
	})

	t.Run("when the key wasn't found", func(t *testing.T) {
		result, err := cache.Find("inexistent")
		if err == nil {
			t.Error("Should've returned an error")
		}

		if result != nil {
			t.Error("Should've returned nil")
		}
	})
}

type StructTest struct {
	Value int
}

func Test_Add(t *testing.T) {
	cache := New()

	t.Run("when we add successfully", func(t *testing.T) {
		t.Run("when we add an int", func(t *testing.T) {
			cache.Add("test", 25)

			result, ok := cache.cache["test"]

			if !ok {
				t.Error("The key test wasn't found")
			}

			if result.(int) != 25 {
				t.Errorf("The result %d should be 25", result.(int))
			}
		})

		t.Run("when we add a string", func(t *testing.T) {
			value := "test"
			cache.Add("test-string", value)

			result, ok := cache.cache["test-string"]
			if !ok {
				t.Error("The key test wasn't found")
			}

			if result.(string) != "test" {
				t.Errorf("The result %s should be 25", result.(string))
			}
		})

		t.Run("when we add a struct", func(t *testing.T) {
			value := StructTest{25}
			cache.Add("test-struct", value)

			result, ok := cache.cache["test-struct"]
			if !ok {
				t.Error("The key test wasn't found")
			}

			if result.(StructTest).Value != 25 {
				t.Errorf("The result %d should be 25", result.(StructTest).Value)
			}
		})

	})
}

func Test_Del(t *testing.T) {
	cache := New()

	t.Run("when the delete succeed", func(t *testing.T) {
		cache.Add("test", 25)

		err := cache.Del("test")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("when the delete doesn't found a key", func(t *testing.T) {
		err := cache.Del("inexistent")

		if err == nil {
			t.Error("Should've returned an error")
		}
	})
}
