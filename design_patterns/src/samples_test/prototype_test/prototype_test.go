package prototype_test

import (
	"../../samples/prototype"
	"testing"
)

func TestClone(t *testing.T) {
	jobManager := prototype.NewJobManager()
	teacher := &prototype.Teacher{
		Name: "teacher",
	}
	jobManager.Set("teacher", teacher)
	cache := jobManager.Get("teacher")
	t.Logf("%p", cache)
	result := cache.Clone()
	t.Logf("%p", result)

}
