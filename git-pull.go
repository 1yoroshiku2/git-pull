package main

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {
	// 用户名和密码
	username := "huangshuai@keytop.com.cn"
	password := "ahourashi!12"

	// 通过命令行参数获取分支变量
	// args := os.Args
	// if len(args) < 2 {
	// 	fmt.Println("Please provide the branch name as a command-line argument.")
	// 	return
	// }
	// branch := args[1]
	//branch := "release_v1.0.0"
	branch := "release_v1111"
	//targetDirectory := "2.sql脚本" // 要拉取的目标目录

	// 设置要拉取的远程仓库URL和本地目录路径
	remoteURL := "https://e.coding.net/ktsoft/yongcepingtaipro2.0/deployment-information.git"
	localPath := "C:\\Users\\hs\\Desktop\\test-go\\git-pull"
	//localPath := "/home/tools/git"

	// 克隆仓库
	git.PlainClone(localPath, false, &git.CloneOptions{
		URL:      remoteURL,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		ReferenceName: plumbing.ReferenceName("refs/heads/" + branch),
		SingleBranch:  true,
		//Depth:         1, // 限制克隆的历史记录深度
	})
	// if err != nil {
	// 	fmt.Println("Failed to clone repository:", err)
	// 	return
	// }

	fmt.Println("Code successfully cloned and checked out!")
}
