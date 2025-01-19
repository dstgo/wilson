package conf

import "time"

// Config 资源服务配置结构
type Config struct {
	// DefaultMaxSize 默认最大文件大小（单位：MB）
	DefaultMaxSize uint32
	// DefaultAcceptTypes 默认允许的文件类型列表
	DefaultAcceptTypes []string
	// ChunkSize 文件分块大小（单位：MB）
	ChunkSize uint32
	// Export 导出相关配置
	Export Export
	// Storage 存储相关配置
	Storage Storage
}

// Export 导出配置结构
type Export struct {
	// ServerURL 导出服务器URL
	ServerURL string
	// LocalDir 导出文件本地存储目录
	LocalDir string
	// Expire 导出文件过期时间
	Expire time.Duration
}

// Storage 存储配置结构
type Storage struct {
	// Type 存储类型（如：local, s3, oss等）
	Type string
	// Endpoint 存储服务端点
	Endpoint string
	// Id 存储服务访问ID
	Id string
	// Secret 存储服务访问密钥
	Secret string
	// Bucket 存储桶名称
	Bucket string
	// Region 存储区域
	Region string
	// LocalDir 本地存储目录（仅用于本地存储类型）
	LocalDir string
	// ServerURL 存储服务访问URL
	ServerURL string
	// TemporaryExpire 临时文件过期时间
	TemporaryExpire time.Duration
}
