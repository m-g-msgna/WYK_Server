package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"wyk_server.src/interfaces"
	"wyk_server.src/models"
)

type WYKController struct {
	interfaces.IWYKService
}

/*
 *
 */
func (controller *WYKController) Initialize(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var request_json models.WYKUser
	var response_json models.WYKUser

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request_json);
	if err != nil {
		response_json = models.WYKUser {
			UserID: -1,
			Code: http.StatusInternalServerError,
		}

		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(response_json)
				
		return
	} else {
		uid, err1 := controller.Add_User(request_json.LoginData.Hash)
		if err1 != nil {
			response_json = models.WYKUser {
				UserID: -1,
				Code: http.StatusInternalServerError,
			}
	
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(response_json)
			
			return
		}

		response_json = models.WYKUser {
			UserID: uid,
			Code: http.StatusOK,
		}
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response_json)
}

/*
 * Update hash value at the request of WYK Service from the mobile.
 */
func (controller *WYKController) Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var request_json models.WYKUser
	var response_json models.WYKUser

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request_json);
	if err != nil {
		response_json = models.WYKUser {
			UserID: -1,
			Code: http.StatusInternalServerError,
		}

		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(response_json)
		
		return
	} else {
		result, err1 := controller.Update_Hash(request_json.UserID, request_json.LoginData.Hash)
		if err1 != nil {
			response_json = models.WYKUser {
				UserID: -1,
				Code: http.StatusInternalServerError,
			}

			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(response_json)
			
			return
		}

		response_json = models.WYKUser {
			UserID: request_json.UserID,
			Code: result,
		}
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response_json)
}

/*
 * Authenticate hash value.
 */
func (controller *WYKController) Authenticate(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var request_json models.WYKUser
	var response_json models.WYKUser

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request_json);
	if err != nil {
		response_json = models.WYKUser {
			UserID: -1,
			Code: http.StatusInternalServerError,
		}

		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(response_json)
		
		return
	} else {
		result, err1 := controller.Verify_Hash(request_json.UserID, request_json.LoginData.Hash)
		if err1 != nil {
			response_json = models.WYKUser {
				UserID: -1,
				Code: http.StatusInternalServerError,
			}
	
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(response_json)
			
			return
		}

		response_json = models.WYKUser {
			UserID: request_json.UserID,
			Code: result,
		}
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response_json)
}

func (controller *WYKController) GetUserData(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	keys := req.URL.Query()
	user_id, err := strconv.Atoi( keys.Get("uid") )

	if err != nil {
		res_json := models.WYKUser {
			UserID: -1,
			Code: -1,
		}
	
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(res_json)
		
		return
	}

	response_json, err1 := controller.Get_User_Data( int64 (user_id) ) 
	if err1 != nil {
		res_json := models.WYKUser {
			UserID: -1,
			Code: -1,
		}
	
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(res_json)
		
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response_json)
}