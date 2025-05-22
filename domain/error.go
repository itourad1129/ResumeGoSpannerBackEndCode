package domain

import "errors"

// ErrUserNameConflict ユーザー名の重複エラー
var ErrUserNameConflict = errors.New("user already exists")
var ErrConvertClaims = errors.New("failed to convert claims")
