package ipc

import "testing"

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{method, params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, err1 := client1.Call("get", "param1")
	resp2, err2 := client2.Call("post", "param2")

	if err1 != nil || err2 != nil {
		t.Error("Client Call function failed.")
	}
	if resp1.Body != "param1" || resp1.Code != "get" || resp2.Body != "param2" || resp2.Code != "post" {
		t.Error("IpcClient.Call failed. resp1:", resp1, " resp2:", resp2)
	}
	client1.Close()
	client2.Close()
}
