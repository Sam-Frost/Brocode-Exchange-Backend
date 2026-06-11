package repo

import (
	"context"
	"fmt"

	db "github.com/Sam-Frost/db/generated"
	"github.com/Sam-Frost/web-server/internal/dto"
	"github.com/Sam-Frost/web-server/internal/util"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepo struct {
	server *util.Server
}

func NewUserRepo(server *util.Server) UserRepo {
	return UserRepo{
		server: server,
	}
}

func (u *UserRepo) GetUserIdByAffiliateCode(affiliateCode string) (int32, error) {
	userId, err := u.server.Query.FindUserIdByAffiliateCode(context.Background(), affiliateCode)

	if err != nil {
		return 0, fmt.Errorf("GetUserIdByAffiliateCode: %w", err)
	}

	return userId, nil
}

func (u *UserRepo) CreateNewUser(requestBody dto.CreateUserRequest, affiliateCode, passwordHash string) (db.User, error) {
	user, err := u.server.Query.CreateUser(context.Background(), db.CreateUserParams{
		Name:  requestBody.Name,
		Email: requestBody.Email,
		BirthDate: pgtype.Date{
			Time:  requestBody.BirthDate,
			Valid: true,
		},
		AffiliateCode: affiliateCode,
		PasswordHash:  passwordHash,
	})

	if err != nil {
		return db.User{}, fmt.Errorf("Saving user to databse without referrer ID: %w", err)
	}

	return user, nil
}

func (u *UserRepo) CreateNewUserWithReferral(requestBody dto.CreateUserRequest, affiliateCode, passwordHash string, referrerId int32) (db.User, error) {
	user, err := u.server.Query.CreateUserWithReferral(context.Background(), db.CreateUserWithReferralParams{
		Name:  requestBody.Name,
		Email: requestBody.Email,
		BirthDate: pgtype.Date{
			Time:  requestBody.BirthDate,
			Valid: true,
		},
		AffiliateCode: affiliateCode,
		PasswordHash:  passwordHash,
		ReferrerID: pgtype.Int4{
			Int32: referrerId,
			Valid: true,
		},
	})

	if err != nil {
		return db.User{}, fmt.Errorf("Saving user to databse with referrer ID: %w", err)
	}

	return user, nil
}
