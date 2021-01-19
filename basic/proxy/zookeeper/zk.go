package zookeeper

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type ZkManager struct {
	hosts      []string
	conn       *zk.Conn
	pathPrefix string
}

func NewZkManager(hosts []string) *ZkManager {
	return &ZkManager{
		hosts:      hosts,
		pathPrefix: "/micro_gateway_",
	}
}

// 连接zk服务器
func (this *ZkManager) BuildConnect() error {
	conn, _, err := zk.Connect(this.hosts, time.Second*5)
	if err != nil {
		return err
	}
	this.conn = conn
	return nil
}

// 关闭服务器
func (this *ZkManager) Close() {
	if this.conn != nil {
		this.conn.Close()
	}
}

func (this *ZkManager) Get(path string) ([]byte, *zk.Stat, error) {
	return this.conn.Get(path)
}

func (this *ZkManager) Set(path string, config []byte, version int32) error {
	// 当前节点不存在，则创建
	hasPath, dStat, _ := this.conn.Exists(path)
	if !hasPath {
		_, err := this.conn.Create(path, config, 0, zk.WorldACL(zk.PermAll))
		return err
	}
	// 根据版本设置数据
	_, err := this.conn.Set(path, config, dStat.Version)
	if err != nil {
		return err
	}
	return nil
}

// 创建临时节点
// path是永久节点，path+/+host是临时节点
func (this *ZkManager) CreateTempNode(path, host string) error {
	// 先检查永久节点是否存在，不存在则先创建
	hasPath, _, err := this.conn.Exists(path)
	if err != nil {
		return err
	}
	if !hasPath {
		_, err = this.conn.Create(path, nil, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			return err
		}
	}

	// 创建临时节点
	tempPath := path + "/" + host
	hasPath, _, err = this.conn.Exists(tempPath)
	if err != nil {
		return err
	}
	if !hasPath {
		_, err = this.conn.Create(tempPath, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
		if err != nil {
			return err
		}
	}

	return nil
}

// 获取服务列表
func (this *ZkManager) GetServiceList(path string) ([]string, error) {
	list, _, err := this.conn.Children(path)
	return list, err
}

// watch机制，服务器断开或者重连
func (this *ZkManager) WatchServiceList(path string) (chan []string, chan error) {
	conn := this.conn
	snapshots := make(chan []string)
	errors := make(chan error)
	go func() {
		for {
			snapshot, _, eventchan, err := conn.ChildrenW(path)
			if err != nil {
				errors <- err
			}
			snapshots <- snapshot
			select {
			case evt := <-eventchan:
				if evt.Err != nil {
					errors <- evt.Err
				}
				fmt.Printf("ChildrenW Event Path:%v Type:%v\n", evt.Path, evt.Type)
			}
		}
	}()

	return snapshots, errors
}

// watch机制，监听节点值的变化
func (this *ZkManager) WatchPathData(path string) (chan []byte, chan error) {
	conn := this.conn
	snapshots := make(chan []byte)
	errors := make(chan error)

	go func() {
		for {
			data, _, eventchan, err := conn.GetW(path)
			if err != nil {
				errors <- err
				return
			}
			snapshots <- data

			select {
			case evt := <-eventchan:
				if evt.Err != nil {
					errors <- evt.Err
					return
				}
				fmt.Printf("GetW Event Path:%v Type:%v\n", evt.Path, evt.Type)
			}
		}
	}()

	return snapshots, errors
}
