package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Rishan-Jadva/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	rssFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	ffRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    rssFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow")
	}
	printFields(rssFeed)
	printFeedFollow(ffRow.UserName, ffRow.FeedName)
	return nil
}

func printFields(feed database.Feed) {
	fmt.Printf("* ID:          %s\n", feed.ID)
	fmt.Printf("* Created:     %s\n", feed.CreatedAt)
	fmt.Printf("* Updated:     %s\n", feed.UpdatedAt)
	fmt.Printf("* Name:        %s\n", feed.Name)
	fmt.Printf("* URL:         %s\n", feed.Url)
	fmt.Printf("* UserID:      %s\n", feed.UserID)
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		return fmt.Errorf("no feeds found: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name:     %s\n", feed.Name)
		fmt.Printf("Feed URL:      %s\n", feed.Url)

		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("no user with given ID: %w", err)
		}

		fmt.Printf("User:          %s\n", user.Name)
	}
	return nil
}
