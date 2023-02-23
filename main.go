package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("domain hasMX hasSPF SPFresord hasDMARC dmarcRecord\n")

	for scanner.Scan(){
		checkDomain(scanner.Text())  // Scanner.Text is value that user type like mailchimp.com we can check multiple emails at the same time thats why for loop 

	}
}


func checkDomain(domain string){

	var hasDMARC , hasMX , hasSPF bool
	var SPFresord , dmarcRecord string

	mxRecord,err := net.LookupMX(domain)

	if err!=nil{
		log.Printf("Error %v",err)
	}

	if len(mxRecord)>0{
		hasMX=true
	}

	txtRecords,err := net.LookupTXT(domain)
	if err!=nil{
		log.Printf("Error %v",err)
	}

	for _ ,record:= range txtRecords{

		if strings.HasPrefix(record,"v=spf1"){
			hasSPF=true
			SPFresord=record
			break
		}

	}

	dmarcRecords,err:=net.LookupTXT("_dmarc."+domain)
	if err!=nil{
		log.Printf("Error %v",err)

	}

	for _,record :=range dmarcRecords{

		if strings.HasPrefix(record,"v=DMARC1"){
			hasDMARC=true
			dmarcRecord=record
			break
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v",domain,hasMX,hasDMARC,dmarcRecord,hasSPF,SPFresord)



}