package errors

const (
	ERROR_GET_POST_FAIL = "20001"
)

var PostsErrorMsg = map[string]string{
	ERROR_GET_POST_FAIL: "Post(s) is not available",
}