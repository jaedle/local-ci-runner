package assets

import "embed"

//go:embed builder/*
var assets embed.FS

func GetString(path string) (string, error) {
	asset, err := assets.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(asset), nil
}
