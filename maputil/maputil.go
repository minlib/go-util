package maputil

// ContainsKey returns the keys of the map m contains a key k.
func ContainsKey[M ~map[K]V, K comparable, V any](m M, k K) bool {
	_, found := m[k]
	return found
}
