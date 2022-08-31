package putty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"puttycmd/constant"
	"strings"
)

func DownloadLog(DownloadUrl string, sb *strings.Builder) {
	session, err := SSH(host, port, username, password)
	if err != nil {
		sb.WriteString(err.Error())
		return
	}
	//var DownloadOCUUrl = fmt.Sprintf("http://%s:%d/%s/", constant.Domain, constant.OCUPost, constant.ServicemanagementServices)
	//RequestGet(DownloadOCUUrl, &sb)
	res, err := session.Output(DownloadUrl)
	sb.WriteString(DownloadUrl)
	sb.WriteString("\n")
	var str bytes.Buffer
	if err != nil {
		sb.WriteString("...Run failed: ")
		sb.WriteString(err.Error())
	} else {
		sb.WriteString("...Run successfully\n\n")
	}
	if err := json.Indent(&str, res, "", "    "); err == nil {
		sb.WriteString(str.String())
	} else {
		sb.Write(res)
	}
}
func OCU() {
	sb := strings.Builder{}
	sb.WriteString("run GetOCNLogFunc...\n")
	var DownloadOCUUrl = fmt.Sprintf("cd /mnt ; ./curl -L http://%s:%d/%s/", constant.Domain, constant.OCUPost, constant.ServicemanagementServices)
	DownloadLog(DownloadOCUUrl, &sb)
	Outputlog.SetText(sb.String())

}

func CNS() {

	sb := strings.Builder{}
	sb.WriteString("run GetCNSLogFunc...\n")
	var DownloadCNSUrl = fmt.Sprintf("cd /mnt ; ./curl -L http://%s:%d/%s/", constant.Domain, constant.CNSPost, constant.ServicemanagementServices)
	DownloadLog(DownloadCNSUrl, &sb)
	Outputlog.SetText(sb.String())
}
