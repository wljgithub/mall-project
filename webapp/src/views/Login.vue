<!--
 * 严肃声明：
 * 开源版本请务必保留此注释头信息，若删除我方将保留所有法律责任追究！
 * 本系统已申请软件著作权，受国家版权局知识产权以及国家计算机软件著作权保护！
 * 可正常分享和学习源码，不得用于违法犯罪活动，违者必究！
 * Copyright (c) 2020 陈尼克 all rights reserved.
 * 版权所有，侵权必究！
 *
-->

<template>
  <div class="login">
    <s-header :back="'/home'" :name="type == 'login' ? '登录' : '注册'"></s-header>
    <img alt="" class="logo" src="//s.weituibao.com/1582958061265/mlogo.png">
    <div class="login-body login" v-if="type == 'login'">
      <van-form @submit="onSubmit">
        <van-field
          :rules="[{ required: true, message: '请填写用户名' }]"
          label="用户名"
          name="username"
          placeholder="用户名"
          v-model="username"
        />
        <van-field
          :rules="[{ required: true, message: '请填写密码' }]"
          label="密码"
          name="password"
          placeholder="密码"
          type="password"
          v-model="password"
        />
        <div class="verify">
          <Verify :fontSize="'16px'" :height="'40px'" :showButton="false" :type="2" :width="'100%'"
                  @error="error" @success="success" ref="loginVerifyRef"></Verify>
        </div>
        <div style="margin: 16px;">
          <div @click="toggle('register')" class="link-register">立即注册</div>
          <van-button block color="#1baeae" native-type="submit" round type="info">登录</van-button>
        </div>
      </van-form>
    </div>
    <div class="login-body register" v-else>
      <van-form @submit="onSubmit">
        <van-field
          :rules="[{ required: true, message: '请填写用户名' }]"
          label="用户名"
          name="username1"
          placeholder="用户名"
          v-model="username1"
        />
        <van-field
          :rules="[{ required: true, message: '请填写密码' }]"
          label="密码"
          name="password1"
          placeholder="密码"
          type="password"
          v-model="password1"
        />
        <div class="verify">
          <Verify :fontSize="'16px'" :height="'40px'" :showButton="false" :type="2" :width="'100%'"
                  @error="error" @success="success" ref="loginVerifyRef"></Verify>
        </div>
        <div style="margin: 16px;">
          <div @click="toggle('login')" class="link-login">已有登录账号</div>
          <van-button block color="#1baeae" native-type="submit" round type="info">注册</van-button>
        </div>
      </van-form>
    </div>
  </div>
</template>

<script>
  import sHeader from '@/components/SimpleHeader'
  import {login, register, getUserInfo} from '../service/user'
  import {setLocal, getLocal} from '@/common/js/utils'
  import {Toast} from 'vant'
  import Verify from 'vue2-verify'

  export default {
    data() {
      return {
        username: '',
        password: '',
        username1: '',
        password1: '',
        type: 'login',
        verify: false
      }
    },
    components: {
      sHeader,
      Verify
    },
    methods: {
      dealTriVer() {
        // 执行验证码的验证，通过 this.verify 知道验证码是否填写正确
        this.$refs.loginVerifyRef.$refs.instance.checkCode()
      },
      toggle(v) {
        this.verify = false
        this.type = v
      },
      async onSubmit(values) {
        // this.dealTriVer()
        // if (!this.verify) {
        //   Toast.fail('验证码未填或填写错误!')
        //   return
        // }
        if (this.type == 'login') {
          login({
            "loginName": values.username,
            "passwordMd5": this.$md5(values.password)
          }).then(
            res => {
              localStorage.setItem('token',res.data.Token)
              this.$router.push("/")
              // window.location.href = '/'
            }
          ).catch(err => {
            console.log(err)
          })
        } else {
          const {data} = await register({
            "loginName": values.username1,
            "password": this.$md5(values.password1)
          })
          Toast.success('注册成功')
          this.type = 'login'
        }
      },
      success(obj) {
        this.verify = true
        // 回调之后，刷新验证码
        obj.refresh()
      },
      error(obj) {
        this.verify = false
        // 回调之后，刷新验证码
        obj.refresh()
      }
    },
  }
</script>

<style lang="less">
  .login {
    .logo {
      width: 120px;
      height: 120px;
      display: block;
      margin: 80px auto 0px;
    }

    .login-body {
      padding: 0 20px;
    }

    .login {
      .link-register {
        font-size: 14px;
        margin-bottom: 20px;
        color: #1989fa;
        display: inline-block;
      }
    }

    .register {
      .link-login {
        font-size: 14px;
        margin-bottom: 20px;
        color: #1989fa;
        display: inline-block;
      }
    }

    .verify-bar-area {
      margin-top: 24px;

      .verify-left-bar {
        border-color: #1baeae;
      }

      .verify-move-block {
        background-color: #1baeae;
        color: #fff;
      }
    }

    .verify {
      > div {
        width: 100%;
      }

      display: flex;
      justify-content: center;

      .cerify-code-panel {
        margin-top: 16px;
      }

      .verify-code {
        width: 40% !important;
        float: left !important;
      }

      .verify-code-area {
        float: left !important;
        width: 54% !important;
        margin-left: 14px !important;

        .varify-input-code {
          width: 90px;
          height: 38px !important;
          border: 1px solid #e9e9e9;
          padding-left: 10px;
          font-size: 16px;
        }

        .verify-change-area {
          line-height: 44px;
        }
      }
    }
  }
</style>
