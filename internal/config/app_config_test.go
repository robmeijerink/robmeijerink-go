package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	// 1. Save the original environment variables to prevent test pollution
	originalPort := os.Getenv("PORT")
	originalEnv := os.Getenv("APP_ENV")

	// 2. Defer restoration: This guarantees the environment is restored
	// exactly as it was, even if a test panics or fails.
	defer func() {
		os.Setenv("PORT", originalPort)
		os.Setenv("APP_ENV", originalEnv)
	}()

	// 3. Define the test scenarios
	tests := []struct {
		name         string
		setupPort    string
		setupAppEnv  string
		expectedProd bool
		expectedPort string
	}{
		{
			name:         "Default values when environment variables are missing",
			setupPort:    "",
			setupAppEnv:  "",
			expectedProd: false,
			expectedPort: "3000",
		},
		{
			name:         "Production mode with custom port (e.g., Google App Engine)",
			setupPort:    "8080",
			setupAppEnv:  "production",
			expectedProd: true,
			expectedPort: "8080",
		},
		{
			name:         "Development mode ignores non-production APP_ENV strings",
			setupPort:    "5000",
			setupAppEnv:  "staging", // Anything other than "production" should be false
			expectedProd: false,
			expectedPort: "5000",
		},
	}

	// 4. Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the environment for this specific test case
			if tt.setupPort == "" {
				os.Unsetenv("PORT")
			} else {
				os.Setenv("PORT", tt.setupPort)
			}

			if tt.setupAppEnv == "" {
				os.Unsetenv("APP_ENV")
			} else {
				os.Setenv("APP_ENV", tt.setupAppEnv)
			}

			// Execute the function under test
			cfg := Load()

			// Assertions
			assert.NotNil(t, cfg)
			assert.Equal(t, tt.expectedProd, cfg.IsProd, "IsProd did not match expected value")
			assert.Equal(t, tt.expectedPort, cfg.Port, "Port did not match expected value")
		})
	}
}
