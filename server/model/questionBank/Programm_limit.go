package questionBank

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/25 2:16

 * @Note:

 **/

type ProgrammLimit struct {
	StrictMemoryLimit *int `json:"strictMemoryLimit" form:"strictMemoryLimit" gorm:"column:strict_memory_limit;comment:;"`
	MemoryLimit       *int `json:"memoryLimit" form:"memoryLimit" gorm:"column:memory_limit"`
	CpuLimit          *int `json:"cpuLimit" form:"cpuLimit" gorm:"column:cpu_limit;comment:;"`
	ClockLimit        *int `json:"clockLimit" form:"clockLimit" gorm:"column:clock_limit;comment:;"`
	StackLimit        *int `json:"stackLimit" form:"stackLimit" gorm:"column:stack_limit;comment:;"`
	ProcLimit         *int `json:"procLimit" form:"procLimit" gorm:"column:proc_limit;comment:;"`
	CpuRateLimit      *int `json:"cpuRateLimit" form:"cpuRateLimit" gorm:"column:cpu_rate_limit;comment:;"`
	CpuSetLimit       *int `json:"cpuSetLimit" form:"cpuSetLimit" gorm:"column:cpu_set_limit;comment:;"`
}
