package utils

import "strings"

func ReleaseAssetURL(repoURL, tag, asset string) string {
	return repoURL + "/releases/download/" + tag + "/" + asset
}

func ReleasesAPIURL(repoURL string) string {
	return strings.Replace(repoURL, "https://github.com/", "https://api.github.com/repos/", 1) + "/releases"
}
