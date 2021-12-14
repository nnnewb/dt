//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	if err := sh.Run("go", "mod", "vendor"); err != nil {
		return err
	}

	if err := sh.Run("docker-compose", "build"); err != nil {
		return err
	}

	return nil
}

func Up() error {
	if err := sh.Run("docker-compose", "up", "-d"); err != nil {
		return err
	}
	return nil
}
