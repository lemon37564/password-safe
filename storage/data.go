package storage

import "pass-safe/crypto"

// Data stores evey data in map
type Data struct {
	Map map[string]Pair
	key []byte
	iv  []byte
}

// NewData reutrn a Data
func NewData(key []byte) *Data {
	return &Data{Map: make(map[string]Pair), key: key, iv: crypto.GenerateIV()}
}

// Store into safe file
func (d *Data) Store() {
	store(d.Map, d.key, d.iv)
}

// Load from safe file
func (d *Data) Load() error {
	var err error
	d.Map, err = read(d.key)
	if err != nil {
		return err
	}
	return nil
}

// Get the value from map
func (d *Data) Get(key string) Pair {
	return d.Map[key]
}

// Assign the value into map
func (d *Data) Assign(key string, pair Pair) {
	d.Map[key] = pair
	d.Store()
}

// Delete a value from map
func (d *Data) Delete(key string) {
	delete(d.Map, key)
	d.Store()
}
