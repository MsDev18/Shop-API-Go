package user

import (
	"context"
	userdto "shop/internal/dto/user"
	"shop/internal/entity"
)

func (s Service) UpdateProfile(ctx context.Context, userID uint, req userdto.UpdateProfileRequest) error {
	const op = "user-service-UpdateProfile"
	// 1. process image and upload in server
	avatarURI, processImageErr := s.imageProcessor.ProcessAvatar(req.Avatar, userID)
	if processImageErr != nil {
		return processImageErr
	}
	// 2. create user
	user := entity.User{
		ID:     userID,
		Name:   req.Name,
		Avatar: avatarURI,
	}
	// 3. update user record in database
	// with updateProfile method
	repoErr := s.repository.UpdateProfile(ctx, user)
	if repoErr != nil {
		return repoErr
	}
	return nil
}
