package repository

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"pjdrc/database"
	"pjdrc/domain/user"
	"strconv"
)

type userAreaRepository struct {
	spannerClient *spanner.Client
	tableName     string
}

func NewUserAreaRepository(spannerClient *spanner.Client, tableName string) user.UserAreaRepository {
	return &userAreaRepository{
		spannerClient: spannerClient,
		tableName:     tableName,
	}
}

func (repo *userAreaRepository) GetUserArea(c context.Context, tx *spanner.ReadWriteTransaction, userID string) (user.UserArea, error) {
	var userArea user.UserArea
	columnNames, columns, err := database.GetSpannerColumns(user.UserArea{})
	if err != nil {
		return userArea, err
	}

	userIDInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return userArea, fmt.Errorf("invalid userID: %w", err)
	}

	stmt := spanner.Statement{
		SQL:    fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1", columnNames, repo.tableName, columns[user.UserID]),
		Params: map[string]interface{}{"p1": userIDInt},
	}

	iter := tx.Query(c, stmt)
	defer iter.Stop()

	if err := iter.Do(func(row *spanner.Row) error {
		return row.ToStruct(&userArea)
	}); err != nil {
		return userArea, err
	}
	return userArea, nil
}

func (repo *userAreaRepository) Create(ctx context.Context, tx *spanner.ReadWriteTransaction, userID int64) error {

	_, columns, err := database.GetSpannerColumns(user.UserArea{})
	if err != nil {
		return err
	}

	// 挿入用のミューテーションを作成
	mutation := spanner.Insert(repo.tableName, []string{columns[user.UserID], columns[user.AreaID], columns[user.IsStay]},
		[]interface{}{userID, 1, true})

	// トランザクション内で挿入を行う
	if err := tx.BufferWrite([]*spanner.Mutation{mutation}); err != nil {
		return err
	}

	return nil
}
