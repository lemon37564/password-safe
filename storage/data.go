package storage

import (
	"pass-safe/crypto"
)

// Data stores evey data in map
type Data struct {
	dict map[string]Pair
	key  []byte
	iv   []byte
}

// NewData reutrn a Data
func NewData(key []byte) *Data {
	return &Data{dict: make(map[string]Pair), key: key, iv: crypto.GenerateIV()}
}

// Store into safe file
func (d *Data) Store() {
	store(d.dict, d.key, d.iv)
}

// Load from safe file
func (d *Data) Load() error {
	var err error
	d.dict, err = read(d.key)
	if err != nil {
		return err
	}
	return nil
}

// Get the value from map
func (d *Data) Get(name string) (Pair, bool) {
	val, exist := d.dict[name]
	return val, exist
}

// GetMap return map of the data
func (d *Data) GetMap() map[string]Pair {
	return d.dict
}

// Assign the value into map
func (d *Data) Assign(name string, pair Pair) {
	d.dict[name] = pair
	d.Store()
}

// Delete a value from map
func (d *Data) Delete(key string) {
	delete(d.dict, key)
	d.Store()
}
