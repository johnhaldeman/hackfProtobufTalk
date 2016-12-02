    package main

    import (
        "log"
        "net"
        "bufio"
        "os"
	"fmt"
        "strings"
        "github.com/golang/protobuf/proto"
        "github.com/johnhaldeman/httpmonitorproto"
	"strconv"
        "encoding/binary"
    )
	
	func check(e error) {
		if e != nil {
			panic(e)
		}
	}

    func main() {
		conn, err := net.Dial("tcp", "10.10.9.1:8686")
		check(err)
		defer conn.Close()	
		
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan(){
			splitLogLine := strings.Split(scanner.Text(), "||");
			timeParse, err := strconv.ParseInt(splitLogLine[0], 10, 64)
			check(err)
			httpCodeParse, err := strconv.ParseInt(splitLogLine[10], 10, 32)
			check(err)
			contentSizeParse, err := strconv.ParseInt(splitLogLine[12], 10, 64)
			check(err)
		
			test := &monitor_http.HttpRequest {
				Timestamp: proto.Int64(timeParse),
				Hostname:  proto.String(splitLogLine[1]),
				ServerName: proto.String(splitLogLine[2]),
				ServerIp: proto.String(splitLogLine[3]),
				Protocol: proto.String(splitLogLine[4]),
				HttpUser: proto.String(splitLogLine[5]),
				Method: proto.String(splitLogLine[6]),
				Resource: proto.String(splitLogLine[7]),
				Query: proto.String(splitLogLine[8]),
				FullRequest: proto.String(splitLogLine[9]),
				HttpCode: proto.Int32( int32(httpCodeParse) ),
				ConnStatus: proto.String(splitLogLine[11]),
				ContentSize: proto.Int64(contentSizeParse),
				TimeToServe: proto.String(splitLogLine[13]),
				HeaderReferer: proto.String(splitLogLine[14]),
				HeaderUserAgent: proto.String(splitLogLine[15]),
				HeaderAccept: proto.String(splitLogLine[16]),
				HeaderAcceptLanguage: proto.String(splitLogLine[17]),
				File: proto.String(splitLogLine[18]),
				
			}
			data, err := proto.Marshal(test)
			if err != nil {
				log.Fatal("marshaling error: ", err)
			}
			
			dataLength := int32(len(data))
			fmt.Printf("Writing what I think is: %d bytes", dataLength);
                        binary.Write(conn, binary.BigEndian, dataLength)
			conn.Write(data)
		}
    }

	
	
