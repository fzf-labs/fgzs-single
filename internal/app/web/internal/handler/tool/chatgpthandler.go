package tool

import (
	"fgzs-single/internal/app/web/internal/svc"
	"net/http"
	"time"

	"github.com/fzf-labs/fpkg/cache/collectioncache"
	"github.com/fzf-labs/fpkg/third_api/openai"
	"github.com/fzf-labs/fpkg/util/uuidutil"
	"github.com/gorilla/websocket"
)

func ChatGPTHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrade, err := ChatGPTWsUpgrade(w, r)
		if err != nil {
			return
		}
		chatGPT := openai.NewChatGPT(svcCtx.Config.OpenAI.ChatGPT)
		ChatGPTMessageHandle(upgrade, chatGPT, svcCtx.CollectionCacheChatGpt)
	}
}

// ChatGPTWsUpgrade websocket 连接
func ChatGPTWsUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
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

// ChatGPTMessageHandle 发送消息
func ChatGPTMessageHandle(conn *websocket.Conn, gpt *openai.ChatGPT, chatGptCache *collectioncache.Cache) {
	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
	}()
	clientId := uuidutil.KSUidByTime()
	for {
		_, req, err := conn.ReadMessage()
		if err != nil {
			return
		}
		messages := make([]openai.ChatMessage, 0)
		var resp []byte
		oldMessage, ok := chatGptCache.Get(clientId)
		if ok {
			chatMessages, ok := oldMessage.([]openai.ChatMessage)
			if ok {
				messages = append(messages, chatMessages...)
			}
		}
		messages = append(messages, openai.ChatMessage{
			Role:    "user",
			Content: string(req),
		})
		completions, err := gpt.ChatCompletions(messages)
		if err != nil {
			return
		}
		chatGptCache.Set(clientId, messages)
		resp = []byte(completions.Content)
		err = conn.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			return
		}
	}
}
