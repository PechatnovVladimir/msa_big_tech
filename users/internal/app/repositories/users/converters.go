package users

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"strings"
	"time"
)

type UserProfileRow struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Nickname  string    `db:"nickname"`
	Bio       string    `db:"bio"`
	AvatarURL string    `db:"avatar_url"`
	CreatedAt time.Time `db:"created_at"`
}

func (row *UserProfileRow) Values() []any {
	return []any{
		row.ID, row.Email, row.Nickname, row.Bio, row.AvatarURL, row.CreatedAt,
	}
}

func ToModel(r *UserProfileRow) *users.UserProfile {
	if r == nil {
		return nil
	}
	return &users.UserProfile{
		ID:       r.ID,
		Email:    r.Email,
		Nickname: r.Nickname,
		Bio:      r.Bio,
		Avatar:   r.AvatarURL,
		CreateAt: r.CreatedAt,
	}
}

func FromModel(m *users.UserProfile) UserProfileRow {
	if m == nil {
		return UserProfileRow{}
	}
	return UserProfileRow{
		ID:        m.ID,
		Email:     strings.ToLower(strings.TrimSpace(m.Email)),
		Nickname:  m.Nickname,
		Bio:       m.Bio,
		AvatarURL: m.Avatar,
		CreatedAt: m.CreateAt,
	}
}
