package config

type DatabaseConfiguration struct {
	Mongo_URL string `required:"true"`
}

type AuthorizationConfiguration struct {
	JWT_SECRET         string `required:"true"`
	ACCESS_TTL_MINUTES int
	REFRESH_TTL_DAYS   int
	BCRYPT_COST        int
}
