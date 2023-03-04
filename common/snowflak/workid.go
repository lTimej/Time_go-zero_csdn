package snowflak

import (
	"errors"
	"time"
)

var (
	// 64位ID的划分
	WORKER_ID_BITS     int64 = 5
	DATACENTER_ID_BITS int64 = 5
	SEQUENCE_BITS      int64 = 12
	// 最大取值计算
	MAX_WORKER_ID     int64 = -1 ^ (-1 << WORKER_ID_BITS)
	MAX_DATACENTER_ID int64 = -1 ^ (-1 << DATACENTER_ID_BITS)
	// 移位偏移计算
	WOKER_ID_SHIFT       int64 = SEQUENCE_BITS
	DATACENTER_ID_SHIFT  int64 = SEQUENCE_BITS + WORKER_ID_BITS
	TIMESTAMP_LEFT_SHIFT int64 = SEQUENCE_BITS + WORKER_ID_BITS + DATACENTER_ID_BITS
	// 序号循环掩码
	SEQUENCE_MASK int64 = -1 ^ (-1 << SEQUENCE_BITS)
	// Twitter元年时间戳
	TWEPOCH int64 = 1288834974657
)

type WorkID struct {
	worker_id      int64
	datacenter_id  int64
	sequence       int64
	last_timestamp int64
}

func NewSnowFlak(worker_id, datacenter_id, sequence, last_timestamp int) (*WorkID, error) {
	if int64(worker_id) > MAX_WORKER_ID || worker_id < 0 {
		return nil, errors.New("worker_id值越界")
	}
	if int64(datacenter_id) > MAX_DATACENTER_ID || datacenter_id < 0 {
		return nil, errors.New("datacenter_id值越界")
	}
	return &WorkID{
		worker_id:      int64(worker_id),
		datacenter_id:  int64(datacenter_id),
		sequence:       int64(sequence),
		last_timestamp: int64(last_timestamp),
	}, nil
}
func (wi *WorkID) getTimestamp() int64 {
	return time.Now().UnixMilli()
}
func (wi *WorkID) tilNextMillis(last_timestamp int64) int64 {
	timestamp := wi.getTimestamp()
	for timestamp <= last_timestamp {
		timestamp = wi.getTimestamp()
	}
	return timestamp
}
func (wi *WorkID) GetId() int64 {
	timestamp := wi.getTimestamp()

	if timestamp < wi.last_timestamp {
		return 0
	}
	if timestamp == wi.last_timestamp {
		wi.sequence = (wi.sequence + 1) & SEQUENCE_MASK
	}
	if wi.sequence == 0 {
		timestamp = wi.tilNextMillis(wi.last_timestamp)
	} else {
		wi.sequence = 0
	}
	wi.last_timestamp = timestamp
	new_id := ((timestamp - TWEPOCH) << TIMESTAMP_LEFT_SHIFT) | (wi.datacenter_id << DATACENTER_ID_SHIFT) | (wi.worker_id << WOKER_ID_SHIFT) | wi.sequence
	return new_id
}
