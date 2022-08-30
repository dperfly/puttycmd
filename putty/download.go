package putty

import (
	"fmt"
	"puttycmd/constant"
	"strings"
)

func OCU() {

	sb := strings.Builder{}
	sb.WriteString("run GetOCNLogFunc\n")
	session, err := SSH(host, port, username, password)
	fmt.Println(host, port, username, password)
	if err != nil {
		sb.WriteString(err.Error())
		return
	}
	//var DownloadOCUUrl = fmt.Sprintf("http://%s:%d/%s/", constant.Domain, constant.OCUPost, constant.ServicemanagementServices)
	//RequestGet(DownloadOCUUrl, &sb)
	var DownloadOCUUrl = fmt.Sprintf("cd /mnt ; ./curl -L http://%s:%d/%s/", constant.Domain, constant.OCUPost, constant.ServicemanagementServices)
	res, err := session.Output(DownloadOCUUrl)

	if err != nil {
		sb.WriteString(err.Error())
	} else {
		sb.WriteString(string(res))
	}
	Output.SetText(sb.String())
}

func CNS() {

	sb := strings.Builder{}
	sb.WriteString("run GetOCNLogFunc\n")
	session, err := SSH(host, port, username, password)
	if err != nil {
		sb.WriteString(err.Error())
		return
	}
	//var DownloadOCUUrl = fmt.Sprintf("http://%s:%d/%s/", constant.Domain, constant.CNDPost, constant.ServicemanagementServices)
	//RequestGet(DownloadOCUUrl, &sb)
	var DownloadOCUUrl = fmt.Sprintf("cd /mnt ; ./curl -L http://%s:%d/%s/", constant.Domain, constant.CNDPost, constant.ServicemanagementServices)
	res, err := session.Output(DownloadOCUUrl)

	if err != nil {
		sb.WriteString(err.Error())
	} else {
		sb.WriteString(string(res))
	}
	Output.SetText(sb.String())
}
