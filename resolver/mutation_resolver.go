package resolver

import (
	"context"
	"strconv"

	"github.com/AKSHAYK0UL/graph-ql/graph/model"
	"github.com/AKSHAYK0UL/graph-ql/models"
	"github.com/AKSHAYK0UL/graph-ql/respository"
)

func CreateUser(ctx context.Context, name, email string) (*model.User, error) {
	user := models.User{
		Name:  name,
		Email: email,
	}

	err := respository.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    strconv.Itoa(int(user.ID)),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func DeleteUser(ctx context.Context, id string) (bool, error) {
	return respository.DeleteUser(id)
}
