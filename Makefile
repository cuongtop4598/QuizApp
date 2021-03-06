KillProcess:
	sudo kill -9 `sudo lsof -t -i:8000`
CreateCA:
	openssl req -new -key server.key -out server.pem\n
	openssl req -newkey rsa:4096 -nodes -keyout server.key -out server.csr\n
	openssl x509 -signkey server.key -in server.csr -req -days 365 -out server.crt