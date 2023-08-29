package service

import (
	"gorm.io/gorm"
	"log"
	"main/pkg/models"
	"time"
)

type Service struct {
	Db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{Db: db}
}

func (srv Service) CheckUser(login, password string) (bool, int) {
	sqlQuery := `select * from users`
	Users := make([]models.User, 0)
	err := srv.Db.Raw(sqlQuery).Scan(&Users).Error
	if err != nil {
		log.Println(err)
		return false, 0
	}
	for _, user := range Users {
		if user.Login == login && user.Password == password {
			return true, user.Id
		}
	}
	return false, 0
}

func (srv Service) AddNoteToDb(userId int, newContent models.Note) error {
	sqlQuery := `insert into notes (user_id, content)
values (?,?)`
	err := srv.Db.Exec(sqlQuery, userId, newContent.Content).Error
	if err != nil {
		return err
	}
	return nil
}

func (srv Service) Read(userId, id int) (*models.Note, error) {
	sqlQuery := `select * from notes where user_id = ? and id = ?`

	var Note models.Note
	err := srv.Db.Raw(sqlQuery, userId, id).Scan(&Note).Error
	if err != nil {
		return nil, err
	}
	return &Note, nil
}

func (srv Service) UpdateNote(userId, id int, NewContent models.Note) (*models.Note, error) {
	sqlQuery := `update notes set content = ?
where user_id = ? and id = ?`
	err := srv.Db.Exec(sqlQuery, NewContent.Content, userId, id).Error
	if err != nil {
		return nil, err
	}

	var Note models.Note
	sqlQuery = `select * from notes where user_id = ? and id = ?`
	err = srv.Db.Raw(sqlQuery, userId, id).Scan(&Note).Error
	if err != nil {
		return nil, err
	}
	return &Note, nil
}

func (srv Service) DeleteNote(userId, id int) error {
	sqlQuery := `update notes 
set active = ?, deleted_at = ?
where user_id = ? and id = ?`
	return srv.Db.Exec(sqlQuery, false, time.Now(), userId, id).Error
}
