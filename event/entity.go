package event
// 实体ID
type Entity interface {
	// EntityID returns the ID of the entity.
	EntityID() UUID
}
