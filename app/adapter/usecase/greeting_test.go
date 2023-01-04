package usecase_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"shimabox/example-s3-to-lambda-go/app/adapter/domain/errs"
	"shimabox/example-s3-to-lambda-go/app/adapter/domain/model"
	"shimabox/example-s3-to-lambda-go/app/adapter/usecase"
	"shimabox/example-s3-to-lambda-go/testhelper"
)

func Test_Handle_挨拶がなければerror_EventNotFoundが返る(t *testing.T) {
	testInput := TestInput{f: emptyGreetingEventList}
	greetingUsecase := usecase.NewGreetingUsecase()

	got := greetingUsecase.Handle(context.Background(), testInput, TestRep{})
	want := &errs.EventNotFound{}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Handle is mismatch (-want +got):\n%s", diff)
	}
}

func Test_Handle_挨拶の読み込みに失敗したらerror_ReadEventExceptionが返る(t *testing.T) {
	testInput := TestInput{f: greetingEventList}
	var expectedErr = &TestErr{"test read error"}
	testRep := TestRep{
		func() (*model.Greeting, error) {
			return nil, expectedErr
		},
	}
	greetingUsecase := usecase.NewGreetingUsecase()

	got := greetingUsecase.Handle(context.Background(), testInput, testRep)
	want := &errs.ReadEventException{Err: expectedErr}

	if diff := cmp.Diff(want.Error(), got.Error()); diff != "" {
		t.Errorf("Handle is mismatch (-want +got):\n%s", diff)
	}
}

func Test_Handle_挨拶の読み込みに成功したら標準出力される(t *testing.T) {
	testInput := TestInput{f: greetingEventList}
	testRep := TestRep{f: getGreeting}
	greetingUsecase := usecase.NewGreetingUsecase()

	got := testhelper.ExtractStdout(t, func() {
		err := greetingUsecase.Handle(context.Background(), testInput, testRep)
		if err != nil {
			return
		}
	})
	want := "テスト太郎: こんにちは"

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Handle is mismatch (-want +got):\n%s", diff)
	}
}

type TestErr struct {
	m string
}

func (t TestErr) Error() string {
	return t.m
}

type TestInput struct {
	f func() *model.GreetingEventList
}

func (t TestInput) EventToGreetingEventList() *model.GreetingEventList {
	return t.f()
}

func emptyGreetingEventList() *model.GreetingEventList {
	return &model.GreetingEventList{}
}
func greetingEventList() *model.GreetingEventList {
	list := []*model.GreetingEvent{
		model.NewGreetingEvent(model.NewGreetingEventKey("test.json")),
		model.NewGreetingEvent(model.NewGreetingEventKey("test2.json")),
	}
	return model.NewGreetingEventList(list)
}

type TestRep struct {
	f func() (*model.Greeting, error)
}

func (t TestRep) Read(ctx context.Context, key *model.GreetingEventKey) (*model.Greeting, error) {
	return t.f()
}

func getGreeting() (*model.Greeting, error) {
	return &model.Greeting{
		UserName: "テスト太郎",
		Message:  "こんにちは",
	}, nil
}
