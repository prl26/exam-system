package problemType

type ProblemType uint

const (
	EASY = ProblemType(1 + iota)
	MEDIUM
	HARD
)

const (
	easy_name   = "EASY"
	medium_name = "MEDIUM"
	hard_name   = "HARD"
)

var problemTypeNameToType map[string]ProblemType = map[string]ProblemType{
	easy_name:   EASY,
	medium_name: MEDIUM,
	hard_name:   HARD,
}

var problemTypeToName = map[ProblemType]string{
	EASY:   easy_name,
	MEDIUM: medium_name,
	HARD:   hard_name,
}
