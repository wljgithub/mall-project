/**
 * 严肃声明：
 * 开源版本请务必保留此注释头信息，若删除我方将保留所有法律责任追究！
 * 本系统已申请软件著作权，受国家版权局知识产权以及国家计算机软件著作权保护！
 * 可正常分享和学习源码，不得用于违法犯罪活动，违者必究！
 * Copyright (c) 2020 陈尼克 all rights reserved.
 * 版权所有，侵权必究！
 */
import axios from 'axios'
import {Toast} from 'vant'
import router from '../router'

axios.defaults.baseURL = process.env.NODE_ENV == 'development' ? '//localhost:3000/api/v1' : '/api/v1'
// axios.defaults.baseURL = process.env.NODE_ENV == 'development' ? '//localhost:3000/api/v1' : '//47.99.134.126:28019/api/v1'
// axios.defaults.withCredentials = true
// axios.defaults.headers['X-Requested-With'] = 'XMLHttpRequest'
// axios.defaults.headers['Authorization'] = 'Bearer ' + localStorage.getItem('token') || ''
// axios.defaults.headers.post['Content-Type'] = 'application/json'

axios.interceptors.request.use(
  function(config) {
    // carried token if exist in local storage
    let token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = 'Bearer ' + token;
    }
    return config;
  },
  function(error) {
    return Promise.reject(error);
  }
);
axios.interceptors.response.use(res => {
  if (typeof res.data !== 'object') {
    Toast.fail('服务端异常！')
    return Promise.reject(res)
  }
  if (res.data.resultCode != 0) {
    if (res.data.message) Toast.fail(res.data.message)
    return Promise.reject(res.data)
  }

  return Promise.resolve(res.data)
}, error => {
  if (error.response.status === 401) {
    router.push({path: '/login'})
    return Promise.resolve(error.response.data)
  }
})


export default axios
