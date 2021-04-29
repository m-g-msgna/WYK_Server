package interfaces

import "wyk_server.src/models"

type IWYKService interface {
	Add_User(hash string) (int64, error)
	Verify_Hash(user_id int64, hash string) (int, error)
	Update_Hash(user_id int64, new_hash string) (int, error)
	Get_User_Data(user_id int64) (models.WYKUser, error)
}
