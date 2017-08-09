package game

// Scene ゲームシーンのインターフェイス
type Scene interface {
	Initialize() error
	Progress()
	Draw()
}
