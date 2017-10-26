# CLI 命令行实用程序开发实战 - Agenda
---
### **小组成员**
#### 李佳    15331151
#### 李辉旭  15331150
#### 李果    15331145
---
## 1、概述
  命令行实用程序并不是都象 cat、more、grep 是简单命令。[go](https://go-zh.org/cmd/go/)项目管理程序，类似 java 项目管理 maven、Nodejs 项目管理程序 npm、git 命令行客户端、 docker 与 kubernetes 容器管理工具等等都是采用了较复杂的命令行。即一个实用程序同时支持多个子命令，每个子命令有各自独立的参数，命令之间可能存在共享的代码或逻辑，同时随着产品的发展，这些命令可能发生功能变化、添加新命令等。因此，符合 [OCP](https://en.wikipedia.org/wiki/Open/closed_principle)原则 的设计是至关重要的编程需求。
  
### 任务目标

1.熟悉 go 命令行工具管理项目
2.综合使用 go 的函数、数据结构与接口，编写一个简单命令行应用 agenda
3.使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令,不会影响其他命令的代码
4.项目部署在 Github 上，合适多人协作，特别是代码归并
5.支持日志（原则上不使用debug调试程序）

---
## tip1: 相关学习资料
- [面向对象设计思想与 golang 编程](http://blog.csdn.net/pmlpml/article/details/78326769)
- [《Go程序设计语言》要点总结——程序结构](http://time-track.cn/gopl-notes-program-structure.html)
- [《Go程序设计语言》要点总结——数据类型](http://time-track.cn/gopl-notes-types.html)
- [《Go程序设计语言》要点总结——函数](http://time-track.cn/gopl-notes-function.html)
- [《Go程序设计语言》要点总结——方法](http://time-track.cn/gopl-notes-function.html)
- [《Go程序设计语言》要点总结——接口](http://time-track.cn/gopl-notes-interface.html)
- [golang语法总结（二十一）：方法method](http://blog.csdn.net/qq245671051/article/details/50722802)

---

## tip2: 安装 cobra

使用命令 `go get -v github.com/spf13/cobra/cobra` 下载过程中，会出提示如下错误
>Fetching https://golang.org/x/sys/unix?go-get=1
>
>https fetch failed: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: i/o timeout

这是熟悉的错误，请在 `$GOPATH/src/golang.org/x` 目录下用 `git clone` 下载 `sys` 和 `text` 项目，然后使用 `go install github.com/spf13/cobra/cobra`, 安装后在 `$GOBIN` 下出现了 `cobra` 可执行程序。

## Cobra 的简单使用

创建一个处理命令 `agenda register -uTestUser` 或 `agenda register --user=TestUser` 的小程序。

简要步骤如下：
>cobra init
>
>cobra add register

需要的文件就产生了。 你需要阅读 `main.go` 的 `main()` ; `root.go` 的 `Execute()`; 最后修改 `register.go`, `init()` 添加：

>registerCmd.Flags().StringP("user", "u", "Anonymous", "Help message for username")

Run 匿名回调函数中添加：

>username, _ := cmd.Flags().GetString("user")
>
>fmt.Println("register called by " + username)

测试命令：
>$ go run main.go register --user=TestUser
>
>register called by TestUser

---
## 2.agenda 开发项目
### 需求描述

- 业务需求：见后面需求

- 功能需求： 设计一组命令完成 agenda 的管理，例如：
>agenda help ：列出命令说明
>agenda register -uUserName --password pass -email=a@xxx.com ：注册用户
>agenda help register ：列出 register 命令的描述
>agenda cm ... : 创建一个会议
>
>原则上一个命令对应一个业务功能

- 持久化要求：
>使用 json 存储 User 和 Meeting 实体
>当前用户信息存储在 curUser.txt 中

- 开发需求:
>团队：2-3人，一人作为 master 创建程序框架，其他人 fork 该项目，所有人同时开发。团队 不能少于 2 人
时间：两周完成

- 项目目录
>cmd ：存放命令实现代码
>entity ：存放 User 和 Meeting 对象读写与处理逻辑

- 其他目录 ： 自由添加
- 日志服务
>使用 log 包记录命令执行情况

---

### Agenda 业务需求

- 用户注册
注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。

- 用户登录
用户使用用户名和密码登录 Agenda 系统。
用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。

- 用户登出
已登录的用户登出系统后，只能使用用户注册和用户登录功能。

- 用户查询
已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。

- 用户删除

1.已登录的用户可以删除本用户账户（即销号）。

2.操作成功，需反馈一个成功注销的信息；否则，反馈一个失败注销的信息。

3.删除成功则退出系统登录状态。删除后，该用户账户不再存在。

4.用户账户删除以后：
>
  以该用户为发起者的会议将被删除
>  
  以该用户为参与者的会议将从参与者列表中移除该用户。若因此造成会议参与者人数为0，则会议也将被删除。

- 创建会议

1.已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册
用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：
>
 会议主题(title)（在会议列表中具有唯一性）
> 
  会议参与者(participator)
> 
  会议起始时间(start time)
> 
  会议结束时间(end time)
  
2.注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。

3.用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。

- 增删会议参与者。

1.已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。

2.增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。

3.删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。

- 查询会议

1.已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。

2.用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。

3.在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。

4.注意，查询会议的结果应包括用户作为 发起者或参与者 的会议。

- 取消会议

1.已登录的用户可以取消 自己发起 的某一会议安排。

2.取消会议时，需提供唯一标识：会议主题（title）。

- 退出会议

1.已登录的用户可以退出 自己参与 的某一会议安排。

2.退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议 参与者人数为0，则会议也将被删除。

- 清空会议

已登录的用户可以清空 自己发起 的所有会议安排。

---





