package services

import (
	"errors"
	"reflect"
	"testing"
	"todo-api/internal/entity"
	"todo-api/internal/repository"
	"todo-api/internal/request"

	"github.com/golang/mock/gomock"
)

func Test_todoService_GetAllTodo(t *testing.T) {
	type fields struct {
		todoRepository *repository.MockTodoRepository
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		fields  fields
		want    []entity.Todo
		wantErr bool
		before  func(f fields)
	}{
		{
			name: "fail",
			fields: fields{
				todoRepository: repository.NewMockTodoRepository(ctrl),
			},
			want:    nil,
			wantErr: true,
			before: func(f fields) {
				f.todoRepository.EXPECT().GetAllTodo().Return(nil, errors.New("Qntlo"))
			},
		},
		{
			name: "success",
			fields: fields{
				todoRepository: repository.NewMockTodoRepository(ctrl),
			},
			want: []entity.Todo{
				{
					ID: 1,
				},
			},
			wantErr: false,
			before: func(f fields) {
				f.todoRepository.EXPECT().GetAllTodo().Return([]entity.Todo{
					{
						ID: 1,
					},
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := todoService{
				todoRepository: tt.fields.todoRepository,
			}
			tt.before(tt.fields)
			got, err := tr.GetAllTodo()
			if (err != nil) != tt.wantErr {
				t.Errorf("todoService.GetAllTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoService.GetAllTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_FindByIDTodo(t *testing.T) {
	type fields struct {
		todoRepository *repository.MockTodoRepository
	}
	type args struct {
		id int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Todo
		wantErr bool
		before  func(f fields)
	}{
		// TODO: Add test cases.
		{
			name: "fail",
			fields: fields{
				todoRepository: repository.NewMockTodoRepository(ctrl),
			},
			args: args{
				id: 1,
			},
			want:    entity.Todo{},
			wantErr: true,
			before: func(f fields) {
				f.todoRepository.EXPECT().FindByIDTodo(1).Return(entity.Todo{}, errors.New("tes"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := todoService{
				todoRepository: tt.fields.todoRepository,
			}
			tt.before(tt.fields)
			got, err := tr.FindByIDTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoService.FindByIDTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoService.FindByIDTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_CreateTodo(t *testing.T) {
	type fields struct {
		todoRepository *repository.MockTodoRepository
	}
	type args struct {
		todoRequest request.TodoRequest
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Todo
		wantErr bool
		before  func(f fields)
	}{
		// TODO: Add test cases.
		{
			name: "fail",
			fields: fields{
				todoRepository: repository.NewMockTodoRepository(ctrl),
			},
			args: args{
				todoRequest: request.TodoRequest{},
			},
			want:    entity.Todo{},
			wantErr: true,
			before: func(f fields) {
				f.todoRepository.EXPECT().CreateTodo(entity.Todo{}).Return(entity.Todo{}, errors.New("Error create todo"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := todoService{
				todoRepository: tt.fields.todoRepository,
			}

			tt.before(tt.fields)

			got, err := tr.CreateTodo(tt.args.todoRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoService.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoService.CreateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
