// UserInfo
package main

import (
	"net"
	"sync"
)

/*连接与id的映射*/
type userMap struct {
	sync.Mutex
	M map[net.Conn]int
}

/*ID与玩家属性对象映射*/
type userInfoMap struct {
	sync.Mutex
	M map[int]userInfo
}

type userInfo struct {
	Id   int
	Name string
	conn net.Conn
}

var UserMap = &userMap{M: make(map[net.Conn]int)}

var UserInfoMap = userInfoMap{M: make(map[int]userInfo)}

func (user *userMap) Get(key net.Conn) int {
	user.Lock()
	result := user.M[key]
	user.Unlock()
	return result
}

func (user *userMap) Delete(key net.Conn) {
	user.Lock()
	delete(user.M, key)
	user.Unlock()
}

func (user *userMap) Put(key net.Conn, value int) {
	user.M[key] = value
}

func (user *userMap) GetMap() map[net.Conn]int {
	return user.M
}

func (info *userInfoMap) Get(key int) userInfo {
	info.Lock()
	result := info.M[key]
	info.Unlock()
	return result
}

func (info *userInfoMap) Delete(key int) {
	info.Lock()
	delete(info.M, key)
	info.Unlock()
}

func (info *userInfoMap) Put(key int, value userInfo) {
	info.Lock()
	info.M[key] = value
	info.Unlock()
}

func (info *userInfoMap) GetMap() map[int]userInfo {
	return info.M
}
