package config_reader

type StaticConfig struct {
}

func (sc *StaticConfig) GetGroupConfig(idx int) (*GroupConfig, error) {
	switch idx {
	case 0:
		return &GroupConfig{
			Rootdir:    "/home/ikuchin/src",
			Subfolders: []string{"dev.connme.ru", "dev.bestbounty.ru"},
			Files: []string{
				"CMakeLists.txt",
				"po/CMakeLists.txt",
				"supplemental/git/CMakeLists.txt",
				"templates/CMakeLists.txt",
				".github/workflows/release-deployment.yml",
				".github/workflows/qa.yml",
				".github/workflows/sast.yml",
				".github/workflows/spell_check.yml",
				".github/workflows/super_linter.yml",
				".github/workflows/release-drafter.yml",
				"supplemental/docker/prod/Dockerfile",
				"supplemental/docker/qa/Dockerfile",
				"supplemental/apache2/conf-available/mod_expires.conf",
				"supplemental/apache2/sites-available/099-images.conf",
				"supplemental/apache2/sites-available/099-www.conf",
				"supplemental/ssmtp/CMakeLists.txt",
				"supplemental/ssmtp/ssmtp.conf.in",
				"supplemental/libpdf-font/CMakeLists.txt",
				"supplemental/config/CMakeLists.txt",
				"supplemental/archive/CMakeLists.txt",
				"supplemental/maxmind/CMakeLists.txt",
				"supplemental/git/CMakeLists.txt",
				"supplemental/cron/CMakeLists.txt",
			},
		}, nil
	case 1:
		return &GroupConfig{
			Rootdir:    "/home/ikuchin/src",
			Subfolders: []string{"dev.connme.ru-admin", "dev.bestbounty.ru-admin"},
			Files: []string{
				"CMakeLists.txt",
				"templates/CMakeLists.txt",
				"src/admin/CMakeLists.txt",
				"supplemental/docker/prod/Dockerfile",
				"supplemental/docker/qa/Dockerfile",
				".github/workflows/release-deployment.yml",
				"supplemental/apache2/conf-available/mod_expires.conf",
				"supplemental/apache2/sites-available/099-images.conf",
				"supplemental/apache2/sites-available/099-www.conf",
				"supplemental/config/CMakeLists.txt",
				"supplemental/git/CMakeLists.txt",
			},
		}, nil
	case 2:
		return &GroupConfig{
			Rootdir:    "/home/http",
			Subfolders: []string{"dev.connme.ru", "dev.bestbounty.ru"},
			Files: []string{
				"supplemental/Docker/prod/Dockerfile",
				".github/workflows/release-deployment.yml",
				".github/workflows/spell-checker.yml",
				".github/workflows/super-linter.yml",
				"supplemental/apache2/conf-available/mod_expires.conf",
				"supplemental/apache2/sites-available/099-images.conf",
				"supplemental/apache2/sites-available/099-www.conf",
				"user_account_properties.js",
				"user_account_properties.htmlt",
			},
		}, nil
	}
	return nil, nil
}

func NewStaticConfig() (Config, error) {
	return &StaticConfig{}, nil
}
