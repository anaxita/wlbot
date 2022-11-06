package authentificator

type Service struct {
	adminChats map[int64]struct{}
	adminUsers map[string]struct{}
}

func New(adminChats map[int64]struct{}, adminUsers map[string]struct{}) *Service {
	return &Service{
		adminChats: adminChats,
		adminUsers: adminUsers,
	}
}
func (s *Service) IsAdmin(chatID int64, username string) bool {
	return s.IsAdminChat(chatID) || s.IsAdminUser(username)
}

func (s *Service) IsAdminChat(chatID int64) bool {
	_, ok := s.adminChats[chatID]
	return ok
}

func (s *Service) IsAdminUser(username string) bool {
	_, ok := s.adminUsers[username]
	return ok
}
