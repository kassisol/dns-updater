package infoblox

type Config struct {
	IPAMHost   string `yaml:"ipam_host"`
	APIVersion string `yaml:"api_version"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}

type RecordHost struct {
	Ref     string   `json:"_ref,omitempty"`
	Aliases []string `json:"aliases"`
}

type UpdateRecordHostAliases struct {
	Aliases []string `json:"aliases"`
}
