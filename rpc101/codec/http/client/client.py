import requests

req = {
    "id": 0,
    "params": ["luxcgo"],
    "method": "HelloService.Hello"
}

resp = requests.post("http://localhost:1234/jsonrpc", json=req)
print(resp.text)
