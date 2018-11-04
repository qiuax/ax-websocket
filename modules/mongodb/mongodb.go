package mongodb

func Close() {
	if SessionLogs != nil {
		SessionLogs.Close()
	}
}
