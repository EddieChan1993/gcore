func InitRedisDb() {
	gredis.Init(static.StaticRedisUrl())
}
