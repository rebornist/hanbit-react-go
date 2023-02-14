package user

import (
	"hanbit-react/types"
)

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
