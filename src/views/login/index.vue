<script lang="ts" setup>
import type {FormInstance} from 'element-plus'
import {reactive, ref} from "vue";

const loginFormRef = ref<FormInstance>()

const loginFormData = reactive({
  username: '',
  password: ''
})

function handleLogin() {
  loginFormRef.value?.validate((valid) => {
    if (valid) {
      console.log('loginFormData', loginFormData)
    }
  })
}
</script>

<template>
  <div class="login">
    <el-card>
      <h2>登录</h2>
      <el-divider></el-divider>
      <el-form class="login-form"
               ref="loginFormRef"
               :model="loginFormData"
               status-icon
      >
        <el-form-item label="" prop="username" required>
          <el-input v-model="loginFormData.username" placeholder="用户名" prefix-icon="User"></el-input>
        </el-form-item>

        <el-form-item label="" prop="password" required>
          <el-input type="password" v-model="loginFormData.password" placeholder="密码" prefix-icon="Lock"></el-input>
        </el-form-item>

        <el-form-item>

          <el-row>
            <el-col class="register">
              还没有账号？请点击
              <el-link type="primary">注册</el-link>
            </el-col>
          </el-row>
        </el-form-item>

        <el-form-item>
          <el-button
              class="login-button"
              type="primary"
              native-type="submit"
              @click.native.prevent="handleLogin"
              style="width: 100%"
          >登录
          </el-button>
        </el-form-item>
      </el-form>

    </el-card>

  </div>
</template>

<style scoped>
.login {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

.login-form {
  width: 400px;
}
</style>