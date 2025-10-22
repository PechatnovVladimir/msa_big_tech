package users

const usersTable = "user_profiles"

const (
	usersTableColumnID        = "id"
	usersTableColumnEmail     = "email"
	usersTableColumnNickname  = "nickname"
	usersTableColumnBio       = "bio"
	usersTableColumnAvatarURL = "avatar_url"
	usersTableColumnCreatedAt = "created_at"
)

var usersTableColumns = []string{
	usersTableColumnID,
	usersTableColumnEmail,
	usersTableColumnNickname,
	usersTableColumnBio,
	usersTableColumnAvatarURL,
	usersTableColumnCreatedAt,
}
