package tool

import (
	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/pkg/openai"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func ChatGPTHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrade, err := ChatGPTWsUpgrade(w, r)
		if err != nil {
			return
		}
		chatGPT := openai.NewChatGPT(svcCtx.Config.OpenAI.ChatGPT)
		ChatGPTMessageHandle(upgrade, chatGPT)
	}
}

// ChatGPTWsUpgrade websocket 连接
func ChatGPTWsUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws := websocket.Upgrader{
		HandshakeTimeout: time.Second * 5,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// ChatGPTMessageHandle 发送消息
func ChatGPTMessageHandle(conn *websocket.Conn, gpt *openai.ChatGPT) {
	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
	}()
	for {
		_, req, err := conn.ReadMessage()
		if err != nil {
			return
		}
		var resp []byte
		completions, err := gpt.Completions(string(req))
		if err != nil {
			return
		}
		resp = []byte(completions)
		err = conn.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			return
		}
	}
}
