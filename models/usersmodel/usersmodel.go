package usersmodel

import (
	"LocalFood/config"
	"LocalFood/entities"
)

func GetAll() []entities.Users {
	rows, err := config.DB.Query("SELECT * FROM users ORDER BY created_at DESC")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var Users []entities.Users

	for rows.Next() {
		var user entities.Users
		if err := rows.Scan(&user.Id_users, &user.Role, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}

		Users = append(Users, user)
	}

	return Users
}

func Create(formUsers entities.Users) bool {
	result, err := config.DB.Exec(`
	INSERT INTO Users (email, username, password, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)`,
		formUsers.Email, formUsers.Username, formUsers.Password, formUsers.CreatedAt, formUsers.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertID > 0
}

func GetByUsername(username string) (entities.Users, error) {
	var user entities.Users

	err := config.DB.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(
		&user.Id_users,
		&user.Role,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
