package GuitarTuningUtil

import (
	"../OBJ"
)

const (
	STANDARD = iota
	SEMITONE_LOWER
	DROP_D
	DROP_C
	DROP_B
	OPEN_G
	OPEN_A
	OPEN_E
	OPEN_D
	OPEN_GSUS4
	OPEN_C6
	DADGAD
	DADEAD
)

var GuitarTuningNameMap = map[int]string{
	STANDARD:       "レギュラー",
	SEMITONE_LOWER: "半音下げ",
	DROP_D:         "ドロップD",
	DROP_C:         "ドロップC",
	DROP_B:         "ドロップB",
	OPEN_G:         "オープンG",
	OPEN_A:         "オープンA",
	OPEN_E:         "オープンE",
	OPEN_D:         "オープンD",
	OPEN_GSUS4:     "オープンGsus4",
	OPEN_C6:        "オープンC6",
	DADGAD:         "DADGAD",
	DADEAD:         "DAEAD",
}

var GuitarTuningToneMap = map[int]string{
	STANDARD:       "E A D G B E",
	SEMITONE_LOWER: "D# G# C# F# A# D#",
	DROP_D:         "D A D G B E",
	DROP_C:         "C G C F A D",
	DROP_B:         "B F# B E G# C#",
	OPEN_G:         "D G D G B D",
	OPEN_A:         "E A E A C# E",
	OPEN_E:         "E B E G# B E",
	OPEN_D:         "D A D F# A D",
	OPEN_GSUS4:     "D G C G C D",
	OPEN_C6:        "C A C G C E",
	DADGAD:         "D A D G A D",
	DADEAD:         "D A E A D",
}

//CONST + チューニングオブジェクトをマップとして返す関数
//引数：なし
//戻り値：チューニングオブジェクトマップ
func MakeTuningMap() map[int]OBJ.GuitarTuning {
	var GuitarTuningOBJMap map[int]OBJ.GuitarTuning = map[int]OBJ.GuitarTuning{}
	for k, v := range GuitarTuningNameMap {
		var gt OBJ.GuitarTuning
		gt.TuningName = v
		gt.TuningTone = GuitarTuningToneMap[k]
		GuitarTuningOBJMap[k] = gt
	}
	return GuitarTuningOBJMap

	//↓memo
	//	var gt OBJ.GuitarTuning
	//	gt.tuningName = "aaa"
	//	gt.tuningTone = "aaa"
	//	GuitarTuningOBJMap := map[int]OBJ.GuitarTuning{
	//		STANDARD:       gt,
	//		SEMITONE_LOWER: OBJ.NewGuitarTuning("半音下げ", "GO"),
	//	}
}
