package main

import (
	"context"
	"fmt"

	"github.com/K1N3tiCs/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage %s", cmd.Name)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get the following details")
	}

	for _, follows := range following {
		fmt.Printf("Feed Name: %s\n", follows.FeedName)
		fmt.Printf("User Name: %s\n", follows.UserName)
		fmt.Println("================================================")
	}
	return nil
}
