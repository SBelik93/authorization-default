package models

const (
	SchemaPublic  = "pubic"
	SchemaBooking = "booking"
	SchemaAuth    = "auth"
)

// ERRORS
const (
	ErrorUnexpected    = "ошибка с соединением, пожалуйста попробуйте позже"
	ErrorFormatEmail   = "неверный формат почты"
	ErrorFormatPhone   = "неверный формат телефона"
	ErrorParsingBody   = "неверное тело запроса"
	ErrorChanel        = "ошибка в канале верификации"
	ErrorUserNotExists = "пользователь с указанным номером не существует"
	ErrorSigIn       = "неверный логин или пароль"
)

// COLORS
const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)
