package main

import "fmt"

func autoError() error {
	return fmt.Errorf("just a fake error")
}

// Database max connection error
func main() {
	defer fmt.Println("defer: return")
	if err := autoError(); err != nil {
		return
	}
}

/*
func (ur *usersRepo) Register(ctx context.Context, req *entities.UsersRegisterReq) (*entities.UsersRegisterRes,error) {
	tx, err := ur.Db.BeginTxx(ctx,nil)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error, transaction is not enable")
	}
	rows, err := tx.NamedQuery(query, req)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error, can't insert user")
	}
	
	user := new(entities.UsersRegisterRes)
	if rows.Next() {
		if err := rows.StructScan(user); err != nil {
			log.Println(err.Error())
			return nil, errors.New("error, can't parse user into the struct")
		}
	}
	defer rows.Close()
}
*/

/*
func (ur *usersRepo) Register(ctx context.Context, req *entities.UsersRegisterReq) (*entities.UsersRegisterRes,error) {
	tx, err := ur.Db.BeginTxx(ctx,nil)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error, transaction is not enable")
	}
	rows, err := tx.NamedQuery(query, req)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error, can't insert user")
	}
	defer rows.Close()

	user := new(entities.UsersRegisterRes)
	if rows.Next() {
		if err := rows.StructScan(user); err != nil {
			log.Println(err.Error())
			return nil, errors.New("error, can't parse user into the struct")
		}
	}
}
*/