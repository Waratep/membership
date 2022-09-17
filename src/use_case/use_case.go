package use_case

import (
	"github.com/Waratep/membership/src/entity/member"
	"github.com/gin-gonic/gin"
)

type UseCase struct {
	memberRepository MemberRepository
}

type MemberRepository interface {
	HealthCheck(ctx gin.Context) error
	CreateMember(ctx gin.Context, m member.Member) (MembershipMember, error)
	GetMembers(ctx gin.Context) error
	GetMemberByID(ctx gin.Context) error
	UpdateMember(ctx gin.Context) error
	DeleteMember(ctx gin.Context) error
	GetMemberByPhone(ctx gin.Context, phone string) (MembershipMember, error)
	GetMemberByEmail(ctx gin.Context, email string) (MembershipMember, error)
}

func New(member MemberRepository) UseCase {
	return UseCase{
		memberRepository: member,
	}
}
