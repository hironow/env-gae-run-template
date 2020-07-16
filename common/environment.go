package common

import (
	"fmt"

	"github.com/mchmarny/gcputil/meta"
)


func GetProjectID(c *Config) (string, error) {
	if !c.IsGAE() && !c.IsRun() {
		return "", fmt.Errorf("not gcp service")
	}

	name, err := meta.GetClient("my-user-agent").ProjectID()
	return name, err
}

func GetZone(c *Config) (string, error) {
	if !c.IsGAE() && !c.IsRun() {
		return "", fmt.Errorf("not gcp service")
	}

	name, err := meta.GetClient("my-user-agent").Zone()
	return name, err
}

// TODO: これがbatchのものにもおよぶときどうやって切り替える？

func GetHostName(c *Config) string {
	if c.IsGAE() {
		projectID, _ := GetProjectID(c)
		return fmt.Sprintf("https://%s-dot-%s-dot-%s.appspot.com/", c.GAE.Version, c.GAE.Service, projectID)
	}
	return ""
}