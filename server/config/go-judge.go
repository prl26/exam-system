package config

/**

 * @Author: AloneAtWar

 * @Date:   2022/10/25 10:04

 * @Note:

 **/

type GoJudge struct {
	Enable    bool      `mapstructure:"enable" json:"enable" yaml:"enable"`
	Address   string    `mapstructure:"address" json:"address" yaml:"address"`
	CLanguage CLanguage `mapstructure:"cLanguage" json:"cLanguage" yaml:"cLanguage"`
}

type CLanguage struct {
	Enable                            bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	GCC_PATH                          string `mapstructure:"gcc_path" json:"gcc_path" yaml:"gcc_path"`
	DEFAULT_COMPILE_CPU_TIME_LIMIT    uint64 `mapstructure:"default_compile_cpu_time_limit" json:"default_compile_cpu_time_limit" yaml:"default_compile_cpu_time_limit"`
	DEFAULT_COMPILE_MEMORY_TIME_LIMIT uint64 `mapstructure:"default_compile_memory_time_limit" json:"default_compile_memory_time_limit" yaml:"default_compile_memory_time_limit"`
	DEFAULT_JUDGE_CPU_TIME_LIMIT      uint64 `mapstructure:"default_judge_cpu_time_limit" json:"default_judge_cpu_time_limit" yaml:"default_judge_cpu_time_limit"`
	DEFAULT_JUDGE_MEMORY_LIMIT        uint64 `mapstructure:"default_judge_memory_limit" json:"default_judge_memory_limit" yaml:"default_judge_memory_limit"`
}
