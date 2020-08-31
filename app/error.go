package app

import "errors"

var (
	APIRequestError         = errors.New("api request error")     // APIのリクエストに失敗したエラー
	WebSocketIsStartedError = errors.New("web socket is started") // WebSocketがつながっている
	WebSocketIsStoppedError = errors.New("web socket is stopped") // WebSocketがつながっていない

	UninitializedError = errors.New("uninitialized error")  // 初期化されていない
	DataNotFoundError  = errors.New("data not found error") // データが見つからないエラー
)
