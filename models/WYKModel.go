package models

type LoginInfo struct {
	Hash				string		`json:"hash,omitempty"`
	LastUpdateTime		int64		`json:"last_update_time,omitempty"`
}

type AuthLog struct {
	AuthID				int			`json:"auth_id,omitempty"`
	AuthTime			int64		`json:"auth_time,omitempty"`
	AuthResult			int16		`json:"auth_result,omitempty"`
}

type HashChangeLog struct {
	LogID				int			`json:"log_id,omitempty"`
	ChangeTime			int64		`json:"change_time,omitempty"`
}


type WYKUser struct {
	UserID          	int64       	`json:"user_id,omitempty"`
	SignupTime			int64			`json:"signup_time,omitempty"`
	LoginData			LoginInfo		`json:"login_info,omitempty"`
	AuthLogData			[]AuthLog		`json:"auth_log,omitempty"`
	HashChangeLogData	[]HashChangeLog	`json:"hash_change_log,omitempty"`
	Code 				int 			`json:"server_code,omitempty"`
}