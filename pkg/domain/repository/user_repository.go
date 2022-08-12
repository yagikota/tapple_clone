package domain

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
)

type IUserRepository interface {
	FindByUserID(ctx context.Context, userID int) (*entity.User, error)                                  // 1ユーザー取得
	FindAll(ctx context.Context) (entity.UserSlice, error)                                               // 全ユーザー取得
	SendMessage(ctx context.Context, userID int, roomID int, m *entity.Message) (*entity.Message, error) // メッセージ送信
}
