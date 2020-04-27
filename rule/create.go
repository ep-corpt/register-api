package rule

import validation "github.com/go-ozzo/ozzo-validation"

var UserDetail = []validation.Rule{validation.Required}
var FirstName = []validation.Rule{validation.Required, validation.Length(2, 30)}
var LastName =[]validation.Rule{validation.Required, validation.Length(2, 30)}
var Email = []validation.Rule{validation.Required, validation.Length(5, 30)}


var CompanyDetail = []validation.Rule{validation.Required}
var CompanyName = []validation.Rule{validation.Required, validation.Length(5, 30)}


var CredentialDetail = []validation.Rule{validation.Required}
var UserName = []validation.Rule{validation.Required, validation.Length(10, 20)}
var Password = []validation.Rule{validation.Required, validation.Length(10, 20)}

