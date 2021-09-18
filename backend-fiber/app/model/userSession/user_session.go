package userSession

type Users_Sessions struct {
	Users_sessions_id uint32
	Fk_user_id        uint32
	SESSION_ID        string
}

type Users_Sessions_Ins struct {
	Fk_user_id uint32
	SESSION_ID string
}
