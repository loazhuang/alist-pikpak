package homecloud

const (
	MetaPersonal    string = "personal"
	MetaFamily      string = "family"
	MetaPersonalNew string = "personal_new"
)

type BaseResp struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type FrontResp struct {
	Ret    int    `json:"ret"`
	Reason string `json:"reason"`
}

type Catalog struct {
	CatalogID   string `json:"catalogID"`
	CatalogName string `json:"catalogName"`
	//CatalogType     int         `json:"catalogType"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	//IsShared        bool        `json:"isShared"`
	//CatalogLevel    int         `json:"catalogLevel"`
	//ShareDoneeCount int         `json:"shareDoneeCount"`
	//OpenType        int         `json:"openType"`
	//ParentCatalogID string      `json:"parentCatalogId"`
	//DirEtag         int         `json:"dirEtag"`
	//Tombstoned      int         `json:"tombstoned"`
	//ProxyID         interface{} `json:"proxyID"`
	//Moved           int         `json:"moved"`
	//IsFixedDir      int         `json:"isFixedDir"`
	//IsSynced        interface{} `json:"isSynced"`
	//Owner           string      `json:"owner"`
	//Modifier        interface{} `json:"modifier"`
	//Path            string      `json:"path"`
	//ShareType       int         `json:"shareType"`
	//SoftLink        interface{} `json:"softLink"`
	//ExtProp1        interface{} `json:"extProp1"`
	//ExtProp2        interface{} `json:"extProp2"`
	//ExtProp3        interface{} `json:"extProp3"`
	//ExtProp4        interface{} `json:"extProp4"`
	//ExtProp5        interface{} `json:"extProp5"`
	//ETagOprType     int         `json:"ETagOprType"`
}

type Content struct {
	ContentID   string `json:"contentID"`
	ContentName string `json:"contentName"`
	//ContentSuffix   string      `json:"contentSuffix"`
	ContentSize int64 `json:"contentSize"`
	//ContentDesc     string      `json:"contentDesc"`
	//ContentType     int         `json:"contentType"`
	//ContentOrigin   int         `json:"contentOrigin"`
	UpdateTime string `json:"updateTime"`
	//CommentCount    int         `json:"commentCount"`
	ThumbnailURL string `json:"thumbnailURL"`
	//BigthumbnailURL string      `json:"bigthumbnailURL"`
	//PresentURL      string      `json:"presentURL"`
	//PresentLURL     string      `json:"presentLURL"`
	//PresentHURL     string      `json:"presentHURL"`
	//ContentTAGList  interface{} `json:"contentTAGList"`
	//ShareDoneeCount int         `json:"shareDoneeCount"`
	//Safestate       int         `json:"safestate"`
	//Transferstate   int         `json:"transferstate"`
	//IsFocusContent  int         `json:"isFocusContent"`
	//UpdateShareTime interface{} `json:"updateShareTime"`
	//UploadTime      string      `json:"uploadTime"`
	//OpenType        int         `json:"openType"`
	//AuditResult     int         `json:"auditResult"`
	//ParentCatalogID string      `json:"parentCatalogId"`
	//Channel         string      `json:"channel"`
	//GeoLocFlag      string      `json:"geoLocFlag"`
	Digest string `json:"digest"`
	//Version         string      `json:"version"`
	//FileEtag        string      `json:"fileEtag"`
	//FileVersion     string      `json:"fileVersion"`
	//Tombstoned      int         `json:"tombstoned"`
	//ProxyID         string      `json:"proxyID"`
	//Moved           int         `json:"moved"`
	//MidthumbnailURL string      `json:"midthumbnailURL"`
	//Owner           string      `json:"owner"`
	//Modifier        string      `json:"modifier"`
	//ShareType       int         `json:"shareType"`
	//ExtInfo         struct {
	//	Uploader string `json:"uploader"`
	//	Address  string `json:"address"`
	//} `json:"extInfo"`
	//Exif struct {
	//	CreateTime    string      `json:"createTime"`
	//	Longitude     interface{} `json:"longitude"`
	//	Latitude      interface{} `json:"latitude"`
	//	LocalSaveTime interface{} `json:"localSaveTime"`
	//} `json:"exif"`
	//CollectionFlag interface{} `json:"collectionFlag"`
	//TreeInfo       interface{} `json:"treeInfo"`
	//IsShared       bool        `json:"isShared"`
	//ETagOprType    int         `json:"ETagOprType"`
}

type GetDiskResp struct {
	BaseResp
	Data struct {
		Result struct {
			ResultCode string      `json:"resultCode"`
			ResultDesc interface{} `json:"resultDesc"`
		} `json:"result"`
		GetDiskResult struct {
			ParentCatalogID string    `json:"parentCatalogID"`
			NodeCount       int       `json:"nodeCount"`
			CatalogList     []Catalog `json:"catalogList"`
			ContentList     []Content `json:"contentList"`
			IsCompleted     int       `json:"isCompleted"`
		} `json:"getDiskResult"`
	} `json:"data"`
}

type UploadResp struct {
	BaseResp
	Data struct {
		Result struct {
			ResultCode string      `json:"resultCode"`
			ResultDesc interface{} `json:"resultDesc"`
		} `json:"result"`
		UploadResult struct {
			UploadTaskID     string `json:"uploadTaskID"`
			RedirectionURL   string `json:"redirectionUrl"`
			NewContentIDList []struct {
				ContentID     string `json:"contentID"`
				ContentName   string `json:"contentName"`
				IsNeedUpload  string `json:"isNeedUpload"`
				FileEtag      int64  `json:"fileEtag"`
				FileVersion   int64  `json:"fileVersion"`
				OverridenFlag int    `json:"overridenFlag"`
			} `json:"newContentIDList"`
			CatalogIDList interface{} `json:"catalogIDList"`
			IsSlice       interface{} `json:"isSlice"`
		} `json:"uploadResult"`
	} `json:"data"`
}

type CloudContent struct {
	ContentID string `json:"contentID"`
	//Modifier         string      `json:"modifier"`
	//Nickname         string      `json:"nickname"`
	//CloudNickName    string      `json:"cloudNickName"`
	ContentName string `json:"contentName"`
	//ContentType      int         `json:"contentType"`
	//ContentSuffix    string      `json:"contentSuffix"`
	ContentSize int64 `json:"contentSize"`
	//ContentDesc      string      `json:"contentDesc"`
	CreateTime string `json:"createTime"`
	//Shottime         interface{} `json:"shottime"`
	LastUpdateTime string `json:"lastUpdateTime"`
	ThumbnailURL   string `json:"thumbnailURL"`
	//MidthumbnailURL  string      `json:"midthumbnailURL"`
	//BigthumbnailURL  string      `json:"bigthumbnailURL"`
	//PresentURL       string      `json:"presentURL"`
	//PresentLURL      string      `json:"presentLURL"`
	//PresentHURL      string      `json:"presentHURL"`
	//ParentCatalogID  string      `json:"parentCatalogID"`
	//Uploader         string      `json:"uploader"`
	//UploaderNickName string      `json:"uploaderNickName"`
	//TreeInfo         interface{} `json:"treeInfo"`
	//UpdateTime       interface{} `json:"updateTime"`
	//ExtInfo          struct {
	//	Uploader string `json:"uploader"`
	//} `json:"extInfo"`
	//EtagOprType interface{} `json:"etagOprType"`
}

type CloudCatalog struct {
	CatalogID   string `json:"catalogID"`
	CatalogName string `json:"catalogName"`
	//CloudID         string `json:"cloudID"`
	CreateTime     string `json:"createTime"`
	LastUpdateTime string `json:"lastUpdateTime"`
	//Creator         string `json:"creator"`
	//CreatorNickname string `json:"creatorNickname"`
}

type QueryContentListResp struct {
	FrontResp
	Data struct {
		Ret       int        `json:"ret"`
		Reason    string     `json:"reason"`
		Total     string     `json:"total"`
		FileInfos []FileInfo `json:"fileInfos"`
	} `json:"data"`
}

type FileInfo struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           int64  `json:"type"`
	Category       int64  `json:"category"`
	Md5            string `json:"md5"`
	Size           string `json:"size"`
	ParentId       string `json:"parentId"`
	GroupID        string `json:"groupId"`
	UserID         string `json:"userId"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
	Ctag           string `json:"ctag"`
	ParentCategory int64  `json:"parentCategory"`
}

type PersonalThumbnail struct {
	Style string `json:"style"`
	Url   string `json:"url"`
}

type PersonalFileItem struct {
	FileId     string              `json:"fileId"`
	Name       string              `json:"name"`
	Size       int64               `json:"size"`
	Type       string              `json:"type"`
	CreatedAt  string              `json:"createdAt"`
	UpdatedAt  string              `json:"updatedAt"`
	Thumbnails []PersonalThumbnail `json:"thumbnailUrls"`
}

type PersonalListResp struct {
	BaseResp
	Data struct {
		Items          []PersonalFileItem `json:"items"`
		NextPageCursor string             `json:"nextPageCursor"`
	}
}

type PersonalPartInfo struct {
	PartNumber int    `json:"partNumber"`
	UploadUrl  string `json:"uploadUrl"`
}

type PersonalUploadResp struct {
	FrontResp
	Data struct {
		Ret         int    `json:"ret"`
		Reason      string `json:"reason"`
		UploadId    string `json:"uploadId"`
		UploadState int    `json:"uploadState"`
		ExpireTime  string `json:"expireTime"`
		FileId      string `json:"fileId"`
		UploadUrl   string `json:"uploadUrl"`
	}
}

type RefreshTokenResp struct {
	FrontResp
	Data struct {
		Ret          int    `json:"ret"`
		ExpiresIn    int    `json:"expiresIn"`
		License      string `json:"license"`
		Scope        string `json:"scope"`
		UserType     int    `json:"userType"`
		AccessToken  string `json:"accessToken"`
		TokenType    string `json:"tokenType"`
		UserID       string `json:"userId"`
		RefreshToken string `json:"refreshToken"`
	}
}
