package services

type ServiceInfo struct {
	ServiceName   string `json:"service_name"`
	PlanName      string `json:"plan_name"`
	Uri            string `json:"uri"`
	AdminUser     string `json:"admin_user,omitempty"`
	AdminPassword string `json:"admin_assword,omitempty"`
	Database       string `json:"database,omitempty"`
	User           string `json:"user"`
	Password       string `json:"password"`
}
