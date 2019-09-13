package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	xmlString := `<?xml version="1.0" encoding="utf-8"?>
<servers version="1">
	<server>
		<serverName>Shanghai_VPN</serverName>
		<serverIP>127.0.0.1</serverIP>
	</server>
	<server>
		<serverName>Beijing_VPN</serverName>
		<serverIP>127.0.0.2</serverIP>
	</server>
</servers>`

	var dat *Recurlyservers
	err := xml.Unmarshal([]byte(xmlString), &dat)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dat.Description)
	fmt.Println(dat.Svs)
	fmt.Println(dat.Version)
	fmt.Println(dat.XMLName)

}
