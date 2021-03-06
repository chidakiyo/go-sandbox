// Code generated by yo. DO NOT EDIT.
// Package models contains the types.
package models

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
)

// Like represents a row from 'Likes'.
type Like struct {
	ID         string    `spanner:"ID" json:"ID"`                 // ID
	Count      int64     `spanner:"Count" json:"Count"`           // Count
	CreatedAt  time.Time `spanner:"CreatedAt" json:"CreatedAt"`   // CreatedAt
	ModifiedAt time.Time `spanner:"ModifiedAt" json:"ModifiedAt"` // ModifiedAt
	TweetID    string    `spanner:"TweetID" json:"TweetID"`       // TweetID
	UserID     string    `spanner:"UserID" json:"UserID"`         // UserID
}

func LikePrimaryKeys() []string {
	return []string{
		"ID",
	}
}

func LikeColumns() []string {
	return []string{
		"ID",
		"Count",
		"CreatedAt",
		"ModifiedAt",
		"TweetID",
		"UserID",
	}
}

func (l *Like) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "ID":
			ret = append(ret, &l.ID)
		case "Count":
			ret = append(ret, &l.Count)
		case "CreatedAt":
			ret = append(ret, &l.CreatedAt)
		case "ModifiedAt":
			ret = append(ret, &l.ModifiedAt)
		case "TweetID":
			ret = append(ret, &l.TweetID)
		case "UserID":
			ret = append(ret, &l.UserID)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (l *Like) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "ID":
			ret = append(ret, l.ID)
		case "Count":
			ret = append(ret, l.Count)
		case "CreatedAt":
			ret = append(ret, l.CreatedAt)
		case "ModifiedAt":
			ret = append(ret, l.ModifiedAt)
		case "TweetID":
			ret = append(ret, l.TweetID)
		case "UserID":
			ret = append(ret, l.UserID)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newLike_Decoder returns a decoder which reads a row from *spanner.Row
// into Like. The decoder is not goroutine-safe. Don't use it concurrently.
func newLike_Decoder(cols []string) func(*spanner.Row) (*Like, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*Like, error) {
		var l Like
		ptrs, err := l.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &l, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (l *Like) Insert(ctx context.Context) *spanner.Mutation {
	return spanner.Insert("Likes", LikeColumns(), []interface{}{
		l.ID, l.Count, l.CreatedAt, l.ModifiedAt, l.TweetID, l.UserID,
	})
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (l *Like) Update(ctx context.Context) *spanner.Mutation {
	return spanner.Update("Likes", LikeColumns(), []interface{}{
		l.ID, l.Count, l.CreatedAt, l.ModifiedAt, l.TweetID, l.UserID,
	})
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (l *Like) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	return spanner.InsertOrUpdate("Likes", LikeColumns(), []interface{}{
		l.ID, l.Count, l.CreatedAt, l.ModifiedAt, l.TweetID, l.UserID,
	})
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (l *Like) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, LikePrimaryKeys()...)

	values, err := l.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "Like.UpdateColumns", "Likes", err)
	}

	return spanner.Update("Likes", colsWithPKeys, values), nil
}

// FindLike gets a Like by primary key
func FindLike(ctx context.Context, db YORODB, id string) (*Like, error) {
	key := spanner.Key{id}
	row, err := db.ReadRow(ctx, "Likes", key, LikeColumns())
	if err != nil {
		return nil, newError("FindLike", "Likes", err)
	}

	decoder := newLike_Decoder(LikeColumns())
	l, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindLike", "Likes", err)
	}

	return l, nil
}

// ReadLike retrieves multiples rows from Like by KeySet as a slice.
func ReadLike(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*Like, error) {
	var res []*Like

	decoder := newLike_Decoder(LikeColumns())

	rows := db.Read(ctx, "Likes", keys, LikeColumns())
	err := rows.Do(func(row *spanner.Row) error {
		l, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, l)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadLike", "Likes", err)
	}

	return res, nil
}

// Delete deletes the Like from the database.
func (l *Like) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := l.columnsToValues(LikePrimaryKeys())
	return spanner.Delete("Likes", spanner.Key(values))
}

// FindLikesByTweetID retrieves multiple rows from 'Likes' as a slice of Like.
//
// Generated from index 'IDX_Likes_TweetID_EE3F39C571261A95'.
func FindLikesByTweetID(ctx context.Context, db YORODB, tweetID string) ([]*Like, error) {
	const sqlstr = "SELECT " +
		"ID, Count, CreatedAt, ModifiedAt, TweetID, UserID " +
		"FROM Likes@{FORCE_INDEX=IDX_Likes_TweetID_EE3F39C571261A95} " +
		"WHERE TweetID = @param0"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = tweetID

	decoder := newLike_Decoder(LikeColumns())

	// run query
	YOLog(ctx, sqlstr, tweetID)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*Like{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindLikesByTweetID", "Likes", err)
		}

		l, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindLikesByTweetID", "Likes", err)
		}

		res = append(res, l)
	}

	return res, nil
}

// ReadLikesByTweetID retrieves multiples rows from 'Likes' by KeySet as a slice.
//
// This does not retrives all columns of 'Likes' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'IDX_Likes_TweetID_EE3F39C571261A95'.
func ReadLikesByTweetID(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*Like, error) {
	var res []*Like
	columns := []string{
		"ID",
		"TweetID",
	}

	decoder := newLike_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "Likes", "IDX_Likes_TweetID_EE3F39C571261A95", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		l, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, l)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadLikesByTweetID", "Likes", err)
	}

	return res, nil
}

// FindLikesByUserID retrieves multiple rows from 'Likes' as a slice of Like.
//
// Generated from index 'IDX_Likes_UserID_C747C4233177C49A'.
func FindLikesByUserID(ctx context.Context, db YORODB, userID string) ([]*Like, error) {
	const sqlstr = "SELECT " +
		"ID, Count, CreatedAt, ModifiedAt, TweetID, UserID " +
		"FROM Likes@{FORCE_INDEX=IDX_Likes_UserID_C747C4233177C49A} " +
		"WHERE UserID = @param0"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = userID

	decoder := newLike_Decoder(LikeColumns())

	// run query
	YOLog(ctx, sqlstr, userID)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*Like{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindLikesByUserID", "Likes", err)
		}

		l, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindLikesByUserID", "Likes", err)
		}

		res = append(res, l)
	}

	return res, nil
}

// ReadLikesByUserID retrieves multiples rows from 'Likes' by KeySet as a slice.
//
// This does not retrives all columns of 'Likes' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'IDX_Likes_UserID_C747C4233177C49A'.
func ReadLikesByUserID(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*Like, error) {
	var res []*Like
	columns := []string{
		"ID",
		"UserID",
	}

	decoder := newLike_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "Likes", "IDX_Likes_UserID_C747C4233177C49A", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		l, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, l)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadLikesByUserID", "Likes", err)
	}

	return res, nil
}
