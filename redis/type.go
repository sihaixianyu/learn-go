package redis

func StringExample() {
	err := client.Set(ctx, "name", "tingxuan", 0).Err()
	if err != nil {
		panic(err)
	}
}
