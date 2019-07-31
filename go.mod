module syrup-plum

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190730183949-1393eb018365
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190730215328-ed3277de2799
)

require (
	github.com/Unknwon/goconfig v0.0.0-20190425194916-3dba17dd7b9e
	github.com/lvxin0315/syrup-plum v0.0.0-20190731014910-9b1d774fc554
	github.com/pkg/errors v0.8.1
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945 // indirect
)
