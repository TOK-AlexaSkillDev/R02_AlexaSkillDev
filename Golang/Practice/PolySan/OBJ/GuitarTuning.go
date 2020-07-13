package OBJ

type GuitarTuning struct {
	//チューニング名称
	TuningName string
	//チューニングトーン
	TuningTone string
}

//コンストラクタ
//引数：オブジェクトに設定する名称とトーン
//戻り値：構造体のポインタ：
func NewGuitarTuning(name string, tone string) *GuitarTuning {
	gt := new(GuitarTuning)
	gt.TuningName = name
	gt.TuningTone = tone
	return gt
}
