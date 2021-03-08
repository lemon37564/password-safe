package storage

// Data stores evey data in map
type Data struct {
	Map map[string]Pair
	key []byte
}

// NewData reutrn a Data
func NewData(key []byte) *Data {
	return &Data{Map: make(map[string]Pair), key: key}
}

// Load from safe file
func (p *Data) Load() {
	p.Map = read(p.key)
}

// Get the value from map
func (p *Data) Get(key string) Pair {
	return p.Map[key]
}

// Assign the value into map
func (p *Data) Assign(key string, pair Pair) {
	p.Map[key] = pair
	store(p.Map, p.key)
}

// Delete a value from map
func (p *Data) Delete(key string) {
	delete(p.Map, key)
	store(p.Map, p.key)
}
