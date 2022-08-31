package putty

//func JsonRequestPost(url string, pm interface{}, sb *strings.Builder) {
//	jsondata, err := json.Marshal(pm)
//	if err != nil {
//		sb.WriteString(err.Error())
//		return
//	}
//	reader := bytes.NewReader(jsondata)
//	sb.WriteString(string(jsondata))
//	sb.WriteString("\n")
//	request, err := http.NewRequest("POST", url, reader)
//	if err != nil {
//		sb.WriteString(err.Error())
//		sb.WriteString("\n")
//		return
//	}
//	request.Header.Set("Content-Type", "application/json;charset=utf-8")
//	client := http.Client{}
//	resp, err := client.Do(request)
//	if err != nil {
//		sb.WriteString(err.Error())
//		sb.WriteString("\n")
//		return
//	}
//	respBytes, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		sb.WriteString(err.Error())
//		sb.WriteString("\n")
//		return
//	}
//	sb.WriteString(string(respBytes))
//	sb.WriteString("\n")
//	sb.WriteString("run success\n")
//}
//
//func RequestGet(url string, sb *strings.Builder) {
//	resp, err := http.Get(url)
//	if err != nil {
//		Outputlog.SetText(err.Error())
//		return
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	sb.WriteString(string(body))
//	sb.WriteString("\n")
//	sb.WriteString("run success\n")
//}
