package v1

type TenantInfo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Phone   string `json:"phone"`
}

type UserInfo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	TenantId   string `json:"tenantId"`
	Department string `json:"department"`
}

type TenantRequestVCodeReq struct {
	PhoneNum string `json:"phonenum"`
}

type TenantRequestVCodeResp struct {
	Message string `json:"message"`
}

type TenantCreateReq struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	PhoneNum string `json:"phonenum"`
	Email    string `json:"email"`
	VCode    string `json:"vcode"`
	Password string `json:"password"`
}

type TenantCreateResp struct {
	Secret  string `json:"secret"`
	Message string `json:"message"`
}

type UserCreateReq struct {
	Name       string `json:"name"`
	Tenant     string `json:"tenant"`
	Company    string `json:"company"`
	Department string `json:"department"`
	PhoneNum   string `json:"phonenum"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type UserCreateResp struct {
	Secret  string `json:"secret"`
	Message string `json:"message"`
}

type ClusterCreateReq struct {
	Name      string `json:"name"`
	Tenant    string `json:"name"`
	Signature string `json:"signature"`
}

type ClusterCreateResp struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
