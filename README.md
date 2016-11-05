# unset-leap-indicator-proxy
unset-leap-indicator-proxy is UDP proxy that unsets leap indicator flag of a NTP packets.

# overview
UDPパケットをプロキシします。NTPパケットをパースはしませんが、ただ上位2bitを上書きし Leap Indicatorを0に上書きします。

パブリックネットワークで使用する場合は、踏み台攻撃に注意してください

```
|Client|  -- NTP Request ->  |Proxy| -- NTP Request -> |Server|
(Proxy does'n parse NTP messages)

|Client|  <- NTP Request(LI:0) --  |Proxy| <- NTP Response(LI:1) -- |Server|
(Proxy just rewrites top top 2bits of udp body(== LI))
```

- 1) The proxy recieves a ntp request from client.
- 2) The proxy sends the recieved data to ntp server (proxy does'n perse ntp messages).
- 3) The proxy recieves a ntp response from server, and just rewrites top 2bits of udp body(== LI) to 0.
- 4) The proxy sends the recieved data to client.

if use in public network, You should be careful a springboard attack.


# usage
```
go run ./proxy.go -v ntp.example.com
```
- specify ntp server
- -v: show a connecting client address

example
```
vagrant@vagrant:~$ sudo go run ./proxy.go -v ntp.example.com
Info: Start Proxy on  :123  to ntp.example.com
Received from 192.168.0.179:123
Received from 192.168.0.179:123
```