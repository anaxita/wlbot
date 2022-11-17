package rest

import (
	"encoding/json"
	"net/http"

	"kms/wlbot/internal/helpers"

	"go.uber.org/zap"
)

func (s *Server) SendHandler() http.Handler {
	type ReqData struct {
		Text string `json:"text"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req ReqData

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Text != "" {
			err = s.notificator.SendToAdminChats(req.Text)
			if err != nil {
				s.l.Error("failed to send message to admin chats", zap.Error(err))
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func (s *Server) AddIPHandler() http.Handler {
	type ReqData struct {
		IP4      string `json:"ip4"`
		UserName string `json:"user_name"`
		Comment  string `json:"comment"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req ReqData

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.mikrotik.AddIPToDefaultMikrotiks(
			r.Context(),
			req.IP4,
			helpers.TranslitRuToEN(req.UserName+" | "+req.Comment),
		)
		if err != nil {
			s.l.Error("failed to add ip to mikrotik", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
