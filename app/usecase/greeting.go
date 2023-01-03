package usecase

import (
	"context"
	"fmt"

	"shimabox/example-s3-to-lambda-go/app/domain/errs"
	"shimabox/example-s3-to-lambda-go/app/domain/repository"
)

type GreetingUsecase interface {
	Handle(ctx context.Context, input GreetingInput, greetingRepository repository.GreetingRepository) error
}

type greetingUsecase struct{}

func NewGreetingUsecase() GreetingUsecase {
	return &greetingUsecase{}
}

func (u greetingUsecase) Handle(
	ctx context.Context,
	input GreetingInput,
	greetingRepository repository.GreetingRepository,
) error {
	// イベントをモデルに変換
	greetingEventList := input.EventToGreetingEventList()

	// 挨拶がなければerror
	if !greetingEventList.Exists() {
		return &errs.EventNotFound{}
	}

	// 挨拶の読み込み(複数来ていてもとりあえず最初の挨拶だけ)
	greetingEvent := greetingEventList.GetFirstEvent()
	greeting, err := greetingRepository.Read(ctx, greetingEvent.Key())

	// 挨拶の読み込みに失敗したらerror
	if err != nil {
		return &errs.ReadEventException{Err: err}
	}

	// 挨拶が取得できれば出力
	fmt.Println(greeting.Hello())

	return nil
}
