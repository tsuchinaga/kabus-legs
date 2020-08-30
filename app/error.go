package app

import "errors"

var APIRequestError = errors.New("api request error") // APIのリクエストに失敗したエラー

var DataNotFoundError = errors.New("data not found error") // データが見つからないエラー
