package services

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/boltdb/bolt"
	"golang.org/x/crypto/bcrypt"

	"github.com/simon-whitehead/react-todo/domain"
)

type UserServicer interface {
	CreateUser(email, password string) (*domain.User, error)
	GetUserById(id uint64) *domain.User
	GetUserByEmail(email string) *domain.User
	AuthenticateUser(email, password string) (*domain.User, bool)
}

type UserService struct {
	db *bolt.DB
}

func NewUserService(db *bolt.DB) UserServicer {
	return &UserService{db: db}
}

func (svc *UserService) CreateUser(email, password string) (*domain.User, error) {
	if svc.GetUserByEmail(email) != nil {
		return nil, errors.New("Email address in use")
	}

	var u *domain.User
	// Create a new user in the Users bucket
	svc.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("Users"))

		id, err := b.NextSequence()
		if err != nil {
			log.Fatal(err)
		}

		if pwd, err := encryptPassword(password); err == nil {

			u = &domain.User{Id: id, Email: email, Password: pwd}

			buf, err := json.Marshal(u)
			if err != nil {
				return err
			}

			b.Put(domain.Itob(id), buf)

		} else {
			return err
		}

		return nil
	})

	return u, nil
}

func (svc *UserService) GetUserById(id uint64) *domain.User {
	var u *domain.User
	// Find the user via their id
	svc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		buf := b.Get(domain.Itob(id))

		err := json.Unmarshal(buf, u)
		if err != nil {
			return err
		}

		return nil
	})

	return u
}

func (svc *UserService) GetUserByEmail(email string) *domain.User {
	var u *domain.User
	// Find the user via their email
	svc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {

			temp := &domain.User{}
			err := json.Unmarshal(v, temp)
			if err != nil {
				continue
			}

			if temp.Email == email {
				u = temp
				break
			}
		}

		return nil
	})

	return u
}

// Bcrypt the user password
func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (svc *UserService) AuthenticateUser(email, password string) (*domain.User, bool) {
	user := svc.GetUserByEmail(email)
	if user == nil {
		return nil, false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
		return user, true
	}

	return nil, false
}
