/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-17 10:21:49
 * @LastEditTime: 2019-08-17 10:34:49
 * @LastEditors: Please set LastEditors
 */
package config

type ConsulConfig interface {
	GetEnabled() bool
	GetHost() string
	GetPort() int
}
type defaultConsulConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

func (c defaultConsulConfig) GetEnabled() bool {
	return c.Enabled
}

func (c defaultConsulConfig) GetHost() string {
	return c.Host
}

func (c defaultConsulConfig) GetPort() int {
	return c.Port
}
