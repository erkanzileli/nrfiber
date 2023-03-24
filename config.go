package nrfiber

const (
	configKeyNoticeErrorEnabled = "NoticeErrorEnabled"
	configKeyStatusCodeIgnored  = "StatusCodeIgnored"
)

type config struct {
	key   string
	value interface{}
}

func ConfigNoticeErrorEnabled(enabled bool) *config {
	return &config{
		key:   configKeyNoticeErrorEnabled,
		value: enabled,
	}
}

func ConfigStatusCodeIgnored(statusCode []int) *config {
	return &config{
		key:   configKeyNoticeErrorEnabled,
		value: statusCode,
	}
}

func createConfigMap(configs ...*config) map[string]interface{} {
	configMap := make(map[string]interface{}, len(configs))
	for _, c := range configs {
		configMap[c.key] = c.value
	}
	return configMap
}

func noticeErrorEnabled(configMap map[string]interface{}) bool {
	if val, ok := configMap[configKeyNoticeErrorEnabled]; ok {
		if boolVal, ok := val.(bool); ok {
			return boolVal
		}
	}
	return false
}

func statusCodeIgnored(configMap map[string]interface{}) []int {
	if val, ok := configMap[configKeyStatusCodeIgnored]; ok {
		if boolVal, ok := val.([]int); ok {
			return boolVal
		}
	}
	return []int{}
}
