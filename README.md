# APITable
An API to manage iptables rules and ipsets.

## Features  
- Create and Delete iptable rules.
- Create and Delete ipset groups.
- Add and remove members from ipset group.
- Swagger documentation for the API.

## Requirements
- go 1.23<
- iptable and ipset should installed.

## Installation  
Clone the repository:  
   ```bash
   git clone https://github.com/miladnami97/APIPtable.git
   cd APIPtable
   go mod tidy
   go build -o apitable
   sudo ./apitable
```
## Usage

You can access the API on `http://localhost:8089`

And the API documents on `http://localhost:8089/docs/index.html`

Or try
create ipset group (name=TEST, type=hash:ip):
```bash
curl -X POST http://localhost:8089/firewall/ipset -H "Content-Type: application/json" -d '{"group_name":"TEST"}'
```

remove ipset group (name=TEST, type=hash:ip): 
```bash
curl -X DELETE http://localhost:8089/firewall/ipset -H "Content-Type: application/json" -d '{"group_name":"TEST"}'`
```

add ip to a ipset:
```bash
curl -X POST http://localhost:8089/firewall/ipset/ip -H "Content-Type: application/json" -d '{"group_name":"TEST","ip":"192.168.1.105"}'`
```

remove ip to a ipset:
```bash
curl -X DELETE http://localhost:8089/firewall/ipset/ip -H "Content-Type: application/json" -d '{"group_name":"TEST","ip":"192.168.1.105"}'`
```

add iptable rule:
```bash
curl -X POST http://127.0.0.1:8089/firewall/rules -H "Content-Type: application/json" -d '{"ipset_name":"TEST","dst_port":"5000","action":"DROP"}'`
```

remove iptable rule:
```bash
curl -X DELETE http://127.0.0.1:8089/firewall/rules -H "Content-Type: application/json" -d '{"ipset_name":"TEST","dst_port":"5000","action":"DROP"}'`
```
