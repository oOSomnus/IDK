package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"I_Dev_Kit/internal/entity"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Service 代表与数据库交互的服务
type Service interface {
	// 获取GORM数据库实例
	DB() *gorm.DB
	// Health 返回健康状态信息的映射
	Health() map[string]string
	// Close 终止数据库连接
	Close() error
}

type service struct {
	db *gorm.DB
}

var (
	dburl      = os.Getenv("BLUEPRINT_DB_URL")
	dbInstance *service
)

func New() Service {
	// 重用连接
	if dbInstance != nil {
		return dbInstance
	}

	// 配置GORM日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 打开数据库连接
	db, err := gorm.Open(sqlite.Open(dburl), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移实体
	err = db.AutoMigrate(
		&entity.Project{},
		&entity.Feature{},
		&entity.Decision{},
		&entity.Diagram{},
		&entity.Todo{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// DB 返回GORM数据库实例
func (s *service) DB() *gorm.DB {
	return s.db
}

// Health 检查数据库连接的健康状况
func (s *service) Health() map[string]string {
	stats := make(map[string]string)

	// 检查数据库连接
	sqlDB, err := s.db.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		return stats
	}

	// Ping数据库
	err = sqlDB.Ping()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		return stats
	}

	// 数据库正常，添加更多统计信息
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// 获取数据库统计信息
	dbStats := sqlDB.Stats()
	stats["max_open_connections"] = fmt.Sprintf("%d", dbStats.MaxOpenConnections)
	stats["open_connections"] = fmt.Sprintf("%d", dbStats.OpenConnections)
	stats["in_use"] = fmt.Sprintf("%d", dbStats.InUse)
	stats["idle"] = fmt.Sprintf("%d", dbStats.Idle)
	stats["wait_count"] = fmt.Sprintf("%d", dbStats.WaitCount)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = fmt.Sprintf("%d", dbStats.MaxIdleClosed)
	stats["max_lifetime_closed"] = fmt.Sprintf("%d", dbStats.MaxLifetimeClosed)

	return stats
}

// Close 关闭数据库连接
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
