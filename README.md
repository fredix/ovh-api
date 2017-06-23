# ovh-api
golang scripts using ovh's IP API. I use move_ip.go with keepalived

first you need to create your token's api  as explain on https://api.ovh.com/g934.first_step_with_api
go to https://eu.api.ovh.com/createApp/
add your application key (AK) and your application secret (AS) to ovh.conf-sample and move it to /etc/ovh.con
last, you need to launch register_ip.go to get a consumer Key (CS) authorized to use ip api.

add your CS to /etc/ovh.conf, now you can use get_service.go and move_ip.go


go run get_service.go -h

Usage of ./get_service:
  -endpoint string
    	example ovh-eu (default "ovh-eu")
  -ip string
    	your ip failover (default "xx.xx.xx.xx")


get_service return the service name (aka your VPS) associate to your ip failover, ie :

serviceName :  {fr {vpsxxxx.ovh.net}}

fr is the country code of your ip, vpsxxx.ovh.net the service name (hostname) of your vps. 

move_ip -h               
Usage of move_ip:
  -endpoint string
    	example ovh-eu (default "ovh-eu")
  -ip string
    	your ip failover (default "xx.xx.xx.xx")

move_ip move the ip failover on the host running move_ip.
