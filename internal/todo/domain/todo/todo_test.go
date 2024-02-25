package todo_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hexisa_go_nal_todo/internal/todo/domain/todo"
)

func TestNewTodo(t *testing.T) {
	expected := todo.NewTodoFixture(
		"01HPCHEC5HJ37MC7D04PRCTJWK",
		"01HQ3M72Y0AXV0MNPHTXP96C4M",
		"test",
		"test",
		0,
	)
	type args struct {
		ulid    string
		userId  string
		title   string
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    *todo.Todo
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				ulid:    "01HPCHEC5HJ37MC7D04PRCTJWK",
				userId:  "01HQ3M72Y0AXV0MNPHTXP96C4M",
				title:   "test",
				content: "test",
			},
			want:    expected,
			wantErr: false,
		},
		{
			name: "準正常系:100文字以上の名前で作成できない",
			args: args{
				ulid:    "01HPCHEC5HJ37MC7D04PRCTJWK",
				userId:  "01HQ3M72Y0AXV0MNPHTXP96C4M",
				title:   strings.Repeat("a", 100),
				content: "test",
			},
			want:    &todo.Todo{},
			wantErr: true,
		},
		{
			name: "準正常系:1000文字以上の内容で作成できない",
			args: args{
				ulid:    "01HPCHEC5HJ37MC7D04PRCTJWK",
				userId:  "01HQ3M72Y0AXV0MNPHTXP96C4M",
				title:   "test",
				content: strings.Repeat("a", 1000),
			},
			want:    &todo.Todo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := todo.NewTodo(tt.args.ulid, tt.args.userId, tt.args.title, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(todo.Todo{})
			if d := cmp.Diff(got, tt.want, opt); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
