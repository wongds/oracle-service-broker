package models

type Metadata struct {
	Costs []struct {
		Unit   string `yaml:"unit", json:"unit"`
		Amount struct {
			Value string `yaml:"value", json:"value"`
		}
	}
}
