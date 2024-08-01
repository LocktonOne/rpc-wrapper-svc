/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "strconv"

type Key struct {
	ID           string       `json:"id"`
	ResourceType ResourceType `json:"type"`
}

func NewKeyInt64(id int64, resourceType ResourceType) Key {
	return Key{
		ID:           strconv.FormatInt(id, 10),
		ResourceType: resourceType,
	}
}

func (r *Key) GetKey() Key {
	return *r
}

func (r Key) GetKeyP() *Key {
	return &r
}

func (r Key) AsRelation() *Relation {
	return &Relation{
		Data: r.GetKeyP(),
	}
}
