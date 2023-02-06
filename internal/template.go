package internal

import "time"

type templateData map[string]interface{}

func generateTemplateData(cfg *Config) (templateData, error) {

	data := make(map[string]interface{})

	var err error
	data["age"], err = age(cfg)
	if err != nil {
		return nil, err
	}

	data["github"], err = getStats(cfg)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func age(cfg *Config) (int, error) {
	born, err := time.Parse(DateFormat, cfg.Born)
	if err != nil {
		return 0, err
	}

	delta := time.Since(born)
	return int(delta.Hours() / 24 / 365), nil
}
