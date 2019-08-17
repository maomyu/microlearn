/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-17 10:22:12
 * @LastEditTime: 2019-08-17 10:24:31
 * @LastEditors: Please set LastEditors
 */
package config

type Profiles interface {
	GetInclude() string
}
type defaultProfiles struct {
	Include string `json:"include"`
}

// 实现接口方法
func (p defaultProfiles) GetInclude() string {
	return p.Include
}
