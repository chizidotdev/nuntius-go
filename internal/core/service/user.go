package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"github.com/chizidotdev/nuntius/config"
	"github.com/chizidotdev/nuntius/internal/core/domain"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"log"
	"net/http"
)

type UserService struct {
	repo       domain.UserRepository
	authConfig oauth2.Config
}

func NewUserService(repo domain.UserRepository) *UserService {
	gob.Register(domain.User{})

	oauthConfig := oauth2.Config{
		ClientID:     config.EnvVars.GoogleClientID,
		ClientSecret: config.EnvVars.GoogleClientSecret,
		RedirectURL:  config.EnvVars.AuthCallbackUrl,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
	return &UserService{
		repo:       repo,
		authConfig: oauthConfig,
	}
}

func (u *UserService) GetGoogleAuthConfig() oauth2.Config {
	return u.authConfig
}

func (u *UserService) GoogleCallback(ctx context.Context, code string) (*domain.User, error) {
	userData, err := u.getGoogleUserData(code)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Email:         userData.Email,
		EmailVerified: userData.VerifiedEmail,
		FirstName:     userData.GivenName,
		LastName:      userData.FamilyName,
		Image:         userData.Picture,
		GoogleID:      userData.Id,
	}

	err = u.repo.Upsert(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type UserData struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func (u *UserService) getGoogleUserData(code string) (UserData, error) {
	token, err := u.authConfig.Exchange(context.Background(), code)
	if err != nil {
		return UserData{}, err
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return UserData{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return UserData{}, err
	}

	gob.Register(UserData{})
	var user UserData
	err = json.Unmarshal(data, &user)
	if err != nil {
		return UserData{}, err
	}

	return user, nil
}

func (u *UserService) GenerateAuthState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.RawURLEncoding.EncodeToString(b)

	return state, nil
}
