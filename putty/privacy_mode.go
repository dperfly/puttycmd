package putty

import (
	"encoding/json"
	"fmt"
	"puttycmd/constant"
	"strconv"
	"strings"
)

var PrivacyModeUrl = fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.PrivacysetupAddress, constant.Privacy)

func PrivacySSH(pm PrivacyMode, sb *strings.Builder) {
	jsondata, err := json.Marshal(pm)
	if err != nil {
		sb.WriteString(err.Error())
		return
	}
	strJson := strconv.Quote(string(jsondata))
	session, err := SSH(host, port, username, password)
	if err != nil {
		sb.WriteString("...Run failed: ")
		sb.WriteString(err.Error())
		return
	}
	cmd := fmt.Sprintf("cd /mnt; ./curl -H \"Content-Type: application/json\" -X POST %s -d %s", PrivacyModeUrl, strJson)
	sb.WriteString(cmd)
	sb.WriteString("\n")
	res, err := session.Output(cmd)
	if err != nil {
		sb.WriteString("...Run failed: ")
		sb.WriteString(err.Error())
	} else {
		sb.WriteString("...Run successfully.\n\n")
		sb.WriteString(string(res))
	}
}
func Incognito() {
	sb := strings.Builder{}
	sb.WriteString("run Incognito...\n")
	pm := PrivacyMode{}
	//JsonRequestPost(PrivacyModeUrl, pm, &sb)
	PrivacySSH(pm, &sb)
	Outputlog.SetText(sb.String())

}

func Tracking() {
	sb := strings.Builder{}
	sb.WriteString("run Tracking...\n")
	pm := PrivacyMode{
		PersonalGps:        true,
		PersonalUser:       true,
		PseudonymousGps:    true,
		PersonalVehicle:    true,
		AnonymousDataNoGPS: true,
	}
	//JsonRequestPost(PrivacyModeUrl, pm, &sb)

	PrivacySSH(pm, &sb)
	Outputlog.SetText(sb.String())
}

func Location() {
	sb := strings.Builder{}
	sb.WriteString("run Location...\n")
	pm := PrivacyMode{
		PersonalUser:       true,
		PseudonymousGps:    true,
		PersonalVehicle:    true,
		AnonymousDataNoGPS: true,
	}
	//JsonRequestPost(PrivacyModeUrl, pm, &sb)

	PrivacySSH(pm, &sb)
	Outputlog.SetText(sb.String())
}

func Personal() {
	sb := strings.Builder{}
	sb.WriteString("run Personal...\n")
	pm := PrivacyMode{
		PersonalUser:       true,
		PersonalVehicle:    true,
		AnonymousDataNoGPS: true,
	}
	//JsonRequestPost(PrivacyModeUrl, pm, &sb)

	PrivacySSH(pm, &sb)
	Outputlog.SetText(sb.String())
}
