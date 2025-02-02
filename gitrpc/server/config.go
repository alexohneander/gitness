// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"errors"
	"time"
)

const (
	ModeInMemory = "inmemory"
	ModeRedis    = "redis"
	ModeNone     = "none"
)

// Config represents the configuration for the gitrpc server.
type Config struct {
	// Port specifies the port used to bind the grpc server.
	Port int `envconfig:"GITRPC_SERVER_PORT" default:"3001"`
	// GitRoot specifies the directory containing git related data (e.g. repos, ...)
	GitRoot string `envconfig:"GITRPC_SERVER_GIT_ROOT"`
	// TmpDir (optional) specifies the directory for temporary data (e.g. repo clones, ...)
	TmpDir string `envconfig:"GITRPC_SERVER_TMP_DIR"`
	// GitHookPath points to the binary used as git server hook.
	GitHookPath string `envconfig:"GITRPC_SERVER_GIT_HOOK_PATH"`

	HTTP struct {
		Port int `envconfig:"GITRPC_SERVER_HTTP_PORT" default:"4001"`
	}
	MaxConnAge      time.Duration `envconfig:"GITRPC_SERVER_MAX_CONN_AGE" default:"630720000s"`
	MaxConnAgeGrace time.Duration `envconfig:"GITRPC_SERVER_MAX_CONN_AGE_GRACE" default:"630720000s"`

	// LastCommitCache holds configuration options for the last commit cache.
	LastCommitCache struct {
		// Mode determines where the cache will be. Valid values are "inmemory" (default), "redis" or "none".
		Mode string `envconfig:"GITRPC_LAST_COMMIT_CACHE_MODE" default:"inmemory"`

		// DurationSeconds defines cache duration in seconds of last commit, default=12h.
		DurationSeconds int `envconfig:"GITRPC_LAST_COMMIT_CACHE_SECONDS" default:"43200"`
	}

	Redis struct {
		Endpoint           string `envconfig:"GITRPC_REDIS_ENDPOINT"             default:"localhost:6379"`
		MaxRetries         int    `envconfig:"GITRPC_REDIS_MAX_RETRIES"          default:"3"`
		MinIdleConnections int    `envconfig:"GITRPC_REDIS_MIN_IDLE_CONNECTIONS" default:"0"`
		Password           string `envconfig:"GITRPC_REDIS_PASSWORD"`
		SentinelMode       bool   `envconfig:"GITRPC_REDIS_USE_SENTINEL"         default:"false"`
		SentinelMaster     string `envconfig:"GITRPC_REDIS_SENTINEL_MASTER"`
		SentinelEndpoint   string `envconfig:"GITRPC_REDIS_SENTINEL_ENDPOINT"`
	}
}

func (c *Config) Validate() error {
	if c == nil {
		return errors.New("config is required")
	}
	if c.Port < 0 {
		//nolint: stylecheck // that's the name of the field
		return errors.New("Port is required")
	}
	if c.GitRoot == "" {
		return errors.New("GitRoot is required")
	}
	if c.GitHookPath == "" {
		return errors.New("GitHookPath is required")
	}
	if c.MaxConnAge == 0 {
		return errors.New("MaxConnAge is required")
	}
	if c.MaxConnAgeGrace == 0 {
		return errors.New("MaxConnAgeGrace is required")
	}
	if m := c.LastCommitCache.Mode; m != "" && m != ModeInMemory && m != ModeRedis && m != ModeNone {
		return errors.New("LastCommitCache.Mode has unsupported value")
	}

	return nil
}
