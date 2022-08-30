package putty

import (
	"encoding/json"
	"fmt"
	"log"
	"puttycmd/constant"
	"strconv"
	"strings"
)

func ServicesSSH(enableStatusReportUrl string, disableMode DisableMode, sb *strings.Builder) {
	session, err := GetSession()
	if err != nil {
		sb.WriteString(err.Error())
		return
	}

	jsondata, err := json.Marshal(disableMode)
	if err != nil {
		sb.WriteString(err.Error())
		return
	}
	strJson := strconv.Quote(string(jsondata))
	cmd := fmt.Sprintf("cd /mnt; ./curl -H \"Content-Type: application/json\" %s -d %s", enableStatusReportUrl, strJson)
	log.Println(cmd)
	res, err := session.Output(cmd)
	if err != nil {
		sb.WriteString(err.Error())
	} else {
		sb.WriteString(string(res))
	}

}
func EnableStatusReport() {
	sb := strings.Builder{}
	sb.WriteString("run EnableStatusReport\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Statusreport)
	disableMode := DisableMode{DisableReasons: []string{}}

	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	Output.SetText(sb.String())
}
func DisableStatusReport() {
	sb := strings.Builder{}
	sb.WriteString("run DisableStatusReport\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Statusreport)
	disableMode := DisableMode{DisableReasons: []string{LocalVehicleDisabled}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Output.SetText(sb.String())
}

func EnableVehicleHealth() {
	sb := strings.Builder{}
	sb.WriteString("run EnableVehicleHealth\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Vehiclehealth)
	disableMode := DisableMode{DisableReasons: []string{}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Output.SetText(sb.String())

}

func DisableVehicleHealth() {
	sb := strings.Builder{}
	sb.WriteString("run DisableVehicleHealth\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Vehiclehealth)
	disableMode := DisableMode{DisableReasons: []string{LocalVehicleDisabled}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)

	Output.SetText(sb.String())
}

func EnableCarFinder() {

	sb := strings.Builder{}
	sb.WriteString("run EnableCarFinder\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Carfinder)
	disableMode := DisableMode{DisableReasons: []string{}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)

	Output.SetText(sb.String())
}

func DisableCarFinder() {

	sb := strings.Builder{}
	sb.WriteString("run DisableCarFinder\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Carfinder)
	disableMode := DisableMode{DisableReasons: []string{LocalVehicleDisabled}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Output.SetText(sb.String())
}

func EnableRhonk() {

	sb := strings.Builder{}
	sb.WriteString("run EnableRhonk\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Rhonk)
	disableMode := DisableMode{DisableReasons: []string{}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Output.SetText(sb.String())
}
func DisableRhonk() {
	sb := strings.Builder{}
	sb.WriteString("run DisableRhonk\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Rhonk)
	disableMode := DisableMode{DisableReasons: []string{LocalVehicleDisabled}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Output.SetText(sb.String())
}
