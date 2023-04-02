package repo

import (
	"balansnack/internal/entity"
	"database/sql"
	"errors"
	"fmt"
)

type AvatarRepo struct {
	db *sql.DB
}

func NewAvatarRepo(db *sql.DB) *AvatarRepo {
	repo := AvatarRepo{db: db}
	return &repo
}

func (r *AvatarRepo) Create(nickname string, profile *string) (*entity.Avatar, error) {
	result, err := r.db.Exec("insert into avatars (nickname, profile) values (?, ?)", nickname, profile)
	if err != nil {
		return nil, err // https://github.com/go-gorm/gorm/issues/4037
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	avatar := entity.Avatar{
		ID:       id,
		Nickname: nickname,
		Profile:  profile,
	}

	return &avatar, nil
}

func (r *AvatarRepo) UpdateNickname(id int64, nickname string) error {
	result, err := r.db.Exec("update avatars set nickname = ? where id = ?", nickname, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("no avatar updated")
	}

	return nil
}

func (r *AvatarRepo) UpdateProfile(id int64, profile *string) error {
	result, err := r.db.Exec("update avatars set profile = ? where id = ?", profile, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("no avatar updated")
	}

	return nil
}

func (r *AvatarRepo) Delete(id int64) error {
	result, err := r.db.Exec("delete from avatars where id = ?", id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("no avatar deleted")
	}

	return nil
}

func (r *AvatarRepo) GetByID(id int64) (*entity.Avatar, error) {
	var avatar entity.Avatar

	row := r.db.QueryRow("select * from avatars where id = ?", id)
	err := row.Scan(&avatar.ID, &avatar.Nickname, &avatar.Profile)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("avatar not found with id: %v", id)
		}
		return nil, err
	}

	return &avatar, nil
}

func (r *AvatarRepo) CheckDuplicateNickname(nickname string) error {
	var count int
	err := r.db.QueryRow("select count(*) from avatars where nickname = ?", nickname).Scan(&count)
	if err != nil {
		return err
	}

	if count != 0 {
		return errors.New("duplicated nickname")
	}

	return nil
}
