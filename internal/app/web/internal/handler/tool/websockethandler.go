package tool

import (
	"fgzs-single/internal/app/web/internal/svc"
	"github.com/fzf-labs/fpkg/third_api/openai"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"time"
)

func WebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrade, err := WsUpgrade(w, r)
		if err != nil {
			return
		}
		chatGPT := openai.NewChatGPT(svcCtx.Config.OpenAI.ChatGPT)
		MessageHandle(upgrade, chatGPT)
	}
}

// WsUpgrade websocket 连接
func WsUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws := websocket.Upgrader{
		HandshakeTimeout: time.Minute * 5,
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

// MessageHandle 发送消息
func MessageHandle(conn *websocket.Conn, gpt *openai.ChatGPT) {
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
		if strings.ToLower(string(req)) == "ping" {
			resp = []byte("pong")
		} else {
			completions, err := gpt.Completions(string(req))
			if err != nil {
				return
			}
			resp = []byte(completions)
		}
		err = conn.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			return
		}
	}
}
