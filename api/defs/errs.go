package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json"error_code"`
}

type ErrResponse struct {
	HttpSc int
	Error Err
}

var(
	ErrorRequestBodeParseFailed = ErrResponse{
		HttpSc: 400,
		Error: Err{
			Error:"request body is not correct",
			ErrorCode: "001"},
	}

	ErrorNotAuthUser = ErrResponse{
		HttpSc: 401,
		Error: Err{
			Error:"User not authentication failed",
			ErrorCode: "002"},
	}
)