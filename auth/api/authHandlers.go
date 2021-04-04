package api

import (
	"encoding/json"
	"net/http"

	"github.com/alekssro/banking/auth/application"
	"github.com/alekssro/banking/lib/dto"
	"github.com/alekssro/banking/lib/logger"
)

type AuthHandler struct {
	service application.AuthService
}

func (h AuthHandler) NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	WriteResponse(w, http.StatusOK, "Handler not implemented...")
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, appErr := h.service.Login(loginRequest)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			WriteResponse(w, http.StatusOK, *token)
		}
	}
}

/*
  Sample URL string
 http://localhost:8181/auth/verify?token=somevalidtokenstring&routeName=GetCustomer&customer_id=2000&account_id=95470
*/
func (h AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	urlParams := make(map[string]string)

	// converting from Query to map type
	for k := range r.URL.Query() {
		urlParams[k] = r.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {
		appErr := h.service.Verify(urlParams)
		if appErr != nil {
			WriteResponse(w, appErr.Code, notAuthorizedResponse(appErr.Message))
		} else {
			WriteResponse(w, http.StatusOK, authorizedResponse())
		}
	} else {
		WriteResponse(w, http.StatusForbidden, notAuthorizedResponse("missing token"))
	}
}

func (h AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var refreshRequest dto.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&refreshRequest); err != nil {
		logger.Error("Error while decoding refresh token request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, appErr := h.service.Refresh(refreshRequest)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			WriteResponse(w, http.StatusOK, *token)
		}
	}
}

func notAuthorizedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"isAuthorized": false,
		"message":      msg,
	}
}

func authorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": true}
}
