## 简介
一个视频剪切接口的demo，登录使用jwt认证，剪切功能使用`ffmpeg`，使用[ffmpeg-go](https://github.com/u2takey/ffmpeg-go)操作，消息推送使用websocket

## 使用前的配置
#### 下载[ffmpeg](https://ffmpeg.org/download.html)
#### 根据环境设置环境变量(开发环境设置DEV为true)
```shell
export DEV=true
```
#### 执行`go mod tidy`安装依赖包

## 接口文档
#### 用户登录
```
ip:host/v1/cut/user/login?username=xxx&password=xxx
```
#### 用户注册
```
ip:host/v1/cut/user/login?username=xxx&password=xxx
```
### 剪切参数
```
ip:host/v1/cut/func/commit?start=xxx&duration=xxx
```
### 消息推送
```
ip:host/v1/cut/func/push
```
### 视频下载
```
ip:host/v1/cut/func/download
```

## 还能改进的点
#### 注册需要邮箱或者电话号码
> 使用阿里云短信服务验证，使用redis缓存验证码
#### 登录时验证密码
#### 用户密码加密存储，而不是存储明文密码
#### 考虑用户传入的参数是否有效
> 例如参数只能为正数， 如果前端有拖动进度条，验证start位置是否在end位置之前
#### 下载连接只能用一次(下载链接存入记录中)
#### 使用oss对象存储
> 用户从前端上传视频后，存储在oss，后台完成剪切任务或者用户需要下载时再剪切，剪切完成的视频存到oss，并生成下载链接。减小服务器的压力
#### 

## 未全部完成的点
#### 单元测试
#### 用户上传视频
#### 任务进度
