package main

import (
	"fmt"
	"math/rand"
	"time"
)

//他のパッケージからはアクセス不可能な構造体
//※構造体名の先頭を大文字でPublicな構造体になる。
type rapper struct {
	name      string
	punchLine string
}

func (r rapper) ShootPunchLine() {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(5)
	if randInt == 0 {
		fmt.Println("もうダメだ、こいつ強い( ﾉД`)ｼｸｼｸ…")
	} else {
		fmt.Println(r.punchLine)
	}
}

func main() {
	//rapper構造体のインスタンス化
	var ryofu rapper
	ryofu.name = "呂布カルマ"
	ryofu.punchLine = "喋ってから終わんな"

	var suna rapper
	suna.name = "スナフキン"
	suna.punchLine = "火つけるよ"

	fmt.Println("alexa! proto")
	fmt.Println("TK	DA 黒縁")
	//fmt.Println("喋ってから終わんな")
	for i := 0; i < 3; i++ {
		suna.ShootPunchLine()
	}
	ryofu.ShootPunchLine()
}
