package social

const (
	friendsTable = "friends"
)

const (
	friendsTableColumnUserID       = "user_id"
	friendsTableColumnFriendUserID = "friend_user_id"
	friendsTableColumnCreatedAt    = "created_at"
)

var friendsTableColumns = []string{
	friendsTableColumnUserID,
	friendsTableColumnFriendUserID,
	friendsTableColumnCreatedAt,
}
