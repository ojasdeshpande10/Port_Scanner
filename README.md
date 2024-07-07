# CSE508-Assignment 4

Creating a synprobe

## Instructions to run the server and client side proxies


Build the synprobe program:

go build synprobe.go

Server Command:

./synprobe -p [port range/port number] host address


Example output:

┌──(kali㉿kali)-[~/NS_A2]
└─$ ./synprobe -p 21 ftp.dlptest.com
Response from port 21 TCP server [server-initiated] : 220 Welcome to the DLP Test FTP Server


                                                                                                                   
┌──(kali㉿kali)-[~/NS_A2]
└─$ ./synprobe -p 130 compute.cs.stonybrook.edu
Response from port 130 TCP server [server-initiated] : SSH-2.0-OpenSSH_7.4


                                                                                                                                                                                                                                            
┌──(kali㉿kali)-[~/NS_A2]
└─$ ./synprobe -p 993 imap.gmail.com           
Response from port 993 TLS server [server-initiated] : * OK Gimap ready for requests from 129.49.252.167 u14mb146278123qta


                                                                                                                                                                                                                                            
┌──(kali㉿kali)-[~/NS_A2]
└─$ ./synprobe -p 465 smtp.gmail.com
Response from port 465 TLS server [server-initiated] : 220 smtp.gmail.com ESMTP g16-20020a0cf850000000b006a0d781b105sm3427171qvo.93 - gsmtp


                                                                                                                                                                                                                                            
┌──(kali㉿kali)-[~/NS_A2]
└─$ ./synprobe -p 443 www.cs.stonybrook.edu
Response from port 443 TLS HTTPS server [client-initiated]: HTTP/1.1 404 Unknown site
Connection: close
Content-Length: 566
Retry-After: 0
Server: Pantheon
Cache-Control: no-cache, must-revalidate
Content-Type: text/html; charset=utf-8
X-pantheon-serious-reason: The page could not be loaded properly.
Date: Mon, 06 May 2024 04:06:13 GMT
X-Served-By: cache-lga21944-LGA
X-Cache: MISS
X-Cache-Hits: 0
X-Timer: S1714968373.390788,VS0,VE36
Vary: Cookie
Age: 0
Accept-Ranges: bytes
Via: 1.1 varnish



──(kali㉿kali)-[~/NS_A2]
└─$ ./synprobe  www.cs.stonybrook.edu 
Response from port 443 TLS HTTPS server [client-initiated]: HTTP/1.1 404 Unknown site
Connection: close
Content-Length: 566
Retry-After: 0
Server: Pantheon
Cache-Control: no-cache, must-revalidate
Content-Type: text/html; charset=utf-8
X-pantheon-serious-reason: The page could not be loaded properly.
Date: Mon, 06 May 2024 04:09:48 GMT
X-Served-By: cache-lga21979-LGA
X-Cache: MISS
X-Cache-Hits: 0
X-Timer: S1714968588.226715,VS0,VE31
Vary: Cookie
Age: 0
Accept-Ranges: bytes
Via: 1.1 varnish



Response from port 80 TCP HTTP server [client-initiated]: HTTP/1.1 404 Unknown site
Connection: close
Content-Length: 566
Retry-After: 0
Server: Pantheon
Cache-Control: no-cache, must-revalidate
Content-Type: text/html; charset=utf-8
X-pantheon-serious-reason: The page could not be loaded properly.
Date: Mon, 06 May 2024 04:09:48 GMT
X-Served-By: cache-lga21923-LGA
X-Cache: MISS
X-Cache-Hits: 0
X-Timer: S1714968588.227837,VS0,VE66
Vary: Cookie
Age: 0
Accept-Ranges: bytes
Via: 1.1 varnish

<!DOCTYPE HTML>
      <html>
        <head>
          <title>404 - Unknown site</title>
        </head>
        <body style="font-family:Arial, Helvetica, sans-serif; text-align: center">
          <div style='padding-block: 180px'>
            <h1>
              <div style='font-size: 180px; font-weight: 700'>404</div>
              <div style='font-size: 24px; font-weight: 700'>Unknown site</div>
            </h1>
            <p style="font-size: 16px; font-weight: 400">The page could not be loaded properly.</p>
          </div>
        </body>
      </html>

Port 25 closed (TCP)

Port 853 closed (TCP)

Port 3389 closed (TCP)

Port 110 closed (TCP)

Port 587 closed (TCP)

Port 143 closed (TCP)

Port 21 closed (TCP)

Port 23 closed (TCP)

Port 8080 closed (TCP)

Port 22 closed (TCP)

Port 993 closed (TCP)


### Prerequisites


Tested on Kali Linux

Go should be installed >1.16
