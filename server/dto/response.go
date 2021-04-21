package dto

var (
	Success          = &baseResponse{code: 0, message: "OK"}
	AlreadyConnected = &baseResponse{code: 204, message: "Handshake already completed"}
	IllegalArgument  = &baseResponse{code: 403, message: "Illeagal argument exceptiono"}
	Failed           = &baseResponse{code: 404, message: "Unknown Error"}
)

type baseResponse struct {
	code    int
	message string
}
