package notifications

type Notifications struct {
	DeviceUUID    string   `json:"device_uuid"`
	CompaniesUUID []string `json:"companies_uuid"`
	Intern        bool     `json:"intern"`
	Entry         bool     `json:"entry"`
	Mid           bool     `json:"mid"`
	Senior        bool     `json:"senior"`
	Manager       bool     `json:"manager"`
}

type NotificationsUpdate struct {
	DeviceUUID     string   `json:"device_uuid"`
	CompaniesUUID  []string `json:"companies_uuid"`
	RemoveCompUUID []string `json:"remove_comp_uuid"`
	Intern         bool     `json:"intern"`
	Entry          bool     `json:"entry"`
	Mid            bool     `json:"mid"`
	Senior         bool     `json:"senior"`
	Manager        bool     `json:"manager"`
}
