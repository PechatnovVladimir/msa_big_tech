package social

const (
	friendRequestsTable = "friend_requests"
)

const (
	friendRequestsTableColumnID         = "id"
	friendRequestsTableColumnFromUserID = "from_user_id"
	friendRequestsTableColumnToUserID   = "to_user_id"
	friendRequestsTableColumnStatus     = "status"
	friendRequestsTableColumnCreatedAt  = "created_at"
)

var friendRequestsTableColumns = []string{
	friendRequestsTableColumnID,
	friendRequestsTableColumnFromUserID,
	friendRequestsTableColumnToUserID,
	friendRequestsTableColumnStatus,
	friendRequestsTableColumnCreatedAt,
}
