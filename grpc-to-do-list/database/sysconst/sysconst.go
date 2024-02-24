package sysconst

import (
	"fmt"
	"grpc-to-do-list/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success                             = iota // 請求成功
	InvalidId                           = 4000 // 無效ID
	WrongParameterFormat                = 5000 // 參數格式錯誤
	AccountOrPasswordError              = 5001 // 帳號密碼錯誤
	TokenHasExpired                     = 5002 // Token 過期
	Unauthorized                        = 5003 // 未經授權
	PermissionDenied                    = 5004 // 沒有權限
	InvalidVerificationCode             = 5005 // 無效驗證碼
	InvalidTokenClaims                  = 5006 // Token解析無效
	InvalidMemberID                     = 5007 // 無效會員ID
	InvalidArticleID                    = 5008 // 無效文章ID
	InvalidArticleCategoryID            = 5009 // 無效文章類別ID
	InvalidCompetitionAndTypeID         = 5010 // 無效賽事類別ID與運動類型
	InvalidTokenMemberID                = 5011 // 無效的Token會員ID
	InvalidPublisherMemberID            = 5012 // 無效的發佈者ID
	InvalidApproverMemberID             = 5013 // 無效的審核者ID
	DuplicateMemberAccount              = 5014 // 帳號重複
	InvalidTag                          = 5015 // 無效的標籤
	InvalidURL                          = 5016 // 無效的網址
	InvalidTemporaryToken               = 5017 // 無效的暫時驗證 token
	DuplicateMemberUserName             = 5018 // 使用者名稱重複
	InvalidManagerId                    = 5019 // 無效管理者ID
	InvalidCategoryID                   = 5020 // 無效選單ID
	InvalidRoleID                       = 5021 // 無效角色ID
	InvalidContentFilterId              = 5022 // 無效內容過濾 ID
	InvalidAnalystGradeRequired         = 5023 // 無效分析師等級需求
	InvalidAnalystGradeID               = 5024 // 無效的分析師等級 ID
	InsufficientBalance                 = 5025 // 餘額不足
	InsufficientPoints                  = 5026 // 點數不足
	DuplicatedRoleName                  = 5027 // 角色名稱重複
	InvalidAnalystID                    = 5028 // 無效的分析師 ID
	InvalidOperation                    = 5029 // 無效操作
	IsDisableAccount                    = 5030 // 帳號停用中
	BalanceOverflow                     = 5031 // 金額溢出(超出範圍限制)
	IsBroadcastedNotification           = 5032 // 通知已播送
	InvalidImage                        = 5033 // 無效的圖片
	InvalidTime                         = 5034 // 無效的時間
	InvalidLink                         = 5035 // 無效的連結
	DataOutOfSync                       = 5036 // 數據不同步
	IsExceededMatchTime                 = 5037 // 已超過賽事時間
	UpdateError                         = 8001 // 更新失敗
	InsertError                         = 8002 // 新增失敗
	RepeatError                         = 8003 // 重複新增
	DelError                            = 8004 // 刪除失敗
	QueryError                          = 8005 // 查詢失敗
	InternalSystemError                 = 9000 // 系統內部錯誤
	DataAleadyExists                    = 9001 // 資料已存在
	DataNotExists                       = 9002 // 資料不存在
	UploadVideoInitError                = 9050 // 上傳檔案初始有誤
	UploadVideoToServerError            = 9051 // 上傳檔案失敗
	QuantityLimitExceededOfUploadVideos = 9052 // 超出上傳影片數量限制
	QuantityLimitExceededOfUploadImages = 9053 // 超出上傳圖片數量限制
	ArticleAlreadyAudited               = 9100 // 文章已經審核(不可再次審核)
	AnalystAlreadyAudited               = 9101 // 分析師已經審核(不可再次審核)
	AnalystDrawRequestAlreadyAudited    = 9102 // 分析師提領已經審核(不可再次審核)
	AnalystGradeAlreadyUsed             = 9200 // 此等級已有會員正在使用中
	AccountAlreadyExists                = 9201 // 帳號已存在
	NicknameAlreadyExists               = 9202 // 暱稱已存在
	EmailAlreadyExists                  = 9203 // 郵件已存在
	MaintainNotInsert                   = 9301 //維修不能新增
	MaintainTimeError                   = 9302 //維修時間錯誤
	NotMaintaining                      = 9303 //沒有維修在進行
)

func GetResultMessage(rs int) string {
	var result = ""
	switch rs {
	case Success:
		result = "請求成功"
	case WrongParameterFormat:
		result = "參數格式錯誤"
	case AccountOrPasswordError:
		result = "帳號密碼錯誤"
	case TokenHasExpired:
		result = "令牌過期"
	case Unauthorized:
		result = "未經授權"
	case PermissionDenied:
		result = "沒有權限"
	case InvalidOperation:
		result = "無效操作"
	case IsDisableAccount:
		result = "帳號停用中"
	case DataOutOfSync:
		result = "資料不同步"
	case BalanceOverflow:
		result = "點數加總超過最大上限"
	case IsBroadcastedNotification:
		result = "通知已播送"
	case InvalidImage:
		result = "無效圖片"
	case InvalidTime:
		result = "無效時間"
	case InvalidLink:
		result = "無效的連結"
	case IsExceededMatchTime:
		result = "已超過賽事時間"
	case InvalidVerificationCode:
		result = "無效驗證碼"
	case InvalidTokenClaims:
		result = "Token解析無效"
	case InvalidMemberID:
		result = "無效會員ID"
	case InvalidArticleID:
		result = "無效文章ID"
	case InvalidArticleCategoryID:
		result = "無效文章類別ID"
	case InvalidCompetitionAndTypeID:
		result = "無效賽事類別ID與運動類型"
	case InvalidTokenMemberID:
		result = "無效的Token會員ID"
	case InvalidPublisherMemberID:
		result = "無效的發佈者ID"
	case InvalidApproverMemberID:
		result = "無效的審核者ID"
	case InvalidManagerId:
		result = "無效的管理者 ID"
	case InvalidAnalystGradeID:
		result = "無效的分析師等級"
	case InvalidAnalystID:
		result = "無效的分析師 ID"
	case InsufficientBalance:
		result = "餘額不足"
	case InsufficientPoints:
		result = "點數不足"
	case DuplicatedRoleName:
		result = "角色名稱重複"
	case DuplicateMemberAccount:
		result = "帳號重複"
	case DuplicateMemberUserName:
		result = "會員名稱重複"
	case InvalidCategoryID:
		result = "無效選單ID"
	case InvalidRoleID:
		result = "無效角色ID"
	case InvalidContentFilterId:
		result = "無效過濾內容 ID"
	case InvalidTag:
		result = "無效的標籤"
	case InvalidURL:
		result = "無效的網址"
	case InvalidTemporaryToken:
		result = "無效的認證"
	case InvalidAnalystGradeRequired:
		result = "無效的分析師等級需求"
	case InternalSystemError:
		result = "系統內部錯誤"
	case DataAleadyExists:
		result = "資料已存在"
	case DataNotExists:
		result = "資料不存在"
	case UploadVideoInitError:
		result = "上傳檔案初始有誤"
	case UploadVideoToServerError:
		result = "上傳檔案失敗"
	case QuantityLimitExceededOfUploadVideos:
		result = "超出上傳影片數量限制"
	case QuantityLimitExceededOfUploadImages:
		result = "超出上傳圖片數量限制"
	case UpdateError:
		result = "更新失敗"
	case InsertError:
		result = "新增失敗"
	case RepeatError:
		result = "重複新增"
	case DelError:
		result = "刪除失敗"
	case QueryError:
		result = "查詢失敗"
	case ArticleAlreadyAudited:
		result = "此文章已審核"
	case AnalystAlreadyAudited:
		result = "此分析師已審核"
	case AnalystDrawRequestAlreadyAudited:
		result = "此分析師提領已經審核"
	case AnalystGradeAlreadyUsed:
		result = "此等級已有會員正在使用中"
	case AccountAlreadyExists:
		result = "該帳號已存在"
	case NicknameAlreadyExists:
		result = "該暱稱已存在"
	case EmailAlreadyExists:
		result = "該郵件已存在"
	case MaintainNotInsert:
		result = "維修不能新增"
	case MaintainTimeError:
		result = "維修時間錯誤"
	case NotMaintaining:
		result = "沒有維修在進行"
	}

	return result
}

func GetResultHTTPCode(code int) int {
	switch code {
	case Success:
		return http.StatusOK
	case WrongParameterFormat:
		return http.StatusBadRequest
	case AccountOrPasswordError:
		return http.StatusBadRequest
	case TokenHasExpired:
		return http.StatusUnauthorized
	case Unauthorized:
		return http.StatusUnauthorized
	case PermissionDenied:
		return http.StatusUnauthorized
	case InvalidOperation:
		return http.StatusUnauthorized
	case IsDisableAccount:
		return http.StatusUnauthorized
	case BalanceOverflow:
		return http.StatusBadRequest
	case IsBroadcastedNotification:
		return http.StatusBadRequest
	case InvalidImage:
		return http.StatusBadRequest
	case InvalidTime:
		return http.StatusBadRequest
	case InvalidLink:
		return http.StatusBadRequest
	case InvalidVerificationCode:
		return http.StatusBadRequest
	case InvalidTokenClaims:
		return http.StatusUnauthorized
	case InvalidMemberID:
		return http.StatusBadRequest
	case InvalidArticleID:
		return http.StatusBadRequest
	case InvalidArticleCategoryID:
		return http.StatusBadRequest
	case InvalidCompetitionAndTypeID:
		return http.StatusBadRequest
	case InvalidTokenMemberID:
		return http.StatusNotFound
	case InvalidPublisherMemberID:
		return http.StatusBadRequest
	case InvalidApproverMemberID:
		return http.StatusBadRequest
	case InvalidManagerId:
		return http.StatusBadRequest
	case InvalidAnalystGradeID:
		return http.StatusBadRequest
	case InvalidAnalystID:
		return http.StatusBadRequest
	case InsufficientBalance:
		return http.StatusBadRequest
	case InsufficientPoints:
		return http.StatusBadRequest
	case DuplicatedRoleName:
		return http.StatusBadRequest
	case DuplicateMemberAccount:
		return http.StatusBadRequest
	case DuplicateMemberUserName:
		return http.StatusBadRequest
	case InvalidCategoryID:
		return http.StatusBadRequest
	case InvalidRoleID:
		return http.StatusBadRequest
	case InvalidContentFilterId:
		return http.StatusBadRequest
	case InvalidTag:
		return http.StatusBadRequest
	case InvalidURL:
		return http.StatusBadRequest
	case InvalidTemporaryToken:
		return http.StatusBadRequest
	case InvalidAnalystGradeRequired:
		return http.StatusBadRequest
	case InternalSystemError:
		return http.StatusBadRequest
	case DataAleadyExists:
		return http.StatusBadRequest
	case DataNotExists:
		return http.StatusBadRequest
	case UploadVideoInitError:
		return http.StatusBadRequest
	case UploadVideoToServerError:
		return http.StatusBadRequest
	case QuantityLimitExceededOfUploadVideos:
		return http.StatusBadRequest
	case QuantityLimitExceededOfUploadImages:
		return http.StatusBadRequest
	case UpdateError:
		return http.StatusBadRequest
	case InsertError:
		return http.StatusBadRequest
	case RepeatError:
		return http.StatusBadRequest
	case DelError:
		return http.StatusBadRequest
	case QueryError:
		return http.StatusBadRequest
	case ArticleAlreadyAudited:
		return http.StatusBadRequest
	case AnalystAlreadyAudited:
		return http.StatusBadRequest
	case AnalystGradeAlreadyUsed:
		return http.StatusBadRequest
	case AccountAlreadyExists:
		return http.StatusBadRequest
	case NicknameAlreadyExists:
		return http.StatusBadRequest
	case EmailAlreadyExists:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func Response(c *gin.Context, code int, result dtos.ResultDto, msg interface{}) {
	if code != Success {
		result.Error = []dtos.ErrorDto{{Message: fmt.Sprintf("%v %v", GetResultMessage(code), msg)}}
	}
	c.JSON(GetResultHTTPCode(code), result)
}
