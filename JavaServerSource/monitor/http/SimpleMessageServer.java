package monitor.http;
import java.net.*;
import java.time.Instant;
import java.util.Iterator;
import java.util.Map;

import com.google.protobuf.Descriptors.FieldDescriptor;

import monitor.http.Monitorhttp.http_request;

import java.io.*;

public class SimpleMessageServer {
		
	public static void main(String[] args) throws IOException {
         
        int portNumber = 8686;

        ServerSocket serverSocket = new ServerSocket(portNumber);
        Socket clientSocket = serverSocket.accept();
        DataInputStream inStream = new DataInputStream(clientSocket.getInputStream());
        
        while(true){
        	System.out.println("Accepted Connection");
        	System.out.println("Waiting for Data");
        	int ContentSize = inStream.readInt();
        	if(ContentSize != 0){
        		System.out.println("--------------------------------------");
        		System.out.println("Wants to send " + ContentSize + " bytes of data");
            	System.out.println("Reading...");
            	
            	byte[] bytes = new byte[ContentSize];
            	inStream.read(bytes, 0, ContentSize);
            	http_request requestProto = http_request.parseFrom(bytes);

            	System.out.println("Time of Request (milliseconds since epoch): " + requestProto.getTimestamp());
            	Instant inst = Instant.ofEpochMilli(requestProto.getTimestamp());
            	System.out.println("Time of Request (human readable): " + inst.toString());
            	System.out.println("This was the request string: " + requestProto.getFullRequest());
            	System.out.println();
            	
            	Map<FieldDescriptor, Object> httpReqMap = requestProto.getAllFields();
            	Iterator<FieldDescriptor> keys = httpReqMap.keySet().iterator();
            	while(keys.hasNext()){
            		FieldDescriptor protoField = keys.next();
            		System.out.print(protoField.getName() + ": ");
            		System.out.println(httpReqMap.get(protoField));
            	}
        	}
      	
        }
        
    }
}
