package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/diki-haryadi/go-micro-template/pkg/constant"
	"log"
	"net/http"
	"strconv"
)

var (
	ErrBadRequest                    = errors.New("Bad request")
	ErrForbiddenResource             = errors.New("Forbidden resource")
	ErrNotFound                      = errors.New("Not Found")
	ErrPreConditionFailed            = errors.New("Precondition failed")
	ErrInternalServerError           = errors.New("Internal server error")
	ErrTimeoutError                  = errors.New("Timeout error")
	ErrUnauthorized                  = errors.New("Unauthorized")
	ErrConflict                      = errors.New("Conflict")
	ErrMethodNotAllowed              = errors.New("Method not allowed")
	ErrInvalidGrantType              = errors.New("Invalid grant type")
	ErrInvalidClientIDOrSecret       = errors.New("Invalid client ID or secret")
	ErrAuthorizationCodeNotFound     = errors.New("Authorization code not found")
	ErrAuthorizationCodeExpired      = errors.New("Authorization code expired")
	ErrInvalidRedirectURI            = errors.New("Invalid redirect URI")
	ErrInvalidScope                  = errors.New("Invalid scope")
	ErrInvalidUsernameOrPassword     = errors.New("Invalid username or password")
	ErrRefreshTokenNotFound          = errors.New("Refresh token not found")
	ErrRefreshTokenExpired           = errors.New("Refresh token expired")
	ErrRequestedScopeCannotBeGreater = errors.New("Requested scope cannot be greater")
	ErrTokenMissing                  = errors.New("Token missing")
	ErrTokenHintInvalid              = errors.New("Invalid token hint")
	ErrAccessTokenNotFound           = errors.New("Access token not found")
	ErrAccessTokenExpired            = errors.New("Access token expired")
	ErrClientNotFound                = errors.New("Client not found")
	ErrInvalidClientSecret           = errors.New("Invalid client secret")
	ErrClientIDTaken                 = errors.New("Client ID taken")
	ErrRoleNotFound                  = errors.New("Role not found")
	MinPasswordLength                = 6
	ErrPasswordTooShort              = fmt.Errorf(
		"Password must be at least %d characters long",
		MinPasswordLength,
	)
	ErrUserNotFound                         = errors.New("User not found")
	ErrInvalidUserPassword                  = errors.New("Invalid user password")
	ErrCannotSetEmptyUsername               = errors.New("Cannot set empty username")
	ErrUserPasswordNotSet                   = errors.New("User password not set")
	ErrUsernameTaken                        = errors.New("Username taken")
	ErrInvalidAuthorizationCodeGrantRequest = errors.New("Invalid authorization code request")
	ErrInvalidPasswordGrantRequest          = errors.New("Invalid password grant request")
	ErrInvalidClientCredentialsGrantRequest = errors.New("Invalid client credentials grant request")
	ErrInvalidIntrospectRequest             = errors.New("Invalid introspect request")
	ErrSessonNotStarted                     = errors.New("Session not started")
)

const (
	StatusCodeGenericSuccess            = "200000"
	StatusCodeAccepted                  = "202000"
	StatusCodeBadRequest                = "400000"
	StatusCodeAlreadyRegistered         = "400001"
	StatusCodeUnauthorized              = "401000"
	StatusCodeForbidden                 = "403000"
	StatusCodeNotFound                  = "404000"
	StatusCodeConflict                  = "409000"
	StatusCodeGenericPreconditionFailed = "412000"
	StatusCodeOTPLimitReached           = "412550"
	StatusCodeNoLinkerExist             = "412553"
	StatusCodeInternalError             = "500000"
	StatusCodeFailedSellBatch           = "500100"
	StatusCodeFailedOTP                 = "503000"
	StatusCodeServiceUnavailable        = "503000"
	StatusCodeTimeoutError              = "504000"
	StatusCodeMethodNotAllowed          = "405000"
)

func GetErrorCode(err error) string {
	err = getErrType(err)

	switch err {
	case ErrBadRequest:
		return StatusCodeBadRequest
	case ErrForbiddenResource:
		return StatusCodeForbidden
	case ErrNotFound:
		return StatusCodeNotFound
	case ErrPreConditionFailed:
		return StatusCodeGenericPreconditionFailed
	case ErrInternalServerError:
		return StatusCodeInternalError
	case ErrTimeoutError:
		return StatusCodeTimeoutError
	case ErrUnauthorized:
		return StatusCodeUnauthorized
	case ErrConflict:
		return StatusCodeConflict
	case ErrMethodNotAllowed:
		return StatusCodeMethodNotAllowed
	case ErrInvalidGrantType, ErrInvalidClientIDOrSecret, ErrInvalidRedirectURI, ErrInvalidScope,
		ErrInvalidUsernameOrPassword, ErrTokenHintInvalid, ErrInvalidAuthorizationCodeGrantRequest,
		ErrInvalidPasswordGrantRequest, ErrInvalidClientCredentialsGrantRequest, ErrInvalidIntrospectRequest:
		return StatusCodeBadRequest
	case ErrAuthorizationCodeNotFound, ErrRefreshTokenNotFound, ErrTokenMissing, ErrAccessTokenNotFound:
		return StatusCodeNotFound
	case ErrAuthorizationCodeExpired, ErrRefreshTokenExpired, ErrRequestedScopeCannotBeGreater, ErrAccessTokenExpired:
		return StatusCodeBadRequest
	case ErrClientNotFound, ErrInvalidClientSecret, ErrUserNotFound, ErrRoleNotFound:
		return StatusCodeNotFound
	case ErrClientIDTaken, ErrUsernameTaken:
		return StatusCodeConflict
	case ErrPasswordTooShort, ErrCannotSetEmptyUsername, ErrUserPasswordNotSet, ErrInvalidUserPassword:
		return StatusCodeBadRequest
	case nil:
		return StatusCodeGenericSuccess
	default:
		return StatusCodeInternalError
	}
}

func GetHTTPStatus(code int) string {
	switch code {
	case http.StatusOK:
		return "success"
	case http.StatusCreated:
		return "created"
	case http.StatusAccepted:
		return "accepted"
	case http.StatusNonAuthoritativeInfo:
		return "non authoritative information"
	case http.StatusNoContent:
		return "no content"
	case http.StatusResetContent:
		return "reset content"
	case http.StatusPartialContent:
		return "partial content"
	case http.StatusMultipleChoices:
		return "multiple choices"
	case http.StatusMovedPermanently:
		return "moved permanently"
	case http.StatusFound:
		return "found"
	case http.StatusSeeOther:
		return "see other"
	case http.StatusNotModified:
		return "not modified"
	case http.StatusUseProxy:
		return "use proxy"
	case http.StatusTemporaryRedirect:
		return "temporary redirect"
	case http.StatusPermanentRedirect:
		return "permanent redirect"
	case http.StatusBadRequest:
		return "bad request"
	case http.StatusUnauthorized:
		return "unauthorized"
	case http.StatusPaymentRequired:
		return "payment required"
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusNotFound:
		return "not found"
	case http.StatusMethodNotAllowed:
		return "method not allowed"
	case http.StatusNotAcceptable:
		return "not acceptable"
	case http.StatusProxyAuthRequired:
		return "proxy authentication required"
	case http.StatusRequestTimeout:
		return "request timeout"
	case http.StatusConflict:
		return "conflict"
	case http.StatusGone:
		return "gone"
	case http.StatusLengthRequired:
		return "length required"
	case http.StatusPreconditionFailed:
		return "precondition failed"
	case http.StatusRequestEntityTooLarge:
		return "request entity too large"
	case http.StatusRequestURITooLong:
		return "request URI too long"
	case http.StatusUnsupportedMediaType:
		return "unsupported media type"
	case http.StatusRequestedRangeNotSatisfiable:
		return "requested range not satisfiable"
	case http.StatusExpectationFailed:
		return "expectation failed"
	case http.StatusTeapot:
		return "I'm a teapot"
	case http.StatusMisdirectedRequest:
		return "misdirected request"
	case http.StatusUnprocessableEntity:
		return "unprocessable entity"
	case http.StatusLocked:
		return "locked"
	case http.StatusFailedDependency:
		return "failed dependency"
	case http.StatusUpgradeRequired:
		return "upgrade required"
	case http.StatusPreconditionRequired:
		return "precondition required"
	case http.StatusTooManyRequests:
		return "too many requests"
	case http.StatusRequestHeaderFieldsTooLarge:
		return "request header fields too large"
	case http.StatusUnavailableForLegalReasons:
		return "unavailable for legal reasons"
	case http.StatusInternalServerError:
		return "internal server error"
	case http.StatusNotImplemented:
		return "not implemented"
	case http.StatusBadGateway:
		return "bad gateway"
	case http.StatusServiceUnavailable:
		return "service unavailable"
	case http.StatusGatewayTimeout:
		return "gateway timeout"
	case http.StatusHTTPVersionNotSupported:
		return "HTTP version not supported"
	case http.StatusVariantAlsoNegotiates:
		return "variant also negotiates"
	case http.StatusInsufficientStorage:
		return "insufficient storage"
	case http.StatusLoopDetected:
		return "loop detected"
	case http.StatusNotExtended:
		return "not extended"
	case http.StatusNetworkAuthenticationRequired:
		return "network authentication required"
	default:
		return "undefined"
	}
}

func GetHTTPCode(code string) int {
	s := code[0:3]
	i, _ := strconv.Atoi(s)
	return i
}

type JSONResponse struct {
	Data        interface{}            `json:"data,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Code        string                 `json:"code"`
	StatusCode  int                    `json:"status_code"`
	Status      string                 `json:"status"`
	ErrorString string                 `json:"error,omitempty"`
	Error       error                  `json:"-"`
	RealError   string                 `json:"-"`
	Latency     string                 `json:"latency,omitempty"`
	Log         map[string]interface{} `json:"-"`
	HTMLPage    bool                   `json:"-"`
	Result      interface{}            `json:"result,omitempty"`
}

func NewJSONResponse() *JSONResponse {
	return &JSONResponse{Code: StatusCodeGenericSuccess, StatusCode: GetHTTPCode(StatusCodeGenericSuccess), Status: GetHTTPStatus(http.StatusOK), Log: map[string]interface{}{}}
}

func (r *JSONResponse) SetData(data interface{}) *JSONResponse {
	r.Data = data
	return r
}

func (r *JSONResponse) SetStatus(status string) *JSONResponse {
	r.Status = status
	return r
}

func (r *JSONResponse) SetCode(code string) *JSONResponse {
	r.Code = code
	return r
}

func (r *JSONResponse) SetStatusCode(statusCode int) *JSONResponse {
	r.StatusCode = statusCode
	return r
}

func (r *JSONResponse) SetHTML() *JSONResponse {
	r.HTMLPage = true
	return r
}

func (r *JSONResponse) SetResult(result interface{}) *JSONResponse {
	r.Result = result
	return r
}

func (r *JSONResponse) SetMessage(msg string) *JSONResponse {
	r.Message = msg
	return r
}

func (r *JSONResponse) SetLatency(latency float64) *JSONResponse {
	r.Latency = fmt.Sprintf("%.2f ms", latency)
	return r
}

//func (r *JSONResponse) SetLog(key string, val interface{}) *JSONResponse {
//	_, file, no, _ := runtime.Caller(1)
//	log.Errorln(log.Fields{
//		"code":            r.Code,
//		"err":             val,
//		"function_caller": fmt.Sprintf("file %v line no %v", file, no),
//	}).Errorln("Error API")
//	r.Log[key] = val
//	return r
//}

func getErrType(err error) error {
	switch err.(type) {
	case ErrChain:
		errType := err.(ErrChain).Type
		if errType != nil {
			err = errType
		}
	}
	return err
}

func (r *JSONResponse) SetError(err error, a ...string) *JSONResponse {
	r.Code = GetErrorCode(err)
	// r.SetLog("error", err)
	r.RealError = fmt.Sprintf("%+v", err)
	err = getErrType(err)
	r.Error = err
	r.ErrorString = err.Error()
	r.StatusCode = GetHTTPCode(r.Code)
	r.Status = GetHTTPStatus(r.StatusCode)

	if r.StatusCode == http.StatusInternalServerError {
		r.ErrorString = "Internal Server error"
	}
	if len(a) > 0 {
		r.ErrorString = a[0]
	}
	return r
}

func (r *JSONResponse) GetBody() []byte {
	b, _ := json.Marshal(r)
	return b
}

func (r *JSONResponse) Send(w http.ResponseWriter) {
	if r.HTMLPage {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(r.StatusCode)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.StatusCode)
		err := json.NewEncoder(w).Encode(r)
		if err != nil {
			log.Println("err", err.Error())
		}
	}
}

// APIStatusSuccess for standard request api status success
func (r *JSONResponse) APIStatusSuccess() *JSONResponse {
	r.Code = constant.StatusCode(constant.StatusSuccess)
	r.Message = constant.StatusText(constant.StatusSuccess)
	return r
}

// APIStatusCreated
func (r *JSONResponse) APIStatusCreated() *JSONResponse {
	r.StatusCode = constant.StatusCreated
	r.Code = constant.StatusCode(constant.StatusCreated)
	r.Message = constant.StatusText(constant.StatusCreated)
	return r
}

// APIStatusAccepted
func (r *JSONResponse) APIStatusAccepted() *JSONResponse {
	r.StatusCode = constant.StatusAccepted
	r.Code = constant.StatusCode(constant.StatusAccepted)
	r.Message = constant.StatusText(constant.StatusAccepted)
	return r
}

// APIStatusNoContent
func (r *JSONResponse) APIStatusNoContent() *JSONResponse {
	r.StatusCode = constant.StatusNoContent
	r.Code = constant.StatusCode(constant.StatusNoContent)
	r.Message = constant.StatusText(constant.StatusNoContent)
	return r
}

// APIStatusErrorUnknown
func (r *JSONResponse) APIStatusErrorUnknown() *JSONResponse {
	r.StatusCode = constant.StatusErrorUnknown
	r.Code = constant.StatusCode(constant.StatusErrorUnknown)
	r.Message = constant.StatusText(constant.StatusErrorUnknown)
	return r
}

// APIStatusInvalidAuthentication
func (r *JSONResponse) APIStatusInvalidAuthentication() *JSONResponse {
	r.StatusCode = constant.StatusInvalidAuthentication
	r.Code = constant.StatusCode(constant.StatusInvalidAuthentication)
	r.Message = constant.StatusText(constant.StatusInvalidAuthentication)
	return r
}

// APIStatusUnauthorized
func (r *JSONResponse) APIStatusUnauthorized() *JSONResponse {
	r.StatusCode = constant.StatusUnauthorized
	r.Code = constant.StatusCode(constant.StatusUnauthorized)
	r.Message = constant.StatusText(constant.StatusUnauthorized)
	return r
}

// APIStatusForbidden
func (r *JSONResponse) APIStatusForbidden() *JSONResponse {
	r.StatusCode = constant.StatusForbidden
	r.Code = constant.StatusCode(constant.StatusForbidden)
	r.Message = constant.StatusText(constant.StatusForbidden)
	return r
}

// APIStatusBadRequest
func (r *JSONResponse) APIStatusBadRequest() *JSONResponse {
	r.StatusCode = constant.StatusErrorForm
	r.Code = constant.StatusCode(constant.StatusErrorForm)
	r.Message = constant.StatusText(constant.StatusErrorForm)
	return r
}

// APIStatusNotFound
func (r *JSONResponse) APIStatusNotFound() *JSONResponse {
	r.StatusCode = constant.StatusNotFound
	r.Code = constant.StatusCode(constant.StatusNotFound)
	r.Message = constant.StatusText(constant.StatusNotFound)
	return r
}
