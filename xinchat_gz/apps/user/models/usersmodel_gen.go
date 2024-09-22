// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsersIdPrefix    = "cache:users:id:"
	cacheUsersPhonePrefix = "cache:users:phone:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Users, error)
		FindOneByPhone(ctx context.Context, phone string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id string) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id        string         `db:"id"`
		Avatar    string         `db:"avatar"`
		Nickname  string         `db:"nickname"`
		Phone     string         `db:"phone"`
		Password  sql.NullString `db:"password"`
		Status    sql.NullInt64  `db:"status"`
		Sex       sql.NullInt64  `db:"sex"`
		CreatedAt sql.NullTime   `db:"created_at"`
		UpdatedAt sql.NullTime   `db:"updated_at"`
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id string) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, data.Phone)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usersIdKey, usersPhoneKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id string) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, usersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByPhone(ctx context.Context, phone string) (*Users, error) {
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, phone)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersPhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, phone); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, data.Phone)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Avatar, data.Nickname, data.Phone, data.Password, data.Status, data.Sex)
	}, usersIdKey, usersPhoneKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, data.Phone)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Avatar, newData.Nickname, newData.Phone, newData.Password, newData.Status, newData.Sex, newData.Id)
	}, usersIdKey, usersPhoneKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}