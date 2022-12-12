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

	if true {

	}

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
	// _, err = golang.Exec(dagger.ContainerExecOpts{
	// 	Args: []string{"go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1"},
	// }).Stdout(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// コンテナ内のsrcディレクトリへマウント
	// srcディレクトリをワークディレクトリに設定
	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")

	// ビルドコマンドを実行
	path := "build/"
	golang = golang.WithExec([]string{"go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1"})
	golang = golang.WithExec([]string{"golangci-lint", "run", "./..."})
	// golang = golang.WithExec([]string{"go", "test", "-v", "./..."})
	// golang = golang.WithExec([]string{"go", "build", "-o", path})

	// コンテナ内のbuildディレクトリへの参照を取得
	output := golang.Directory(path)

	// コンテナからホストへbuildディレクトリの内容を書き込む
	_, err = output.Export(ctx, path)
	if err != nil {
		return err
	}

	return nil
}
