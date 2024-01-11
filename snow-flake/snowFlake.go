package snow_flake

import (
	"sync"
	"time"
)

// SnowFlake
// @Description: 全局唯一雪花算法
// 1位无用位； 41位时间戳; 5位机器ID; 5位机房ID; 12位序列
type SnowFlake struct {
	// 数据中心(机房) id
	datacenterId int64
	// 机器ID
	workerId int64
	// 同一时间的序列
	sequence int64

	// 开始时间戳（2020-5-22 17:30:00）(UTC)
	twepoch int64

	// 机房号，的ID所占的位数 5个bit 最大:11111(2进制)--> 31(10进制)
	datacenterIdBits int64

	// 机器ID所占的位数 5个bit 最大:11111(2进制)--> 31(10进制)
	workerIdBits int64

	// 5 bit最多只能有31个数字，就是说机器id最多只能是32以内
	maxWorkerId int64

	// 5 bit最多只能有31个数字，机房id最多只能是32以内
	maxDatacenterId int64

	// 同一时间的序列所占的位数 12个bit 111111111111 = 4095  最多就是同一毫秒生成4096个
	sequenceBits int64

	// workerId的偏移量
	workerIdShift int64

	// datacenterId的偏移量
	datacenterIdShift int64

	// timestampLeft的偏移量
	timestampLeftShift int64

	// 序列号掩码 4095 (0b111111111111=0xfff=4095)
	// 用于序号的与运算，保证序号最大值在0-4095之间
	sequenceMask int64

	// 最近一次时间戳
	lastTimestamp int64
}

// GetSnowFlak
//
//	@Description: 			创建类
//	@receiver s
//	@param workerId			机器ID
//	@param datacenterId		机房ID
//	@param sequence			同一时间的序列
//	@return *SnowFlak
//	@return error
func GetSnowFlak(workerId int64, datacenterId int64) (*SnowFlake, error) {
	snowFlake := &SnowFlake{
		datacenterId:     datacenterId,
		workerId:         workerId,
		sequence:         0,
		twepoch:          1590160200,
		datacenterIdBits: 5,
		workerIdBits:     5,
		//maxWorkerId:        0,
		//maxDatacenterId:    0,
		sequenceBits: 12,
		//workerIdShift:      0,
		//datacenterIdShift:  0,
		//timestampLeftShift: 0,
		//sequenceMask:       0,
		lastTimestamp: -1,
	}
	snowFlake.maxWorkerId = -1 ^ (-1 << snowFlake.workerIdBits)
	snowFlake.maxDatacenterId = -1 ^ (-1 << snowFlake.datacenterIdBits)
	snowFlake.workerIdShift = snowFlake.sequenceBits
	snowFlake.datacenterIdShift = snowFlake.sequenceBits + snowFlake.workerIdBits
	snowFlake.timestampLeftShift = snowFlake.sequenceBits + snowFlake.workerIdBits + snowFlake.datacenterIdBits
	snowFlake.sequenceMask = -1 ^ (-1 << snowFlake.sequenceBits)

	//if workerId > snowFlake.maxWorkerId || workerId < 0 {
	//	return nil, status.Error(codes.FailedPrecondition, "worker Id can't be greater than %d or less than 0")
	//}
	//if datacenterId > snowFlake.maxDatacenterId || datacenterId < 0 {
	//	return nil, status.Error(codes.FailedPrecondition, "datacenter Id can't be greater than %d or less than 0")
	//}

	return snowFlake, nil
}

// GetWorkerId 获取机器ID
func (s *SnowFlake) GetWorkerId() int64 {
	return s.workerId
}

// GetDatacenterId 获取机房ID
func (s *SnowFlake) GetDatacenterId() int64 {
	return s.datacenterId
}

// GetLastTimestamp 获取最新一次获取的时间戳
func (s *SnowFlake) GetLastTimestamp() int64 {
	return s.lastTimestamp
}

// NextId 获取下一个随机的ID
func (s *SnowFlake) NextId() (int64, error) {
	// 加锁互斥
	lock := &sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()

	// 获取当前时间戳，单位毫秒
	timestamp := time.Now().Unix()

	//if timestamp < s.lastTimestamp {
	//	return 0, status.Errorf(codes.Unknown, "Clock moved backwards.  Refusing to generate id for "+strconv.FormatInt(s.lastTimestamp-timestamp, 10)+" milliseconds")
	//}

	// 去重
	if s.lastTimestamp == timestamp {
		s.sequence = (s.sequence + 1) & s.sequenceMask
		// sequence序列大于4095
		if s.sequence == 0 {
			// 调用到下一个时间戳的方法
			timestamp = s.tilNextMillis(s.lastTimestamp)
		}
	} else {
		// 如果是当前时间的第一次获取，那么就置为0
		s.sequence = 0
	}

	// 记录上一次的时间戳
	s.lastTimestamp = timestamp

	// 偏移计算
	return ((timestamp - s.twepoch) << s.timestampLeftShift) |
		(s.datacenterId << s.datacenterIdShift) |
		(s.workerId << s.workerIdShift) |
		s.sequence, nil
}

func (s *SnowFlake) tilNextMillis(lastTimestamp int64) int64 {
	// 获取最新时间戳
	timestamp := s.timeGen()
	// 如果发现最新的时间戳小于或者等于序列号已经超4095的那个时间戳
	for timestamp <= lastTimestamp {
		// 不符合则继续
		timestamp = s.timeGen()
	}
	return timestamp
}

func (s *SnowFlake) timeGen() int64 {
	return time.Now().Unix()
}
