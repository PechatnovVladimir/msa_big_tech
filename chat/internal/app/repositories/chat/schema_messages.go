package chat

const (
	messagesTable = "messages"
)

const (
	messagesTableColumnID        = "id"
	messagesTableColumnChatID    = "chat_id"
	messagesTableColumnSenderID  = "sender_id"
	messagesTableColumnText      = "text"
	messagesTableColumnCreatedAt = "created_at"
)

var messagesTableColumns = []string{
	messagesTableColumnID,
	messagesTableColumnChatID,
	messagesTableColumnSenderID,
	messagesTableColumnText,
	messagesTableColumnCreatedAt,
}
