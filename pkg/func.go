package constfun

import (
	"time"

	pref "github.com/diverse-inc/jp_prefecture"
)

func PrefCodeToPrefKanji(prefCode int) string {
	location := "その他"
	prefInfo, ok := pref.FindByCode(prefCode)
	if ok {
		location = prefInfo.KanjiShort()
	}

	return location
}

func CalcAge(birthday time.Time) (int, error) {
	// タイムゾーンをJSTに設定
	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	nowJST := nowUTC.In(jst)

	thisYear, thisMonth, thisDay := nowJST.Date()
	age := thisYear - birthday.Year()

	// 誕生日を迎えていない時の処理
	if thisMonth < birthday.Month() && thisDay < birthday.Day() {
		age -= 1
	}

	return age, nil
}