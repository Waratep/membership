package use_case

import (
	"errors"
	"log"

	"github.com/Waratep/membership/src/entity/member"
	"github.com/gin-gonic/gin"
)

var (
	ErrorFirstNameIsRequire = errors.New("first_name is require")
	ErrorLastNameIsRequire  = errors.New("last_name is require")
	ErrorPhoneIsRequire     = errors.New("phone is require")
	ErrorEmailIsRequire     = errors.New("email is require")
	ErrorAddressIsRequire   = errors.New("address is require")
	ErrorDuplicatePhone     = errors.New("Error duplicate phone number")
	ErrorDuplicateEmail     = errors.New("Error duplicate email")
	ErrorItemNotFound       = errors.New("Error item not found")
	ErrorDataTransform      = errors.New("error while transforming data")
)

type MembershipMember struct {
	Member member.Member
	ID     int64
}

func (u UseCase) CreateMember(ctx gin.Context, m member.Member) (int64, error) {
	validateRequireFields(m)

	memberByPhone, err := u.memberRepository.GetMemberByPhone(ctx, m.Phone)
	if err != nil && err != ErrorItemNotFound {
		log.Println("Error get member by phone number")

		return 0, err
	}
	if memberByPhone.Member.Phone != "" {
		return 0, ErrorDuplicatePhone
	}

	memberByEmail, err := u.memberRepository.GetMemberByEmail(ctx, m.Email)
	if err != nil && err != ErrorItemNotFound {
		log.Println("Error get member by email")

		return 0, err
	}
	if memberByEmail.Member.Email != "" {
		return 0, ErrorDuplicateEmail
	}

	_, err = u.memberRepository.CreateMember(ctx, m)
	if err != nil {
		log.Println("Error Create member")

		return 0, err
	}

	memberByEmail, err = u.memberRepository.GetMemberByEmail(ctx, m.Email)
	if err != nil && err != ErrorItemNotFound {
		log.Println("Error get member by email")

		return 0, err
	}

	return memberByEmail.ID, nil
}

func validateRequireFields(m member.Member) error {
	if m.FirstName == "" {
		return ErrorFirstNameIsRequire
	}

	if m.LastName == "" {
		return ErrorLastNameIsRequire
	}

	if m.Phone == "" {
		return ErrorPhoneIsRequire
	}

	if m.Email == "" {
		return ErrorEmailIsRequire
	}

	if m.Address == "" {
		return ErrorAddressIsRequire
	}

	return nil
}
