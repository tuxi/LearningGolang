package main

import (
	"fmt"
	"os/user"
)

/*
os包结构介绍
- Go语言标准库中os包提供了不依赖平台的操作系统接口
- 设计为Unix风格的，而错误处理是go风格的，失败的调用会返回错误值而非错误码，通常错误值里包含更多信息
- os包及子包功能
```
	--os/exec包,负责执行外部命令
	--os/signal对输入信息的访问
	--os/user 通过名称或ID，查询用户账号
```
- 在os/user中提供了User结构体，表示操作系统用户
	* Uid 用户id
	* Gid 所属组id
	* Username 用户名称
	* Name 所属组名
	* HomeDir 用户对应文件夹路径
```go
// User represents a user account.
type User struct {
	// Uid is the user ID.
	// On POSIX systems, this is a decimal number representing the uid.
	// On Windows, this is a security identifier (SID) in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Uid string
	// Gid is the primary group ID.
	// On POSIX systems, this is a decimal number representing the gid.
	// On Windows, this is a SID in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Gid string
	// Username is the login name.
	Username string
	// Name is the user's real or display name.
	// It might be blank.
	// On POSIX systems, this is the first (or only) entry in the GECOS field
	// list.
	// On Windows, this is the user's display name.
	// On Plan 9, this is the contents of /dev/user.
	Name string
	// HomeDir is the path to the user's home directory (if they have one).
	HomeDir string
}
```
* 在os/user中的Group表示用户所属组
  * Gid 组的id
  * Name 组的名称
```go
// Group represents a grouping of users.
//
// On POSIX systems Gid contains a decimal number representing the group ID.
type Group struct {
	Gid  string // group ID
	Name string // group name
}
```
* 整个os/user包中内容比较少,提供了两个错误类型和获取当前用户,查找用户
```go
type UnknownUserError
  func (e UnknownUserError) Error() string
type UnknownUserIdError
  func (e UnknownUserIdError) Error() string
type User
  func Current() (*User, error)
  func Lookup(username string) (*User, error)
  func LookupId(uid string) (*User, error)
```

*/

func main() {
	// 获取当前用户或查找用户后获取用户信息

	// 获取当前登录user
	u, _ := user.Current()
	fmt.Println(u) // &{502 20 xiaoyuan xiaoyuan /Users/xiaoyuan}

	// 按照用户名查找指定用户信息，用户名必须是完整名称
	u1, err := user.Lookup("xiaoyuan")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u1) // &{502 20 xiaoyuan xiaoyuan /Users/xiaoyuan}

	// 用户名称
	fmt.Println(u.Name) // xiaoyuan
	// 组id
	fmt.Println(u.Gid) // 20
	// 用户目录
	fmt.Println(u.HomeDir) // /Users/xiaoyuan
	// 用户 id
	fmt.Println(u.Uid) // 502
	// 用户名称
	fmt.Println(u.Username) // xiaoyuan
}
