package member

type MemberUseCase struct {
}

func NewMemberService() *MemberUseCase {
	return &MemberUseCase{}
}

func (s *MemberUseCase) Register() string {
	return "註冊會員"
}
