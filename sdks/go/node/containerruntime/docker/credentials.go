package docker

import (
	"github.com/docker/docker/api/types/registry"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/distribution/reference"
	"github.com/pkg/errors"
)

func getAuthFromConfig(configPath, imageRef string) (string, error) {
	explicitlyPassedConfig := true
	if configPath == "" {
		home, err := os.UserHomeDir()
		if nil != err {
			return "", err
		}
		configPath = filepath.Join(home, ".docker", "config.json")
		explicitlyPassedConfig = false
	}

	dockerConfig := configfile.New(configPath)
	configFile, err := os.OpenFile(configPath, os.O_RDONLY, fs.FileMode(0))
	if errors.Is(err, os.ErrNotExist) {
		if explicitlyPassedConfig {
			return "", errors.Wrapf(err, "failed to find docker config file at %s", configPath)
		}
		// ignore missing configs if not explicitly passed, since some people don't have it set up
	} else if err != nil {
		return "", errors.Wrap(err, "failed to open docker config file")
	} else {
		if err := dockerConfig.LoadFromReader(configFile); err != nil {
			return "", errors.Wrap(err, "failed to load docker config")
		}
	}

	parsedImageRef, err := reference.ParseNormalizedNamed(imageRef)
	if err != nil {
		return "", errors.Wrap(err, "invalid docker image reference")
	}
	domain := reference.Domain(parsedImageRef)

	authConfig, err := dockerConfig.GetAuthConfig(domain)
	if err != nil {
		return "", errors.Wrap(err, "failed to get docker auth")
	}

	// add an explicit fallback for docker hub, which stores it's domain in config differently than the cli api returns
	if authConfig.ServerAddress == "" && domain == "docker.io" {
		authConfig, err = dockerConfig.GetAuthConfig("https://index.docker.io/v1/")
		if err != nil {
			return "", errors.Wrap(err, "failed to get fallback docker auth for docker.io")
		}
	}

	return registry.EncodeAuthConfig(registry.AuthConfig{
		Username: authConfig.Username,
		Password: authConfig.Password,
	})
}
