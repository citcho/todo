package user_test

import (
	"testing"

	"github.com/citcho/todo/internal/user/domain/user"
	"github.com/citcho/todo/internal/user/domain/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestNewUser(t *testing.T) {
	want := user.NewUserFixture(
		"01HPCHEC5HJ37MC7D04PRCTJWK",
		"test",
		"test@example.com",
		"$2a$10$laM6Chj7KqMTSss1nzp21em0YnX95iYRgRR54zqLgWPJrKA1/bOaK",
	)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mpe := mock.NewMockPasswordEncrypter(mockCtrl)
	mpe.EXPECT().
		Encrypt("Secretp@ssw0rd").
		Return("$2a$10$laM6Chj7KqMTSss1nzp21em0YnX95iYRgRR54zqLgWPJrKA1/bOaK").
		AnyTimes()

	type args struct {
		id       string
		name     string
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *user.User
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id:       "01HPCHEC5HJ37MC7D04PRCTJWK",
				name:     "test",
				email:    "test@example.com",
				password: "Secretp@ssw0rd",
			},
			want:    want,
			wantErr: false,
		},
		{
			name: "準正常系:10文字以上の名前で作成できない",
			args: args{
				id:       "01HPCHEC5HJ37MC7D04PRCTJWK",
				name:     "invalidname",
				email:    "test@example.com",
				password: "Secretp@ssw0rd",
			},
			want:    &user.User{},
			wantErr: true,
		},
		{
			name: "準正常系:不正なメールアドレスで作成できない",
			args: args{
				id:       "01HPCHEC5HJ37MC7D04PRCTJWK",
				name:     "test",
				email:    "invalidemail",
				password: "Secretp@ssw0rd",
			},
			want:    &user.User{},
			wantErr: true,
		},
		{
			name: "準正常系:不正なパスワードで作成できない",
			args: args{
				id:       "01HPCHEC5HJ37MC7D04PRCTJWK",
				name:     "test",
				email:    "test@example.com",
				password: "invalidpassword",
			},
			want:    &user.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := user.NewUser(tt.args.id, tt.args.name, tt.args.email, tt.args.password, mpe)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(user.User{})
			if d := cmp.Diff(got, tt.want, opt); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
