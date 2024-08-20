package commands

import "log/slog"

func logStorageError(logger *slog.Logger, command string, action string, err error) {
	logger.Warn(
		"Unexpected Storage Error",
		"Command", command,
		"Action", action,
		"Err", err,
	)
}
