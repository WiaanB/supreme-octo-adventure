package task

import (
	"bytes"
	"os"
	"testing"
)

// Helper function to capture stdout
func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	_ = w.Close()
	os.Stdout = old
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestList_Add(t *testing.T) {
	type args struct {
		task string
	}
	tests := []struct {
		name string
		l    List
		args args
		want string
	}{
		{
			name: "Test Adding a Task",
			l:    nil,
			args: args{"My new task"},
			want: "added task to the list\n",
		},
		{
			name: "Ensure there is a Task provided",
			l:    nil,
			args: args{""},
			want: "task cannot be empty\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				tt.l.Add(tt.args.task)
			})
			if output != tt.want {
				t.Errorf("expectd output: %v, but got %v", tt.want, output)
			}
		})
	}
}

func TestList_Delete(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		l    List
		args args
		want string
	}{
		{
			name: "Could not find the task",
			l: []Task{{
				Id:        1,
				Task:      "Help me",
				Status:    "todo",
				CreatedAt: "foo",
				UpdatedAt: "bar",
			}},
			args: args{4},
			want: "task not found, check the list again??\n",
		},
		{
			name: "Delete the task successfully",
			l: []Task{{
				Id:        1,
				Task:      "Help me",
				Status:    "todo",
				CreatedAt: "foo",
				UpdatedAt: "bar",
			}},
			args: args{1},
			want: "removed task with id: 1\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				tt.l.Delete(tt.args.id)
			})
			if output != tt.want {
				t.Errorf("expectd output: %v, but got %v", tt.want, output)
			}
		})
	}
}

func TestList_List(t *testing.T) {
	type args struct {
		listType string
	}
	tests := []struct {
		name string
		l    List
		args args
		want string
	}{
		{
			name: "No tasks matching query",
			l: []Task{{
				Id:        1,
				Task:      "Todo test",
				Status:    "done",
				CreatedAt: "",
				UpdatedAt: "",
			}},
			args: args{listType: "todo"},
			want: "no tasks set to 'todo'\n",
		},
		{
			name: "Properly show all tasks for todo",
			l: []Task{{
				Id:        1,
				Task:      "Todo test",
				Status:    "todo",
				CreatedAt: "",
				UpdatedAt: "",
			}, {
				Id:        2,
				Task:      "Todo test 2",
				Status:    "in-progress",
				CreatedAt: "",
				UpdatedAt: "",
			}, {
				Id:        3,
				Task:      "Todo test 3",
				Status:    "done",
				CreatedAt: "",
				UpdatedAt: "",
			}},
			args: args{listType: "todo"},
			want: "id: 1, task: Todo test, status: todo\n",
		},
		{
			name: "Properly show all tasks for in-progress",
			l: []Task{{
				Id:        1,
				Task:      "Todo test",
				Status:    "todo",
				CreatedAt: "",
				UpdatedAt: "",
			}, {
				Id:        2,
				Task:      "Todo test 2",
				Status:    "in-progress",
				CreatedAt: "",
				UpdatedAt: "",
			}, {
				Id:        3,
				Task:      "Todo test 3",
				Status:    "done",
				CreatedAt: "",
				UpdatedAt: "",
			}},
			args: args{listType: "in-progress"},
			want: "id: 2, task: Todo test 2, status: in-progress\n",
		},
		{
			name: "Properly show all tasks for done",
			l: []Task{{
				Id:        1,
				Task:      "Todo test",
				Status:    "todo",
				CreatedAt: "",
				UpdatedAt: "",
			}, {
				Id:        2,
				Task:      "Todo test 2",
				Status:    "in-progress",
				CreatedAt: "",
				UpdatedAt: "",
			}, {
				Id:        3,
				Task:      "Todo test 3",
				Status:    "done",
				CreatedAt: "",
				UpdatedAt: "",
			}},
			args: args{listType: "done"},
			want: "id: 3, task: Todo test 3, status: done\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				tt.l.List(tt.args.listType)
			})
			if output != tt.want {
				t.Errorf("expectd output: %v, but got %v", tt.want, output)
			}
		})
	}
}

func TestList_MarkAs(t *testing.T) {
	type args struct {
		id     int
		status string
	}
	tests := []struct {
		name string
		l    List
		args args
		want string
	}{
		{
			name: "Mark a task as in-progress",
			l: []Task{{
				Id:        1,
				Task:      "My task",
				Status:    "todo",
				CreatedAt: "",
				UpdatedAt: "",
			}},
			args: args{id: 1, status: "in-progress"},
			want: "marked task with id: 1 as 'in-progress'\n",
		},
		{
			name: "Could not find the task",
			l: []Task{{
				Id:        1,
				Task:      "My task",
				Status:    "todo",
				CreatedAt: "",
				UpdatedAt: "",
			}},
			args: args{id: 2, status: "in-progress"},
			want: "task not found, check the list again??\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				tt.l.MarkAs(tt.args.id, tt.args.status)
			})
			if output != tt.want {
				t.Errorf("expectd output: %v, but got %v", tt.want, output)
			}
		})
	}
}

func TestList_Update(t *testing.T) {
	type args struct {
		id   int
		task string
	}
	tests := []struct {
		name string
		l    List
		args args
		want string
	}{
		{
			name: "Update a task",
			l: []Task{{
				Id:        1,
				Task:      "My task",
				Status:    "todo",
				CreatedAt: "",
				UpdatedAt: "",
			}},
			args: args{id: 1, task: "My updated task"},
			want: "updated task with id: 1\n",
		},
		{
			name: "Could not find the task",
			l: []Task{{
				Id:        1,
				Task:      "My task",
				Status:    "todo",
				CreatedAt: "",
				UpdatedAt: "",
			}},
			args: args{id: 2, task: "My updated task"},
			want: "task not found, check the list again??\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				tt.l.Update(tt.args.id, tt.args.task)
			})
			if output != tt.want {
				t.Errorf("expectd output: %v, but got %v", tt.want, output)
			}
		})
	}
}
