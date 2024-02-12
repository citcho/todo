package user_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hexisa_go_nal_todo/internal/user/domain/user"
	"github.com/hexisa_go_nal_todo/internal/user/domain/user/mock"
)

func TestUserService_Exists(t *testing.T) {
	u := user.NewUserFixture(
		"01HPCHEC5HJ37MC7D04PRCTJWK",
		"test",
		"test@example.com",
		"$2a$10$laM6Chj7KqMTSss1nzp21em0YnX95iYRgRR54zqLgWPJrKA1/bOaK",
	)

	type fields struct {
		prepareMockFn func(*mock.MockIUserRepository)
	}
	type args struct {
		ctx context.Context
		u   *user.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "正常系:ユーザーが存在することを確認できる",
			fields: fields{
				prepareMockFn: func(mur *mock.MockIUserRepository) {
					mur.EXPECT().
						Exists(context.Background(), u).
						Return(true, nil)
				},
			},
			args: args{
				ctx: context.Background(),
				u:   u,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "正常系:ユーザーが存在しないことを確認できる",
			fields: fields{
				prepareMockFn: func(mur *mock.MockIUserRepository) {
					mur.EXPECT().
						Exists(context.Background(), u).
						Return(false, nil)
				},
			},
			args: args{
				ctx: context.Background(),
				u:   u,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mur := mock.NewMockIUserRepository(mockCtrl)
			tt.fields.prepareMockFn(mur)

			sut := user.NewUserService(mur)
			got, err := sut.Exists(tt.args.ctx, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
