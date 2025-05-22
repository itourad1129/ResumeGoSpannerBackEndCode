package repository

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"pjdrc/database"
	"pjdrc/domain/master"
)

type areaRepository struct {
	spannerClient *spanner.Client
	tableName     string
}

func NewAreaRepository(spannerClient *spanner.Client, tableName string) master.AreaRepository {
	return areaRepository{
		spannerClient: spannerClient,
		tableName:     tableName,
	}
}

func (repo areaRepository) GetArea(c context.Context, areaID int64) (master.Area, error) {
	var area master.Area
	columnNames, columns, err := database.GetSpannerColumns(master.Area{})
	if err != nil {
		return area, err
	}

	selectStmt := spanner.Statement{
		SQL:    fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1", columnNames, repo.tableName, columns[master.AreaID]),
		Params: map[string]interface{}{"p1": areaID},
	}

	iter := repo.spannerClient.Single().Query(c, selectStmt)
	defer iter.Stop()

	if err := iter.Do(func(row *spanner.Row) error {
		return row.ToStruct(&area)
	}); err != nil {
		return area, err
	}
	return area, nil
}
