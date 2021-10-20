package implement

import (
	"context"
	"fmt"
	"grpc/services"
	"strconv"
)

type Member struct{}

func (m *Member) GetUserInfo(ctx context.Context, req *services.GetUserInfoRequest) (res *services.GetUserInfoResponse, err error) {
	fmt.Println("open GetUserInfo")
	res = &services.GetUserInfoResponse{
		Code: 200,
		Info: &services.Member{
			Id:   req.Id,
			Name: "小黑",
			Age:  18,
			Sex:  services.MemberSexEnum_SEX_BOY,
		},
	}
	return
}

func (m *Member) GetUserList(context.Context, *services.GetUserListRequest) (res *services.GetUserListResponse, err error) {
	fmt.Println("open GetUserList")
	var members []*services.Member
	for a := 1; a < 10; a++ {
		members = append(members, &services.Member{
			Id:   int64(a),
			Name: "小黑" + strconv.Itoa(a),
			Age:  18,
			Sex:  services.MemberSexEnum_SEX_BOY,
		})
	}
	res = &services.GetUserListResponse{
		Code: 200,
		Data: members,
	}
	return
}
