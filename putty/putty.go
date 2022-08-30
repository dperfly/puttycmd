package putty

type PrivacyMode struct {
	PersonalGps        bool `json:"personalGps"`
	PersonalUser       bool `json:"personalUser"`
	PseudonymousGps    bool `json:"pseudonymousGps"`
	PersonalVehicle    bool `json:"personalVehicle"`
	AnonymousDataNoGPS bool `json:"anonymousDataNoGPS"`
}

type DisableMode struct {
	DisableReasons []string `json:"disableReasons"`
}

const LocalVehicleDisabled = "localVehicleDisabled"

// default login input
var username string = "sshclient"
var password string = "123456"
var host string = "192.168.4.1"
var port string = "22"
