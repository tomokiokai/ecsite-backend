package main

import (
	"go-rest-api/db"
	"go-rest-api/model"
	"log"
	"math/rand"
	"time"
)

var (
	names       = []string{"さくら", "たくみ", "やまと", "はーもにー", "さんらいず", "にほんばし", "とうきょうてーぶる", "おおさかぐりる", "きょうときっちん", "ほっかいどう"}
	addresses   = []string{"東京都千代田区100-0001", "大阪府大阪市北区530-0001", "愛知県名古屋市中区460-0001", "北海道札幌市中央区060-0000", "福岡県福岡市博多区812-0001"}
	areas       = []string{"東京", "大阪", "愛知", "北海道", "福岡"}
	genres      = []string{"レストラン", "ラーメン", "居酒屋", "イタリアン", "焼き肉", "寿司"}
	descriptions = []string{
		"居心地の良い雰囲気で本格的な日本の味を楽しめます。",
		"東京の街中にいるような本物のラーメンの味を体験してください。",
		"笑いが共有され、思い出が作られる場所です。",
		"日本風のひねりを加えたイタリア料理の豊かな味を堪能してください。",
		"自分で肉を焼いて完璧にし、楽しくインタラクティブな食事の体験をしてください。",
		"海の新鮮な食材から作られた寿司をお楽しみください。",
	}
)

func main() {
	// ランダムな値を生成するためのシードを設定
	rand.Seed(time.Now().UnixNano())

	// データベース接続の取得
	database := db.NewDB()

	// シードデータを生成
	shops := make([]model.Shop, 10)
	for i := 0; i < 10; i++ {
		shops[i] = model.Shop{
			Name:        names[rand.Intn(len(names))],
			Address:     addresses[rand.Intn(len(addresses))],
			Area:        areas[rand.Intn(len(areas))],
			Genre:       genres[rand.Intn(len(genres))],
			Description: descriptions[rand.Intn(len(descriptions))],
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	}

	// データベースにシードデータを挿入
	for _, shop := range shops {
		result := database.Create(&shop)
		if result.Error != nil {
			log.Fatalf("Failed to seed shop: %v", result.Error)
		}
	}

	log.Println("Seeder ran successfully")
}
