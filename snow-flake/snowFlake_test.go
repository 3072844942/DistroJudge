package snow_flake

import (
	"strconv"
	"testing"
	"time"
)

func TestGetSnowFlak(t *testing.T) {
	workerId := int64(1)
	datacenterId := int64(1)

	_, err := GetSnowFlak(workerId, datacenterId)
	if err != nil {
		t.Errorf("Error creating SnowFlake instance: %v", err)
	}
}

func TestSnowFlake_GetWorkerId(t *testing.T) {
	workerId := int64(1)
	datacenterId := int64(1)

	snowFlake, err := GetSnowFlak(workerId, datacenterId)
	if err != nil {
		t.Errorf("Error creating SnowFlake instance: %v", err)
	}

	// 测试获取机器ID和机房ID的方法
	if snowFlake.GetWorkerId() != workerId {
		t.Errorf("Expected workerId %d, got %d", workerId, snowFlake.GetWorkerId())
	}
}

func TestSnowFlake_GetDatacenterId(t *testing.T) {
	workerId := int64(1)
	datacenterId := int64(1)

	snowFlake, err := GetSnowFlak(workerId, datacenterId)
	if err != nil {
		t.Errorf("Error creating SnowFlake instance: %v", err)
	}

	if snowFlake.GetDatacenterId() != datacenterId {
		t.Errorf("Expected datacenterId %d, got %d", datacenterId, snowFlake.GetDatacenterId())
	}
}

func TestSnowFlake_GetLastTimestamp(t *testing.T) {
	workerId := int64(1)
	datacenterId := int64(1)

	snowFlake, err := GetSnowFlak(workerId, datacenterId)
	if err != nil {
		t.Errorf("Error creating SnowFlake instance: %v", err)
	}

	// 测试生成ID的方法
	_, _ = snowFlake.NextId()
	// 测试最后获取的时间戳方法
	lastTimestamp := snowFlake.GetLastTimestamp()
	currentTimestamp := time.Now().UnixMilli()
	if lastTimestamp < 0 || lastTimestamp > currentTimestamp {
		t.Errorf("Invalid last timestamp: %d", lastTimestamp)
	}
}

func TestSnowFlake_NextId(t *testing.T) {
	workerId := int64(1)
	datacenterId := int64(1)

	snowFlake, err := GetSnowFlak(workerId, datacenterId)
	if err != nil {
		t.Errorf("Error creating SnowFlake instance: %v", err)
	}

	// 测试生成ID的方法
	id, err := snowFlake.NextId()
	if err != nil {
		t.Errorf("Error generating ID: %v", err)
	}

	// 检查ID是否符合预期的格式
	// 这里可以根据你的SnowFlake算法的具体实现来编写更多的检查条件
	// 这里只是简单地检查ID的长度是否符合预期
	expectedIdLength := 63
	if len(strconv.FormatInt(id, 2)) != expectedIdLength {
		t.Errorf("Expected ID %s length %d, got %d", strconv.FormatInt(id, 2), expectedIdLength, len(strconv.FormatInt(id, 2)))
	}
}
