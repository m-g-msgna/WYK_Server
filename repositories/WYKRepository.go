package repositories

import (
	"fmt"
	//"time"

	//"wyk_server.src/infrastructures"
	"wyk_server.src/interfaces"
	"wyk_server.src/models"
)

type WYKRepository struct {
	interfaces.IDbHandler
}

/*
 *
 */
func (repository *WYKRepository) Insert_user_info(current_time int64) (int64, error){

	res, err := repository.Execute(fmt.Sprintf("INSERT INTO user_info (signup_time) VALUES (%d)", current_time));

	if err != nil {
		return 0, err
	}

	user_id, err1 := res.LastInsertId();

	if err1 != nil {
		return 0, err1
	}

	return user_id, nil
}

/*
 *
 */
func (repository *WYKRepository) Insert_login_info( user_id int64, hash string, current_time int64) error {
	_, err := repository.Execute( fmt.Sprintf("INSERT INTO login_info(user_id, cl_hash, last_update_time) VALUES (%d, '%s', %d)",
			user_id, hash, current_time ) );

	if err != nil {
		return err
	}

	return nil
}

/*
 *
 */
func (repository *WYKRepository) Insert_hash_change_log(used_id int64, change_time int64) error {
	_, err := repository.Execute(fmt.Sprintf("INSERT INTO hash_change_log(user_id, change_time) VALUES (%d, %d)",
								used_id, change_time ));
	if err != nil {
		return err
	}
	return nil
}

/*
 * 
 */
func (repository *WYKRepository) Update_login_info(user_id int64, new_hash string, current_time int64) (error) {
	_, err := repository.Execute( fmt.Sprintf("UPDATE login_info SET cl_hash = '%s' WHERE user_id = %d", 
					new_hash, user_id) )
	if err !=  nil {
		return err
	}
	return nil
}

func (repository *WYKRepository) Insert_auth_log(user_id int64, result int, current_time int64) (error) {
	_, err := repository.Execute(fmt.Sprintf("INSERT INTO auth_log (user_id, result, auth_time) VALUES (%d, %d, %d) ",
					user_id, result, current_time ));

	if err !=  nil {
		return err
	}

	return nil
}

func (repository *WYKRepository) Get_cl_hash (user_id int64) (string, error) {
	var hash string

	row, err := repository.Query(fmt.Sprintf("SELECT cl_hash FROM login_info WHERE user_id = %d", user_id ))
	if err != nil {
		return "", err
	}

	for row.Next() {
		err = row.Scan( &hash )
		if err != nil {
			return "", err
		}
	}

	return hash, nil
}

/*
 * Check if user already exists in the database
*/
func (repository *WYKRepository) User_Exists(user_id int64) (bool, error) {
	row, err := repository.Query( fmt.Sprintf("SELECT COUNT(*) FROM user_info WHERE user_id = %d", user_id ))
	if err != nil {
		return false, err;
	}

	if row.Next() {
		return true, nil
	}

	return false, nil;
}

/*
 * Get the time the user signed up for this service
*/
func(repository *WYKRepository) Get_User_Signup_Time(used_id int64) (int64, error) {
	var signup_time int64

	row, err := repository.Query( fmt.Sprintf("SELECT signup_time FROM user_info WHERE user_id = %d", used_id) )
	if err != nil {
		return 0, err
	}

	for row.Next() {
		err = row.Scan( &signup_time )
		if err !=  nil {
			return 0, err
		}
	}
	
	return signup_time, nil;
}

/*
 *
*/
func (repository *WYKRepository) Get_auth_log(user_id int64) ([]models.AuthLog, error) {
	var auth_log []models.AuthLog;
	var log_count int

	rows, err := repository.Query( fmt.Sprintf("SELECT COUNT(*) FROM auth_log WHERE user_id = %d", user_id ))
	if err != nil {
		return []models.AuthLog{}, err;
	}
	for rows.Next() {
		err = rows.Scan(&log_count)
		if err != nil {
			return []models.AuthLog{}, err;
		}
	}
	auth_log = make( []models.AuthLog, log_count, log_count );

	rows, err = repository.Query( fmt.Sprintf("SELECT auth_id, auth_time, result FROM auth_log WHERE user_id = %d", user_id ))
	if err != nil{
		return []models.AuthLog{}, err;
	}

	index := 0;
	for rows.Next() {
		err = rows.Scan( &auth_log[index].AuthID, &auth_log[index].AuthTime, &auth_log[index].AuthResult );
		if err != nil {
			return []models.AuthLog{}, err;
		}
		index++;
	}

	return auth_log, nil;
}

/*
 *
*/
func (repository *WYKRepository) Get_hash_change_log(user_id int64) ([]models.HashChangeLog, error) {
	var hash_log []models.HashChangeLog;
	var log_count int

	rows, err := repository.Query( fmt.Sprintf("SELECT COUNT(*) FROM hash_change_log WHERE user_id = %d", user_id ))
	if err != nil {
		return []models.HashChangeLog{}, err;
	}
	for rows.Next() {
		err = rows.Scan(&log_count)
		if err != nil {
			return []models.HashChangeLog{}, err;
		}
	}
	hash_log = make( []models.HashChangeLog, log_count, log_count );

	rows, err = repository.Query( fmt.Sprintf("SELECT log_id, change_time FROM hash_change_log WHERE user_id = %d", user_id ))
	if err != nil{
		return []models.HashChangeLog{}, err;
	}

	index := 0;
	for rows.Next() {
		err = rows.Scan( &hash_log[index].LogID, &hash_log[index].ChangeTime );
		if err != nil {
			return []models.HashChangeLog{}, err;
		}
		index++;
	}

	return hash_log, nil;
}




















