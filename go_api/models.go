package main

import (
	"database/sql"
	"fmt"
)

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserModel struct {
	DB *sql.DB
}

// GetAll retrieves all users from the database
func (m *UserModel) GetAll() ([]User, error) {
	rows, err := m.DB.Query("SELECT id, name, email, age, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (m *UserModel) Create(user *User) error {
	query := "INSERT INTO users(name, email, age, password) VALUES($1, $2, $3, $4) RETURNING id"
	err := m.DB.QueryRow(query, user.Name, user.Email, user.Age, user.Password).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("unable to insert user: %w", err)
	}
	return nil
}

func (m *UserModel) FindByEmailPassword(email string, password string) (*User, error) {
	user := &User{}
	query := "SELECT id, name, email, age, password FROM users WHERE email = $1 AND password = $2"
	err := m.DB.QueryRow(query, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (m *UserModel) Update(user *User) (*User, error) {
	query := "UPDATE users SET name = $1, password = $2, age = $3 WHERE email = $4"
	result, err := m.DB.Exec(query, user.Name, user.Password, user.Age, user.Email)
	if err != nil {
		return nil, fmt.Errorf("unable to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return user, nil
}

func (m *UserModel) Delete(email string, password string) (*User, error) {
	user := &User{}
	queryFind := "SELECT id, name, email, age, password FROM users WHERE email = $1 AND password = $2"
	err := m.DB.QueryRow(queryFind, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	query := "DELETE FROM users WHERE email = $1 AND password = $2"
	result, err := m.DB.Exec(query, email, password)
	if err != nil {
		return nil, fmt.Errorf("unable to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return user, nil
}
