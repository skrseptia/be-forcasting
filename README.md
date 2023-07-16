# Readme

Build binary file using this cmd:
```console
GOOS=linux GOARCH=amd64 go build -o andhiga-pupuk
```

Save file to the server
```console
scp andhiga-pupuk root@192.168.0.237:/root/bin/andhiga-pupuk
```

Restart Supervisor
```console
ssh root@192.168.0.237 -C ./andhiga-pupuk.sh
```
