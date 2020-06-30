package e

// 错误代码包


const(
	SUCCESS = 200
	ERROR = 500

	ErrNotExistRoom = 10001
	ErrRoomPassword = 10002
	ErrRoomExists = 10003

	ErrAuthCheckTokenFail = 20001
	ErrAuthCheckTokenTimeout = 20002
	ErrAuthToken = 20003
	ErrAuth = 20004

	ErrUserExists = 30001
	ErrUserNotExists = 30002
	ErrUserPassword = 30003

	ErrFormat = 40001
	)