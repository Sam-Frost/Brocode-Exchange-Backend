package dto

import "time"

type CreateUserRequest struct {
	Name          string    `json:"name" binding:"required"`
	Email         string    `json:"email" binding:"required,email"`
	BirthDate     time.Time `json:"birthDate" binding:"required"`
	Password      string    `json:"password" binding:"required"`
	AffiliateCode string    `json:"affiliateCode"`
}

type CreateUserResponse struct {
	UserId        int32  `json:"userId"`
	Token         string `json:"token"`
	AffiliateCode string `json:"affiliateCode"`
}
