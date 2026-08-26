package lbugruntime

import (
	"errors"
	"fmt"
	"os"
)

func DatabaseSidecarPaths(dbPath string) []string {
	return []string{
		dbPath + ".wal",
		dbPath + ".lock",
		dbPath + ".wal.checkpoint",
		dbPath + ".checkpoint.apply.lock",
		dbPath + ".checkpoint.intent.lock",
	}
}

func RemoveDatabaseSidecars(dbPath string) error {
	for _, path := range DatabaseSidecarPaths(dbPath) {
		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("remove database sidecar %s: %w", path, err)
		}
	}
	return nil
}

func RemoveDatabaseArtifacts(dbPath string) error {
	if err := os.RemoveAll(dbPath); err != nil {
		return fmt.Errorf("remove database artifact %s: %w", dbPath, err)
	}
	return RemoveDatabaseSidecars(dbPath)
}

func RunWithWALRecovery(dbPath string, operation func() error) error {
	if operation == nil {
		return fmt.Errorf("operation is nil")
	}
	err := operation()
	if err == nil || !IsWALCorruptionError(err) {
		return err
	}
	if cleanupErr := RemoveDatabaseSidecars(dbPath); cleanupErr != nil {
		return cleanupErr
	}
	return operation()
}
