import json
from pydoc import cli
import socket

req = {
    "id": 0,
    "params": ["luxcgo"],
    "method": "HelloService.Hello"
}

client = socket.create_connection(("localhost", 1234))
client.sendall(json.dumps(req).encode())

resp = client.recv(1024)
resp = json.loads(resp.decode())
print(resp)
