package utils

import (
	"errors"
	"sync"
	"time"
)

// 因为snowFlake目的是解决分布式下生成唯一id 所以ID中是包含集群和节点编号在内的

const (
	// 每台机器(节点)的ID位数 10位最大可以有2^10=1024个节点
	workerBits uint8 = 10

	// 表示每个集群下的每个节点, 1毫秒内可生成的id序号的二进制位数 即每毫秒可生成 2^12-1=4096个唯一ID
	numberBits uint8 = 12

	// 这里求最大值使用了位运算, -1 的二进制表示为 1 的补码
	// 节点ID的最大值, 用于防止溢出 (1023)
	workerMax int64 = -1 ^ (-1 << workerBits)

	// 用来表示生成id序号的最大值
	numberMax int64 = -1 ^ (-1 << numberBits)

	// 时间戳向左的偏移量
	timeShift = workerBits + numberBits

	// 节点ID向左的偏移量
	workerShift = numberBits

	// 41位字节作为时间戳数值的话 大约68年就会用完
	// 假如你2010年1月1日开始开发系统 如果不减去2010年1月1日的时间戳 那么白白浪费40年的时间戳啊！
	// 这个一旦定义且开始生成ID后千万不要改了 不然可能会生成相同的ID
	epoch int64 = 1569230536000
)

type Worker struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	timestamp int64      // 记录时间戳
	workerId  int64      // 该节点的ID
	times     int64      // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("invalid worker Id, must between 0 to 1024")
	}

	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		times:     0,
	}, nil
}

func (w *Worker) NextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // 纳秒转毫秒
	if w.timestamp == now {
		w.times++

		if w.times > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.times = 0
		w.timestamp = now // 将机器上一次生成ID的时间更新为当前时间
	}

	// 第一段 now - epoch 为该算法目前已经奔跑了xxx毫秒
	// 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
	ID := (now-epoch)<<timeShift | (w.workerId << workerShift) | (w.times)
	return ID
}
