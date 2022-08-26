package dependency

var loaded = map[string]any{}

func Singleton[T any](key string, provider func() T) T {
	if _, ok := loaded[key]; !ok {
		loaded[key] = provider()
	}

	return loaded[key].(T)
}
