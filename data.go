package main

// 此实体为 etcd中存储的实体
// 这里我坐一些改变,LogPath支持以,分隔的数据
type logConfig struct {
	Topic    string `json:"topic"`
	LogPath  string `json:"log_path"`
	Service  string `json:"service"`
	SendRate int    `json:"send_rate"`
}
