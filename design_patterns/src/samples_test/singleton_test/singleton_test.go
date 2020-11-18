package singleton_test

import (
	"fmt"
	"sync"
	"testing"
)
import "../../samples/singleton"

func TestSingleton(T *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func() {
			instance := singleton.GetInstance()
			// 因为singleton GetInstance 返回的就是一个指针 打印指针的地址空间用下面的例子
			fmt.Printf("%p\n", instance)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestLazySingleton(T *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func() {
			lazyInstance := singleton.GetLazyInstance()
			// 因为singleton GetInstance 返回的就是一个指针 打印指针的地址空间用下面的例子
			fmt.Printf("%p\n", lazyInstance)
			wg.Done()
		}()
	}
	wg.Wait()
}
