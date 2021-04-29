package services

import (
	"time"
	"strings"

	"wyk_server.src/interfaces"
	"wyk_server.src/models"
)

type WYKService struct {
	interfaces.IWYKRepository
}

/*
 *
 */
func (service *WYKService) Add_User(hash string) (int64, error) {
	current_time := time.Now().Unix();

	id, err := service.Insert_user_info(current_time);
	if err != nil {
		return 0, err
	}

	err = service.Insert_login_info(id, hash, current_time)
	if err != nil {
		return 0, err
	}

	err = service.Insert_hash_change_log(id, current_time)
	if err != nil {
		return 0, err
	}

	return id, nil
}

/*
 *
 */
func (service *WYKService) Verify_Hash(user_id int64, hash string) (int, error) {
	current_time := time.Now().Unix()
	var result int

	saved_hash, err := service.Get_cl_hash(user_id)
	if err !=  nil {
		return -100, err
	}

	if strings.TrimSpace(saved_hash) != strings.TrimSpace(hash) {
		result = -1			//Supplied hash is not the same as saved hash.
	} else {
		result = 1
	}

	err = service.Insert_auth_log(user_id, result, current_time)
	if err != nil {
		return result, err
	}

	return result, nil
}

/*
 *
 */
func (service *WYKService) Update_Hash(user_id int64, new_hash string) (int, error) {
	current_time := time.Now().Unix()

	err :=service.Update_login_info(user_id, new_hash, current_time)
	if err != nil {
		return -1, err
	}

	err = service.Insert_hash_change_log(user_id, current_time)
	if err != nil {
		return -1, err
	}

	return 1, nil
}

/*
 * Get user data.
*/
func (service *WYKService) Get_User_Data(user_id int64) (models.WYKUser, error) {
	var user_data models.WYKUser

	exists, err := service.User_Exists(user_id)
	if err != nil {
		return models.WYKUser{}, err
	}

	if !exists {
		return models.WYKUser{}, nil
	}

	user_signup_time, err1 := service.Get_User_Signup_Time( user_id )
	if err1 != nil {
		return models.WYKUser{}, err1
	}

	auth_log, err2 := service.Get_auth_log( user_id )
	if err2 != nil {
		return models.WYKUser{}, err2
	}

	hash_log, err3 := service.Get_hash_change_log( user_id )
	if err3 != nil {
		return models.WYKUser{}, err3
	}

	user_data.UserID = user_id;
	user_data.SignupTime = user_signup_time;
	user_data.AuthLogData = auth_log;
	user_data.HashChangeLogData = hash_log;

	return user_data, nil;
}