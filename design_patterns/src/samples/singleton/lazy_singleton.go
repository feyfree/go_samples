package singleton

type LazySingleton struct {
}

var lazyInstance = new(LazySingleton)

func GetLazyInstance() *LazySingleton {
	return lazyInstance
}
