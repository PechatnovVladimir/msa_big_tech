package consumer

type Deduplicator interface {
	Seen(id string) bool
	MarkSeen(id string)
}
