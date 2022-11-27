package request

import "github.com/prl26/exam-system/server/utils"

var BaseVerify = utils.Rules{
	"ProblemType": {utils.NotEmpty()},
	"CanPractice": {utils.NotEmpty()},
	"CanExam":     {utils.NotEmpty()},
	"Title":       {utils.NotEmpty()},
	//"Describe":    {utils.NotEmpty()},
}
