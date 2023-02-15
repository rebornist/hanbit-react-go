package user

import (
	"fmt"
	"hanbit-react/types"
)

// 모든 유저 정보를 가져오는 함수
func (u *User) GetAllUsers() ([]types.User, error) {
	// fetch users from db
	rows, err := u.Query("SELECT id, phone, username, email FROM public.tb_users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// prepare answer
	users := []types.User{}
	for rows.Next() {
		// scan the result into Player struct
		u := types.User{}
		err = rows.Scan(&u.ID, &u.PhoneNumber, &u.Username, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// 유저명을 통해 특정 유저 정보를 가져오는 함수
func (u *User) GetUserByUsername(username string) (types.User, error) {
	// fetch user from db
	row := u.QueryRow(fmt.Sprintf("SELECT id, phone, username, email, password FROM public.tb_users WHERE username = '%s'", username))

	// prepare answer
	user := types.User{}
	err := row.Scan(&user.ID, &user.PhoneNumber, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return types.User{}, err
	}

	return user, nil
}
