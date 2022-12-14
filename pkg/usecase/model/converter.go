package model

import (
	"time"

	pref "github.com/diverse-inc/jp_prefecture"
)

func prefCodeToPrefKanji(prefCode int) string {
	location := "その他"
	prefInfo, ok := pref.FindByCode(prefCode)
	if ok {
		location = prefInfo.KanjiShort()
	}

	return location
}

// TODO: テストしやすいように引数にtime.Now()渡す
func calcAge(birthday, now time.Time) (int, error) {
	// タイムゾーンをJSTに設定
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := now.In(jst)

	thisYear, thisMonth, thisDay := nowJST.Date()
	age := thisYear - birthday.Year()

	// 誕生日を迎えていない時の処理
	if thisMonth < birthday.Month() || thisMonth == birthday.Month() && thisDay < birthday.Day() {
		age -= 1
	}

	return age, nil
}
