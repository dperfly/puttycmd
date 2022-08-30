package putty

import (
	"encoding/json"
	"fmt"
	"log"
	"puttycmd/constant"
	"strconv"
	"strings"
)

var PrivacyModeUrl = fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.PrivacysetupAddress, constant.Privacy)

func PrivacySSH(pm PrivacyMode, sb *strings.Builder) ([]byte, error) {
	jsondata, err := json.Marshal(pm)
	if err != nil {
		sb.WriteString(err.Error())
		return nil, err
	}
	strJson := strconv.Quote(string(jsondata))
	session, err := SSH(host, port, username, password)
	fmt.Println(host, port, username, password)
	if err != nil {
		sb.WriteString(err.Error())
		return nil, err
	}
	cmd := fmt.Sprintf("cd /mnt; ./curl -H \"Content-Type: application/json\" -X POST %s -d %s", PrivacyModeUrl, strJson)
	log.Println(cmd)
	return session.Output(cmd)
}
func Incognito() {
	sb := strings.Builder{}
	sb.WriteString("run Incognito\n")
	pm := PrivacyMode{}
	//JsonRequestPost(PrivacyModeUrl, pm, &sb)

	res, err := PrivacySSH(pm, &sb)
	if err != nil {
		sb.WriteString(err.Error())
	} else {
		sb.WriteString(string(res))
	}

	Output.SetText(sb.String())

}

func Tracking() {
	sb := strings.Builder{}
	sb.WriteString("run Tracking\n")
	pm := PrivacyMode{
		PersonalGps:        true,
		PersonalUser:       true,
		PseudonymousGps:    true,
		PersonalVehicle:    true,
		AnonymousDataNoGPS: true,
	}
	//JsonRequestPost(PrivacyModeUrl, pm, &sb)

	res, err := PrivacySSH(pm, &sb)
	if err != nil {
		sb.WriteString(err.Error())
	} else {
		sb.WriteString(string(res))
	}

	Output.SetText(sb.String())
}

func Location() {
	sb := strings.Builder{}
	sb.WriteString("run Location\n")
	pm := PrivacyMode{
		PersonalUser:       true,
		PseudonymousGps:    true,
		PersonalVehicle:    true,
		AnonymousDataNoGPS: true,
	}
	//JsonRequestPost(PrivacyModeUrl, pm, &sb)

	res, err := PrivacySSH(pm, &sb)
	if err != nil {
		sb.WriteString(err.Error())
	} else {
		sb.WriteString(string(res))
	}

	Output.SetText(sb.String())
}

func Personal() {
	sb := strings.Builder{}
	sb.WriteString("run Location\n")
	pm := PrivacyMode{
		PersonalUser:       true,
		PersonalVehicle:    true,
		AnonymousDataNoGPS: true,
	}
	//JsonRequestPost(PrivacyModeUrl, pm, &sb)

	res, err := PrivacySSH(pm, &sb)
	if err != nil {
		sb.WriteString(err.Error())
	} else {
		sb.WriteString(string(res))
	}

	Output.SetText(sb.String())
}
