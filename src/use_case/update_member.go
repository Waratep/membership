package use_case

import (
	"github.com/Waratep/membership/src/entity/member"
	"github.com/gin-gonic/gin"
)

func UpdateMember(ctx gin.Context, m member.Member) (int64, error) {
	validateRequireFields(m)

	return 0, nil
}
