package model

type MessageModel struct{
	UserDetail UserDetail `json:"userDetail"`
	CompanyDetail CompanyDetail `json:"companyDetail"`
	CredentialDetail CredentialDetail `json:"credentialDetail"`
}

type UserDetail struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}

type CompanyDetail struct{
	CompanyName string `json:"companyName"`
}

type CredentialDetail struct{
	Username string `json:"userName"`
	Password string `json:"password"`
}