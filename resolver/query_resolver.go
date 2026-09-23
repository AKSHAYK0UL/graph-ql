package resolver

import (
	"context"
	"strconv"

	"github.com/AKSHAYK0UL/graph-ql/graph/model"
	"github.com/AKSHAYK0UL/graph-ql/respository"
)

func GetUsers(ctx context.Context) ([]*model.User, error) {

	users, err := respository.GetUser()

	if err != nil {
		return nil, err
	}

	var results []*model.User

	for _, u := range users {
		results = append(results, &model.User{
			ID:    strconv.Itoa(int(u.ID)),
			Name:  u.Name,
			Email: u.Email,
		})

	}
	return results, nil
}

//https://goplay.tools/snippet/KlZmpoLri5n
