package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/K1N3tiCs/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		log.Fatalf("user doesn't exist in database")
	}

	if name == s.cfg.CurrentUserName {
		log.Fatalf("%s has already logged-in", name)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	_, err := s.db.GetUser(context.Background(), createUserParams.Name)
	if err == nil {
		log.Fatal("user already exists")
	}

	user, err := s.db.CreateUser(context.Background(), createUserParams)
	if err != nil {
		return errors.New("failed to create user")
	}

	return s.cfg.SetUser(user.Name)
}

func handlerDeleteAll(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		log.Fatal("failed to delete users")
	}

	log.Print("Deleted all the users")
	return nil
}

func handlerUsersList(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
