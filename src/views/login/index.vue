<script lang="ts" setup>
import {ElMessage, FormInstance} from 'element-plus'
import {useUserStore} from "@/stores/user";
import {reactive, ref} from "vue";
import {useRouter} from 'vue-router'

const loginFormRef = ref<FormInstance>()

const formData = reactive({
  username: 'zhang',
  password: '123456'
})


const store = useUserStore()
const router = useRouter()

function handleLogin() {
  loginFormRef.value?.validate((valid) => {
    if (valid) {
      store.login(formData).then(() => {
        router.push({path: '/'})
        ElMessage({
          message: "登录成功",
          type: 'success',
          duration: 3 * 1000
        })
      }).catch((e) => {
        console.log('登录失败,', e)
      })
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
               :model="formData"
               status-icon
      >
        <el-form-item label="" prop="username" required>
          <el-input v-model="formData.username" placeholder="用户名" prefix-icon="User"></el-input>
        </el-form-item>

        <el-form-item label="" prop="password" required>
          <el-input type="password" v-model="formData.password" placeholder="密码" prefix-icon="Lock"></el-input>
        </el-form-item>

        <el-form-item>

          <el-row>
            <el-col class="register">
              还没有账号？请点击 <router-link class="to-link" :to="{path: '/register'}"><el-link type="primary" >注册</el-link></router-link>
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