package authentication

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nkien0204/lets-go/configs"
	"github.com/nkien0204/lets-go/internal/db/rdb/mysql"
	"github.com/nkien0204/lets-go/internal/db/rdb/mysql/models"
	"github.com/nkien0204/lets-go/internal/log"
	"github.com/nkien0204/lets-go/internal/network/http_handler/responses"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const AccessTokenExpireTime = time.Duration(2) * time.Hour
const RefreshTokenExpireTime = time.Duration(24) * time.Hour
const AccessTokenKey string = "AccessToken"
const RefreshTokenKey string = "RefreshToken"

type Credentials struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"user_name"`
	jwt.StandardClaims
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	jwtKey := configs.GetConfigs().SecretKey.Key
	var creds Credentials
	logger := log.Logger()
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		logger.Error("decode request failed", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userModel models.User
	dbService := mysql.GetMysqlConnection()
	if result := dbService.Db.Table(models.UsersTable).Where("username = ?", creds.Username).First(&userModel); result.Error != nil {
		logger.Error("find user failed", zap.Error(result.Error))
		responses.CustomResponse(w, responses.ResRetrieveFailed, "find user failed", nil)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(creds.Password)) != nil {
		logger.Error("username or password is not correct")
		responses.CustomResponse(w, responses.ResAuthFailed, "wrong username or password", nil)
		return
	}

	accessTokenExpireTime := time.Now().Add(AccessTokenExpireTime)
	accessToken, err := generateToken(jwtKey, accessTokenExpireTime, creds.Username)
	if err != nil {
		logger.Error("generate accessToken failed")
		responses.CustomResponse(w, responses.ResGenTokenFailed, "generate accessToken failed", nil)
		return
	}
	refreshTokenExpireTime := time.Now().Add(RefreshTokenExpireTime)
	refreshToken, err := generateToken(jwtKey, refreshTokenExpireTime, creds.Username)
	if err != nil {
		logger.Error("generate refreshToken failed")
		responses.CustomResponse(w, responses.ResGenTokenFailed, "generate refreshToken failed", nil)
		return
	}

	logger.Info("sign-in successfully", zap.String("accessToken", accessToken), zap.String("refreshToken", refreshToken))
	data := map[string]string{
		AccessTokenKey:  accessToken,
		RefreshTokenKey: refreshToken,
	}
	responses.CustomResponse(w, responses.ResOk, "ok", data)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	logger := log.Logger()
	jwtKey := configs.GetConfigs().SecretKey.Key

	tknStr := r.Header.Get(AccessTokenKey)

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			logger.Error("invalid signature", zap.Error(err))
			responses.CustomResponse(w, responses.ResInvalidSignature, "invalid signature", nil)
			return
		}
		logger.Error("parse token failed", zap.Error(err))
		responses.CustomResponse(w, responses.ResParseTokenFailed, "parse token failed", nil)
		return
	}
	if !tkn.Valid {
		logger.Error("invalid token")
		responses.CustomResponse(w, responses.ResInvalidToken, "invalid token", nil)
		return
	}

	data := map[string]interface{}{
		"username":         claims.Username,
		"expire time left": claims.ExpiresAt - time.Now().Unix(),
	}
	responses.CustomResponse(w, responses.ResOk, "ok", data)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	logger := log.Logger()
	jwtKey := configs.GetConfigs().SecretKey.Key
	tknStr := r.Header.Get(RefreshTokenKey)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			logger.Error("invalid signature", zap.Error(err))
			responses.CustomResponse(w, responses.ResInvalidSignature, "invalid signature", nil)
			return
		}
		logger.Error("parse token failed", zap.Error(err))
		responses.CustomResponse(w, responses.ResParseTokenFailed, "parse token failed", nil)
		return
	}
	if !tkn.Valid {
		logger.Error("invalid token", zap.Error(err))
		responses.CustomResponse(w, responses.ResInvalidToken, "invalid token", nil)
		return
	}

	if (time.Until(time.Unix(claims.ExpiresAt, 0))) > (RefreshTokenExpireTime - AccessTokenExpireTime) {
		logger.Error("no need to refresh token", zap.Error(err))
		responses.CustomResponse(w, responses.ResInvalidToken, "no need to refresh token", nil)
		return
	}

	accessTokenExpireTime := time.Now().Add(AccessTokenExpireTime)
	accessToken, err := generateToken(jwtKey, accessTokenExpireTime, claims.Username)
	if err != nil {
		logger.Error("generate accessToken failed")
		responses.CustomResponse(w, responses.ResGenTokenFailed, "generate accessToken failed", nil)
		return
	}
	refreshTokenExpireTime := time.Now().Add(RefreshTokenExpireTime)
	refreshToken, err := generateToken(jwtKey, refreshTokenExpireTime, claims.Username)
	if err != nil {
		logger.Error("generate refreshToken failed")
		responses.CustomResponse(w, responses.ResGenTokenFailed, "generate refreshToken failed", nil)
		return
	}

	data := map[string]string{
		AccessTokenKey:  accessToken,
		RefreshTokenKey: refreshToken,
	}
	responses.CustomResponse(w, responses.ResOk, "ok", data)
}

func generateToken(jwtKey []byte, expirationTime time.Time, username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
