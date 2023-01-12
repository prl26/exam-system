package questionType

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/28 12:57

 * @Note:

 **/

type QuestionType uint

const (
	JUDGE = QuestionType(1 + iota)
	PROGRAM
	SUPPLY_BLANK
	SINGLE_CHOICE
	MULTIPLE_CHOICE
	Target //靶场题目
)
