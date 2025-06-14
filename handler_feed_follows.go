package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Rishan-Jadva/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("cannot retrieve user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("cannot retrieve feed: %w", err)
	}

	ffRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow row: %w", err)
	}

	printFeedFollow(ffRow.UserName, ffRow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	ffs, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follow for user: %w", err)
	}

	fmt.Println("Followed feeds for current user:")
	for _, ff := range ffs {
		fmt.Printf("%s\n", ff.FeedName)
	}
	return nil
}

func printFeedFollow(username string, feedname string) {
	fmt.Printf("Username:     %s\n", username)
	fmt.Printf("Feed Name:    %s\n", feedname)
}
