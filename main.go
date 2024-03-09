package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"graduation_project_ent/common"
	"graduation_project_ent/internal/db"
	"graduation_project_ent/pkg/ent_work/user"
)

func main() {
	ctx := context.Background()
	users, err := db.DB.User.Query().Where(user.DeletedAt(common.ZeroTime)).First(ctx)
	if err != nil {
		logrus.Errorf("%+v", err)
	}
	fmt.Printf("%+v\n", users)
}
