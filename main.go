package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	// CombinfalseedOutput()でないと標準出力と標準エラー出力を両方キャッチできないことの確認
	// 以下のような出力となり、標準エラー出力がキャッチするためにはCombinedOutput()を使う必要があることがわかる
	// ----------------
	// Run()でエラーが発生しました: exit status 127
	// -----------------
	// Output()でエラーが発生しました: exit status 127
	// エラー時のOutputの内容: 処理を開始します
	// ステップ1完了
	//
	// -----------------
	// CombinedOutput()でエラーが発生しました: exit status 127
	// エラー時のOutputの内容(標準エラー出力含む): 処理を開始します
	// ステップ1完了
	// ./fail_script.sh: line 8: invalidcommand: command not found

	if true {
		fmt.Println("===================================================")
		fmt.Println("==============fail_script.shを実行する例=============")
		fmt.Println("===================================================")
		// Run()だと標準出力をキャッチできない
		fmt.Println("-----------------")
		cmd := exec.Command("bash", "-c", "./fail_script.sh")
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Run()でエラーが発生しました: %v\n", err)
		}

		// Output()だと標準エラー出力をキャッチできないことの確認
		fmt.Println("-----------------")
		cmd = exec.Command("bash", "-c", "./fail_script.sh")
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Output()でエラーが発生しました: %v\n", err)
			if len(output) > 0 {
				// 標準エラー出力は含まれない(`./script.sh: line 8: invalidcommand: command not found`が出ない)
				fmt.Printf("エラー時のOutputの内容: %s\n", output)
			}
		} else {
			fmt.Printf("Outputの内容: %s\n", output)
		}

		// CombinedOutput()だと標準出力と標準エラー出力を両方キャッチできる
		fmt.Println("-----------------")
		cmd = exec.Command("bash", "-c", "./fail_script.sh")
		output, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("CombinedOutput()でエラーが発生しました: %v\n", err)
			if len(output) > 0 {
				fmt.Printf("エラー時のOutputの内容(標準エラー出力含む): %s\n", output)
			}
		} else {
			fmt.Printf("Outputの内容: %s\n", output)
		}
	}

	// CombinedOutputで今回やりたいことが実現できそうか、具体的に確認
	if true {
		fmt.Println("===================================================")
		fmt.Println("==============catを実行する例=============")
		fmt.Println("===================================================")
		// 既存の動き
		if true {
			fmt.Println("-----------------")
			cmd := exec.Command("bash", "-c", "readlink -f ./dummy.txt")
			output, err := cmd.Output()
			if err != nil {
				fmt.Printf("CombinedOutput()でエラーが発生しました: %v\n", err)
				if len(output) > 0 {
					fmt.Printf("エラー時のOutputの内容(標準エラー出力含む): %s\n", output)
				}
			} else {
				fmt.Printf("Outputの内容: %s\n", output)
			}
			cmd = exec.Command("cat", strings.TrimSpace(string(output)))
			output, err = cmd.Output()
			if err != nil {
				fmt.Printf("CombinedOutput()でエラーが発生しました: %v\n", err)
				if len(output) > 0 {
					fmt.Printf("エラー時のOutputの内容(標準エラー出力含む): %s\n", output)
				}
			} else {
				fmt.Printf("Outputの内容: %s\n", output)
			}

		}
		// CombinedOutputバージョン
		if true {
			// script path を持ってきてそれをcatする
			fmt.Println("-----------------")
			cmd := exec.Command("bash", "-c", "readlink -f ./dummy.txt")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("CombinedOutput()でエラーが発生しました: %v\n", err)
				if len(output) > 0 {
					fmt.Printf("エラー時のOutputの内容(標準エラー出力含む): %s\n", output)
				}
			} else {
				fmt.Printf("Outputの内容: %s\n", output)
			}
			path := strings.TrimSpace((string(output)))
			fmt.Printf("%s\n", path)

			cmd = exec.Command("cat", path)
			//cmd = exec.Command("bash", "-c", "cat", "-f", string(output))
			output, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("CombinedOutput()でエラーが発生しました: %v\n", err)
				if len(output) > 0 {
					fmt.Printf("エラー時のOutputの内容(標準エラー出力含む): %s\n", output)
				}
			} else {
				fmt.Printf("Outputの内容: %s\n", output)
			}

		}
	}
}
