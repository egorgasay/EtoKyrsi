package usecase

import (
	"checkwork/internal/entity"
	"checkwork/internal/repository"
	mock_repository "checkwork/internal/repository/mocks"
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

type mockBehavior func(r *mock_repository.MockIStorage)

func Test_CreateUser(t *testing.T) {
	tests := []struct {
		name         string
		username     string
		password     string
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name:     "OK",
			username: "admin",
			password: "admin",
			wantErr:  false,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().CreateUser("admin", "admin").
					Return(nil).AnyTimes()
			},
		},
		{
			name:     "Duplicate",
			username: "admin",
			password: "admin",
			wantErr:  true,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().CreateUser("admin", "admin").
					Return(repository.AlreadyExistsErr).AnyTimes()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			test.mockBehavior(repos)
			logic := New(repos)
			err := logic.CreateUser(test.username, test.password)
			if test.wantErr && err == nil {
				t.Fail()
			}
		})
	}
}

func Test_CheckPassword(t *testing.T) {
	tests := []struct {
		name         string
		username     string
		password     string
		res          bool
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name:     "OK",
			username: "admin",
			password: "admin",
			res:      true,
			wantErr:  false,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().CheckPassword("admin", "admin").
					Return(true, nil).AnyTimes()
			},
		},
		{
			name:     "Wrong Password",
			username: "admin",
			password: "admin1",
			wantErr:  true,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().CheckPassword("admin", "admin1").
					Return(false, repository.WrongPasswordErr).AnyTimes()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			test.mockBehavior(repos)
			logic := New(repos)
			res, err := logic.CheckPassword(test.username, test.password)
			if test.wantErr && err == nil {
				t.Fail()
			}

			if test.res != res {
				t.Fail()
			}
		})
	}
}

func TestUseCase_CheckIsMentor(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Ok",
			args: args{
				username: "CHANGE ME",
			},
			want: true,
		},
		{
			name: "Not a mentor",
			args: args{
				username: "Rafik",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			logic := New(repos)
			//tt.mockBehavior(repos)

			got, err := logic.CheckIsMentor(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIsMentor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckIsMentor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_CheckIsPending(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name         string
		args         args
		want         bool
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "Ok",
			args: args{
				username: "admin",
			},
			want:    true,
			wantErr: false,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().CheckIsPending("admin").
					Return(true, nil).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			logic := New(repos)
			tt.mockBehavior(repos)
			got, err := logic.CheckIsPending(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIsPending() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckIsPending() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_DeleteTasks(t *testing.T) {
	type args struct {
		username string
		number   string
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "Ok",
			args: args{
				username: "",
				number:   "12",
			},
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().DeleteTask(12).
					Return(nil).AnyTimes()
			},
		},
		{
			name: "Wrong number",
			args: args{
				username: "",
				number:   "12d",
			},
			wantErr: true,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().DeleteTask("12d").
					Return(errors.New("")).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			logic := New(repos)
			tt.mockBehavior(repos)

			if err := logic.DeleteTasks(tt.args.username, tt.args.number); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTasks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetTask(t *testing.T) {
	type args struct {
		username string
		number   string
	}
	tests := []struct {
		name         string
		args         args
		wantTask     entity.Task
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "Ok",
			args: args{
				username: "",
				number:   "-99",
			},
			wantTask: entity.Task{
				Name:   "test",
				Number: -99,
				Text:   "@header\nTest\n",
			},
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetTitle(-99).
					Return("test", nil).AnyTimes()
			},
		},
		{
			name: "No such file",
			args: args{
				username: "",
				number:   "/99",
			},
			wantTask: entity.Task{
				Name:   "",
				Number: 0,
				Text:   "",
			},
			wantErr: true,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetTitle(99).
					Return("test", nil).AnyTimes()
			},
		},
		{
			name: "No such title",
			args: args{
				username: "",
				number:   "-99",
			},
			wantTask: entity.Task{
				Name:   "",
				Number: 0,
				Text:   "",
			},
			wantErr: true,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetTitle(-99).
					Return("", sql.ErrNoRows).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			logic := New(repos)
			tt.mockBehavior(repos)

			gotTask, err := logic.GetTask(tt.args.username, tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("GetTask() gotTask = %v, want %v", gotTask, tt.wantTask)
			}
		})
	}
}

func TestUseCase_GetTaskIDAndMsg(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name         string
		args         args
		want         int
		want1        sql.NullString
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "Ok",
			args: args{
				username: "",
			},
			want: 1,
			want1: sql.NullString{
				String: "Переделывай",
				Valid:  true,
			},
			wantErr: false,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetTaskIDAndMsg("").
					Return(
						1,
						sql.NullString{
							String: "Переделывай",
							Valid:  true,
						},
						nil).AnyTimes()
			},
		},
		{
			name: "No such user",
			args: args{
				username: "",
			},
			want: 0,
			want1: sql.NullString{
				String: "",
				Valid:  false,
			},
			wantErr: true,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetTaskIDAndMsg("").
					Return(
						0,
						sql.NullString{
							String: "",
							Valid:  false,
						},
						sql.ErrNoRows).AnyTimes()
			},
		},
		{
			name: "DB Error",
			args: args{
				username: "",
			},
			want: 0,
			want1: sql.NullString{
				String: "",
				Valid:  false,
			},
			wantErr: true,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetTaskIDAndMsg("").
					Return(
						0,
						sql.NullString{
							String: "",
							Valid:  false,
						},
						errors.New("db error")).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			logic := New(repos)
			tt.mockBehavior(repos)

			got, got1, err := logic.GetTaskIDAndMsg(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTaskIDAndMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTaskIDAndMsg() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetTaskIDAndMsg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUseCase_GetTasks(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         []entity.Task
		wantErr      bool
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetTasks("").
					Return(
						[]entity.Task{
							entity.Task{
								Name:   "test",
								Number: 1,
								Text:   "@header\ntest\n",
							},
						}, nil).AnyTimes()
			},
			args: args{
				username: "",
			},
			want: []entity.Task{
				entity.Task{
					Name:   "test",
					Number: 1,
					Text:   "@header\ntest\n",
				},
			},
		},
		{
			name: "DB Error",
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetTasks("").
					Return(
						[]entity.Task{}, errors.New("db error")).AnyTimes()
			},
			args: args{
				username: "",
			},
			wantErr: true,
			want:    []entity.Task{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			logic := New(repos)
			tt.mockBehavior(repos)

			got, err := logic.GetTasks(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetUsers(t *testing.T) {
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		want         []entity.User
		wantErr      bool
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetUsers().
					Return([]entity.User{
						entity.User{
							Name:        "test",
							Level:       1,
							LastComment: "Переделывай",
						},
					}, nil).AnyTimes()
			},
			want: []entity.User{
				entity.User{
					Name:        "test",
					Level:       1,
					LastComment: "Переделывай",
				},
			},
		},
		{
			name: "No users",
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetUsers().
					Return([]entity.User{}, sql.ErrNoRows).AnyTimes()
			},
			want:    []entity.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			logic := New(repos)
			tt.mockBehavior(repos)

			got, err := logic.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetWorks(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         []repository.Work
		wantErr      bool
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().GetWorks("").
					Return([]repository.Work{
						{
							Student: "test",
							Link:    "test.com/1",
						},
					}, nil).AnyTimes()
			},
			want: []repository.Work{
				{
					Student: "test",
					Link:    "test.com/1",
				},
			},
			args: args{
				username: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			logic := New(repos)
			tt.mockBehavior(repos)

			got, err := logic.GetWorks(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWorks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_HandleUserWork(t *testing.T) {
	type args struct {
		username string
		student  string
		verdict  string
		msg      string
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "Ok",
			args: args{
				username: "",
				student:  "test",
				verdict:  "bad",
				msg:      "bad",
			},
			wantErr: false,
			mockBehavior: func(r *mock_repository.MockIStorage) {
				r.EXPECT().DeletePullRequest("", "test").
					Return("test", nil).AnyTimes()
				r.EXPECT().SetPending("", "bad").
					Return(nil).AnyTimes()
				r.EXPECT().SetVerdict("test", "bad").
					Return(nil).AnyTimes()
				r.EXPECT().UpdateUserScore("test").
					Return(nil).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_repository.NewMockIStorage(c)
			logic := New(repos)
			tt.mockBehavior(repos)
			if err := logic.HandleUserWork(tt.args.username, tt.args.student, tt.args.verdict, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("HandleUserWork() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_SendPullRequest(t *testing.T) {
	type fields struct {
		storage repository.IStorage
	}
	type args struct {
		line    string
		student string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UseCase{
				storage: tt.fields.storage,
			}
			if err := uc.SendPullRequest(tt.args.line, tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("SendPullRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
