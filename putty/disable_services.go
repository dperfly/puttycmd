package putty

import (
	"encoding/json"
	"fmt"
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
	sb.WriteString(cmd)
	sb.WriteString("\n")
	res, err := session.Output(cmd)
	if err != nil {
		sb.WriteString("...Run failed: ")
		sb.WriteString(err.Error())
	} else {
		sb.WriteString("...Run successfully\n\n")
		sb.WriteString(string(res))
	}

}
func RecoverStatusReport() {
	sb := strings.Builder{}
	sb.WriteString("run RecoverStatusReport...\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Statusreport)
	disableMode := DisableMode{DisableReasons: []string{}}

	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	Outputlog.SetText(sb.String())
}
func DisableStatusReport() {
	sb := strings.Builder{}
	sb.WriteString("run DisableStatusReport...\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Statusreport)
	disableMode := DisableMode{DisableReasons: []string{LocalVehicleDisabled}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Outputlog.SetText(sb.String())
}

func RecoverVehicleHealth() {
	sb := strings.Builder{}
	sb.WriteString("run RecoverVehicleHealth...\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Vehiclehealth)
	disableMode := DisableMode{DisableReasons: []string{}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Outputlog.SetText(sb.String())

}

func DisableVehicleHealth() {
	sb := strings.Builder{}
	sb.WriteString("run DisableVehicleHealth...\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Vehiclehealth)
	disableMode := DisableMode{DisableReasons: []string{LocalVehicleDisabled}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)

	Outputlog.SetText(sb.String())
}

func RecoverCarFinder() {

	sb := strings.Builder{}
	sb.WriteString("run RecoverCarFinder...\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Carfinder)
	disableMode := DisableMode{DisableReasons: []string{}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)

	Outputlog.SetText(sb.String())
}

func DisableCarFinder() {

	sb := strings.Builder{}
	sb.WriteString("run DisableCarFinder...\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Carfinder)
	disableMode := DisableMode{DisableReasons: []string{LocalVehicleDisabled}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Outputlog.SetText(sb.String())
}

func RecoverRhonk() {

	sb := strings.Builder{}
	sb.WriteString("run RecoverRhonk...\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Rhonk)
	disableMode := DisableMode{DisableReasons: []string{}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Outputlog.SetText(sb.String())
}
func DisableRhonk() {
	sb := strings.Builder{}
	sb.WriteString("run DisableRhonk...\n")
	enableStatusReportUrl := fmt.Sprintf("http://%s:%d/%s/%s", constant.Domain, constant.OCUPost, constant.ServicemanagementServices, constant.Rhonk)
	disableMode := DisableMode{DisableReasons: []string{LocalVehicleDisabled}}
	//JsonRequestPost(enableStatusReportUrl, disableMode, &sb)
	ServicesSSH(enableStatusReportUrl, disableMode, &sb)
	Outputlog.SetText(sb.String())
}
