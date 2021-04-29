package interfaces

import "wyk_server.src/models"

type IWYKRepository interface {
	Insert_user_info(current_time int64) (int64, error)
	Insert_login_info(user_id int64, hash string, current_time int64 ) (error)
	Update_login_info(user_id int64, new_hash string, current_time int64) (error)
	Insert_hash_change_log(used_id int64, change_time int64) (error)
	Insert_auth_log(user_id int64, result int, current_time int64) (error)
	Get_cl_hash (user_id int64) (string, error)

	User_Exists(user_id int64) (bool, error)
	Get_User_Signup_Time(used_id int64) (int64, error)
	Get_auth_log(user_id int64) ([]models.AuthLog, error)
	Get_hash_change_log(user_id int64) ([]models.HashChangeLog, error)
}
