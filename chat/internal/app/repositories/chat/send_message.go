package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
)

func (r *Repository) SendMessage(ctx context.Context, m *chat.Message) (*chat.Message, error) {
	messageRow, chatMembersRow := fromModelForSendMessage(m)

	//insert into messages (id, ...) values ($1,....)
	queryMessage := r.sb.
		Insert(messagesTable).
		Columns(messagesTableColumns...).
		Values(messageRow.Values()...)

	//insert into chat_members ... Отправитель сообщения становится участником чата. Если уникальность будет нарушена, то пропустить
	queryChatMembers := r.sb.
		Insert(chatMembersTable).
		Columns(chatMembersTableColumns...).
		Values(chatMembersRow.Values()...).
		Suffix("ON CONFLICT (chat_id,user_id) DO NOTHING")

	pool := r.db.GetQueryEngine(ctx)

	_, err := pool.Execx(ctx, queryChatMembers)
	if err != nil {
		return nil, err
	}

	var outRow MessageRow
	err = pool.Getx(ctx, &outRow, queryMessage)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
