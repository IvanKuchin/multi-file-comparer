package config_reader

import "log"

func GetConfig() (*[]GroupConfig, error) {
	gcs := make([]GroupConfig, 0)

	config_reader, err := NewConfig("yaml")
	if err != nil {
		log.Print("ERROR: ", err)
		return nil, err
	}

	i := 0
	for {
		gc, err := config_reader.GetGroupConfig(i)
		if err != nil {
			log.Print("ERROR: ", err)
			return nil, err
		}

		if gc == nil {
			break
		}

		gcs = append(gcs, *gc)
		i++
	}

	return &gcs, nil
}
