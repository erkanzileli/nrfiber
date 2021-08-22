package nrfiber

const (
	configKeyNoticeErrorEnabled = "NoticeErrorEnabled"
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

func createConfigMap(configs ...*config) map[string]interface{} {
	configMap := make(map[string]interface{}, len(configs))
	for _, c := range configs {
		configMap[c.key] = c.value
	}
	return configMap
}
