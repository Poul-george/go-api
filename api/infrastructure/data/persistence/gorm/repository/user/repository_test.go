package user_test

import (
	"context"
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	userRepositoryImpl "github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/repository/user"
	"github.com/Poul-george/go-api/api/util/testhelper"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
	"time"
)

func TestRepository_FindByID(t *testing.T) {
	testhelper.Lock(t)
	testhelper.LoadFixture(t, testfixtures.Directory("testdata/fixtures/findbyid"))
	ctx := context.Background()
	r := userRepositoryImpl.NewRepository(handler.NewHandler(config.GetMySQLConfig()))

	//testLoc, _ := time.LoadLocation("Asia/Tokyo")

	type UserInfo struct {
		externalUserID identifier.ExternalUserID
		userID         identifier.UserID
	}
	tests := []struct {
		name    string
		args    UserInfo
		want    *model.User
		wantErr bool
	}{
		{
			name: "(正)externalUserIDだけを指定してuser情報が取得できる",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID("111"),
				userID:         identifier.UserID(0),
			},
			want: model.ReConstructorUser(
				1,
				"111",
				"test1",
				"$2a$10$tZN5qGGheum3BL9up8VhbOXpojUnlyb5vQEehb.rkPqV8VeP57aHu",
				"test@gmail.com",
				"test comments",
				time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
			),
			wantErr: false,
		},
		{
			name: "(正)UserIDだけを指定してuser情報が取得できる",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID(""),
				userID:         identifier.UserID(2),
			},
			want: model.ReConstructorUser(
				2,
				"222",
				"test2",
				"$2a$10$tZN5qGGheum3BL9up8VhbOXpojUnlyb5vQEehb.rkPqV8VeP57aHu",
				"test@gmail.com",
				"test comments",
				time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
			),
			wantErr: false,
		},
		{
			name: "(正)externalUserID&UserIDを指定してuser情報が取得できる",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID("333"),
				userID:         identifier.UserID(3),
			},
			want: model.ReConstructorUser(
				3,
				"333",
				"test3",
				"$2a$10$tZN5qGGheum3BL9up8VhbOXpojUnlyb5vQEehb.rkPqV8VeP57aHu",
				"test@gmail.com",
				"test comments",
				time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
			),
			wantErr: false,
		},
		{
			name: "(異)externalUserID&存在しないUserIDを指定してエラーになる",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID("333"),
				userID:         identifier.UserID(100),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "(異)削除済みのユーザーを取得できない",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID("444"),
				userID:         identifier.UserID(4),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.FindByID(ctx, tt.args.externalUserID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestUserRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opt := cmp.AllowUnexported(model.User{})
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(model.User{}, "updatedAt"), opt); diff != "" {
				t.Errorf("TestUserRepository.FindByID() diff(-got +want)\n%s", diff)
			}
		})
	}
}
