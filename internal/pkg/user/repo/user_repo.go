package repo

import (
	"bytes"
	"context"
	"crypto/rand"
	"database/sql"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/argon2"

	"github.com/vmedinskiy/highload/internal/pkg/dbutils"
	"github.com/vmedinskiy/highload/internal/pkg/httperr"
	"github.com/vmedinskiy/highload/internal/pkg/user"
)

var (
	errorBadLoginPassword = errors.New("bad login/password pair")
)

type Repo struct {
	DB     sqlx.ExtContext
	saltLn int
}

func NewRepo(db sqlx.ExtContext, saltLn int) *Repo {
	return &Repo{
		DB:     db,
		saltLn: saltLn,
	}
}

func (r *Repo) Create(ctx context.Context, user *user.User) (int64, error) {
	salt := r.makeSalt()
	user.HashedPassword = r.hashPass(user.Password, salt)
	q := `insert 
		  into public.users 
			(first_name, second_name, age, biography, city, pwdhsh) 
          values
			(:first_name, :second_name, :age, :biography, :city, :pwdhsh) 
		  returning id`

	var id int64
	if err := dbutils.NamedGet(ctx, r.DB, &id, q, user); err != nil {
		return -1, errors.Wrap(err, "creating new user")
	}
	return id, nil
}

func (r *Repo) GetByID(ctx context.Context, id int64) (*user.User, error) {
	q := `select id, first_name, second_name, age, biography, city, pwdhsh from public.users where id =:id`
	dest := &user.User{}
	args := map[string]any{
		"id": id,
	}
	err := dbutils.NamedGet(ctx, r.DB, dest, q, args)
	switch {
	case err != nil && errors.Is(err, sql.ErrNoRows):
		return nil, httperr.Wrap(errors.Errorf("no user with id %d", id), http.StatusNotFound)
	case err != nil:
		return nil, errors.Wrap(err, "getting user by id")
	}
	return dest, nil
}

func (r *Repo) Search(ctx context.Context, firstName, secondName string) ([]*user.User, error) {
	if firstName == "" && secondName == "" {
		return nil, httperr.Wrap(errors.New("no conditions to search"), http.StatusBadRequest)
	}
	var searchConds string
	if firstName != "" {
		firstName += "%"
	}
	if secondName != "" {
		secondName += "%"
	}
	args := map[string]any{}
	switch {
	case firstName != "" && secondName != "":
		args["firstName"] = firstName
		args["secondName"] = secondName
		searchConds = "where first_name like :firstName and second_name like :secondName"
	case firstName == "" && secondName != "":
		args["secondName"] = secondName
		searchConds = "where second_name like :secondName"
	case firstName != "" && secondName == "":
		args["firstName"] = firstName
		searchConds = "where first_name like :firstName"
	}
	q := `select id, first_name, second_name, age, biography, city, pwdhsh from public.users ` + searchConds
	var dest []*user.User
	err := dbutils.NamedSelect(ctx, r.DB, &dest, q, args)
	if err != nil {
		return nil, errors.Wrap(err, "searching users")
	}
	return dest, nil
}

func (r *Repo) makeSalt() []byte {
	salt := make([]byte, r.saltLn)
	_, _ = rand.Read(salt)
	return salt
}

func (r *Repo) hashPass(plainPassword string, salt []byte) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	res := salt
	res = append(res, hashedPass...)
	return res
}

func (r *Repo) CheckPasswordByID(ctx context.Context, id int64, pass string) (*user.User, error) {
	resUser, err := r.GetByID(ctx, id)
	switch {
	case err != nil && errors.Is(err, errorBadLoginPassword):
		return nil, httperr.Wrap(err, http.StatusUnauthorized) // 401 no such user
	case err != nil:
		return nil, err // 500 something wrong
	}
	if err := r.isPasswordValid(pass, resUser); err != nil {
		return nil, httperr.Wrap(err, http.StatusUnauthorized) // 401 not valid pass
	}
	return resUser, nil
}

func (r *Repo) isPasswordValid(pass string, user *user.User) error {
	dbPass := user.HashedPassword
	salt := make([]byte, r.saltLn, len(dbPass))
	_ = copy(salt, dbPass[:r.saltLn])
	passHash := r.hashPass(pass, salt)
	if !bytes.Equal(passHash, dbPass) {
		return errorBadLoginPassword
	}
	return nil
}
