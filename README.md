# fluent-attacker-rpm

## 概要

* fluentdサーバーへ負荷をかけるクライアントプログラムです
* Request Per Second,Message Length,Post Countを調整することができます
* Golang

## 使用方法

go run fluent-attacker-rpm.go

## オプション

* -h [str] fluentdサーバー送信先アドレス(Default:127.0.0.1)
* -p [int] fluentdサーバー待ち受けポート(Default:24224)
* -c [int] テストデータ投稿数(Default:3000)
* -t [str] テストデータのTAG(Default:"test.tag")
* -t [str] テストデータのメッセージ長(Default:1000)
* -r [int] 秒間送信数(Request Per Second)(Default:300)


## コンソール画面ログ

```
=====================================
Insert Keys=> 3000 RPS Set => 250 Message Length => 10
=====================================
0 loop - TaskSet Done.
250 loop - TaskSet Done.
500 loop - TaskSet Done.
750 loop - TaskSet Done.
1000 loop - TaskSet Done.
1250 loop - TaskSet Done.
1500 loop - TaskSet Done.
1750 loop - TaskSet Done.
2000 loop - TaskSet Done.
2250 loop - TaskSet Done.
2500 loop - TaskSet Done.
2750 loop - TaskSet Done.
3000 Task - All Job Done.
=====================================
Exec Time => 12.435550Sec Real RPS => 241.243851rps(96.497540%)
=====================================
```

* 最初に設定値が表示されます
* 実行中は定期的に進捗が表示されます
* 最後に「実行時間」と「実RPM」「設定RPMと実RPMとの差異」が表示されます
