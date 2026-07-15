package main

import (
	"context"
	"fmt"

	"github.com/K1N3tiCs/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get the feed: %w", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get the user: %w", err)
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow the feed: %w", err)
	}

	for _, feed := range feed_follow {
		fmt.Printf("Feed Name: %s\n", feed.FeedName)
		fmt.Printf("User Name: %s\n", feed.UserName)
		fmt.Println("================================================")
	}

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage %s", cmd.Name)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
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
