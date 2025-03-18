package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func ScanDirectories(ctx context.Context, directories []string) error {
	for _, directory := range directories {
		slog.Info("Beginning scan of directory", slog.String("directory", directory))
		err := walkDirectory(ctx, directory)
		if err != nil {
			slog.Error("Error while scanning a directory from the config", slog.String("directory", directory), slog.Any("error", err))
		}
		slog.Info("Finished scan of directory", slog.String("directory", directory))
	}

	return nil
}

func walkDirectory(ctx context.Context, directory string) error {
	return filepath.WalkDir(directory, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			slog.Error("Error accessing path", slog.String("path", path), slog.Any("error", err))
			return nil // If there was an issue accessing the path, we can just ignore the error and continue with the Walk
		}

		// Check if context is cancelled, and finish scan early.
		if errors.Is(ctx.Err(), context.Canceled) {
			slog.Info("Context cancelled mid-scan, stopping scan early.")
			return filepath.SkipAll
		}

		// Skip directories that we know aren't going to contain git repos (at least not ones we'd want to index)
		if info.Name() == "node_modules" {
			return filepath.SkipDir
		}

		// Check if the current path is a directory named ".git"
		if info.IsDir() && info.Name() == ".git" {
			slog.Info("Git repository found", slog.String("path", filepath.Dir(path)))

			// Gather information and store on a DB
			err = gatherGitInfo(path)
			if err != nil {
				// Log error but continue
				slog.Error("Error while gathering Git Repo Info", slog.String("path", filepath.Dir(path)), slog.Any("error", err))
			}

			return filepath.SkipDir // No need to walk through the .git folder
		}
		return nil
	})
}

func gatherGitInfo(path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	remotes, err := r.Remotes()
	if err != nil {
		return err
	}

	// List of remotes and their URLs
	for _, rem := range remotes {
		slog.Info("Found Remote on Repository", slog.String("path", filepath.Dir(path)), slog.String("remote_url", rem.Config().URLs[0]), slog.String("remote_url", rem.Config().Name))
	}

	// Get the HEAD reference (active branch)
	headRef, err := r.Head()
	if err != nil {
		return err
	}

	// Get the current branch name
	slog.Info("Current branch on Repository", slog.String("path", filepath.Dir(path)), slog.String("branch", headRef.Name().Short()))
	return nil
}
