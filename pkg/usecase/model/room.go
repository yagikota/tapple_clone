package model

import (
	"sort"
	"strconv"
	"time"

	constant "github.com/CyberAgentHack/2208-ace-go-server/pkg"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
)

type RoomID int

type RoomSlice []*Room

// ルーム一覧で使用
type Room struct {
	ID            RoomID   `json:"id"`
	Unread        int      `json:"unread"`
	IsPinned      bool     `json:"is_pinned"`
	Name          string   `json:"name"`
	SubName       string   `json:"sub_name"`
	IsPrincipal   bool     `json:"is_principal"`
	Icon          string   `json:"icon"`
	LatestMessage *Message `json:"latest_message"`
}

// ルーム一覧で使用
type Rooms struct {
	Rooms RoomSlice `json:"rooms"`
}

// ルーム詳細で使用
type RoomDetail struct {
	ID       RoomID       `json:"id"`
	Name     string       `json:"name"`
	Icon     string       `json:"icon"`
	Messages MessageSlice `json:"messages"`
	IsLast   bool         `json:"is_last"`
}

// ルーム一覧で使用
func RoomFromDomainModel(m *model.Room) *Room {
	u := UserFromDomainModel(m.R.RoomUsers[0].R.User)
	r := &Room{
		ID:            RoomID(m.ID),
		IsPinned:      m.R.RoomUsers[0].IsPinned,
		Name:          u.Name,
		Icon:          u.Icon,
		LatestMessage: MessageFromDomainModel(m.R.Messages[0]),
	}

	unread := 0
	partnerID := UserFromDomainModel(m.R.RoomUsers[0].R.User).ID
	for _, message := range m.R.Messages {
		if message.UserID == int(partnerID) && message.IsRead {
			break
		}
		if message.UserID == int(partnerID) && !message.IsRead {
			unread += 1
		}
	}
	r.Unread = unread

	age, err := calcAge(u.BirthDay, time.Now())
	if err != nil {
		return nil
	}

	// 都道府県コードを県名に変換
	location := u.Location
	r.SubName = strconv.Itoa(age) + "歳・" + location

	r.IsPrincipal = u.IsPrincipal

	return r
}

// ルーム詳細で使用
func RoomDetailFromDomainModel(m *model.Room) *RoomDetail {
	rm := &RoomDetail{
		ID:   RoomID(m.ID),
		Name: m.R.RoomUsers[0].R.User.Name,
		Icon: m.R.RoomUsers[0].R.User.Icon,
	}

	mSlice := make(MessageSlice, 0, len(m.R.Messages))
	for _, message := range m.R.Messages {
		mSlice = append(mSlice, MessageFromDomainModel(message))
	}

	// 古い順番にソート
	sort.Slice(mSlice, func(i, j int) bool {
		return mSlice[i].CreatedAt.Before(mSlice[j].CreatedAt)
	})
	rm.Messages = mSlice

	// 取得したメッセージの数がLIMIT_RECORDより少なくなったらフラグをtureに変更
	if len(mSlice) < constant.LimitMessageRecord {
		rm.IsLast = true
	}

	return rm
}
