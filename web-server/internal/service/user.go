package service

import (
	"fmt"

	"github.com/Sam-Frost/web-server/internal/dto"
	"github.com/Sam-Frost/web-server/internal/repo"
	"github.com/Sam-Frost/web-server/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	server *util.Server
	repo   repo.UserRepo
}

func NewUserService(server *util.Server, repo repo.UserRepo) UserService {
	return UserService{
		server: server,
		repo:   repo,
	}
}

func (u *UserService) CreateUser(requestBody dto.CreateUserRequest) (dto.CreateUserResponse, error) {

	// Response DTO
	createUserResponse := dto.CreateUserResponse{}

	// Hash Password
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 5)
	if err != nil {
		return createUserResponse, fmt.Errorf("Trying to hash password : %w", err)
	}
	passwordHash := string(hashedPasswordBytes)

	affiliateCode := util.GenerateAffiliateCode()

	if len(requestBody.AffiliateCode) != 0 {
		referredUserId, err := u.repo.GetUserIdByAffiliateCode(requestBody.AffiliateCode)
		if err != nil {
			return createUserResponse, fmt.Errorf("Affiliate code not found: %w", err)
		}
		savedUser, err := u.repo.CreateNewUserWithReferral(requestBody, affiliateCode, passwordHash, referredUserId)
		if err != nil {
			return createUserResponse, fmt.Errorf("Saving user with reffered id: %w", err)
		}

		createUserResponse.UserId = savedUser.ID

	} else {
		savedUser, err := u.repo.CreateNewUser(requestBody, affiliateCode, passwordHash)
		if err != nil {
			return createUserResponse, fmt.Errorf("Saving user without reffered id: %w", err)
		}

		createUserResponse.UserId = savedUser.ID
	}

	createUserResponse.AffiliateCode = affiliateCode

	if token, err := util.GenerateToken(string(createUserResponse.UserId)); err != nil {
		return createUserResponse, fmt.Errorf("Genearating JWT token: %w", err)
	} else {
		createUserResponse.Token = token
		return createUserResponse, nil
	}
}
