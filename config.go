package nrfiber

const (
	configKeyNoticeErrorEnabled = "NoticeErrorEnabled"
	configKeyNoticeInternalServerErrorEnabled = "NoticeInternalServerErrorEnabled"
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

func ConfigNoticeInternalServerErrorEnabled(enabled bool) *config {
	return &config{
		key:   configKeyNoticeInternalServerErrorEnabled,
		value: enabled,
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

func noticeInternalServerErrorEnabled(configMap map[string]interface{}) bool {
	if val, ok := configMap[configKeyNoticeInternalServerErrorEnabled]; ok {
		if boolVal, ok := val.(bool); ok {
			return boolVal
		}
	}
	return false
}
