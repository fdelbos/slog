package main

import "github.com/fdelbos/slog"

func main() {
	slog.Enabled = true

	slog.Service("Postgres")
	slog.Value("Host", "192.168.1.1")
	slog.Value("Database", "test")
	slog.Value("User", "me")
	slog.Secret("Password", "my_secret_password")
	slog.Value("SSL", true)
	slog.Value("Pool idle", 10)
	slog.Value("Pool max", 30)

	slog.Service("redis")
	slog.Value("Host", "localhost")
	slog.Value("Port", 6379)
	slog.Secret("Password", "")
	slog.Value("DB", nil)

	slog.Service("Geometry")
	slog.Value("Size", 123.45)
	slog.Value("Maximum", uint64(999999999))
	slog.Value("Minimum", int64(-999999999))

	slog.Service("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	slog.Value("Host", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	slog.Value("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.", "test")
	slog.Value("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
}
