package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// Daggerクライアントの作成
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	// ホストのカレントディレクトリへの参照を取得
	src := client.Host().Directory(".")

	// Goの最新Verのコンテナイメージを取得
	golang := client.Container().From("golang:latest")

	// コンテナ内のsrcディレクトリへマウント
	// srcディレクトリをワークディレクトリに設定
	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")

	// ビルドコマンドを実行
	// golang = golang.WithExec([]string{"go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1"})
	// golang = golang.WithExec([]string{"golangci-lint", "run", "./..."})
	golang = golang.WithExec([]string{"go", "test", "-v", "./..."})

	// path := "./build"
	// golang = golang.WithExec([]string{"go", "build", "-o", path})

	// コンテナ内のbuildディレクトリへの参照を取得
	output, err := golang.Stderr(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", output)

	// コンテナからホストへbuildディレクトリの内容を書き込む
	// _, err = output.Export(ctx, path)
	// if err != nil {
	// 	return err
	// }

	return nil
}
