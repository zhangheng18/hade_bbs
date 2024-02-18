<script lang="ts" setup>
import {reactive, ref} from "vue";
import {useRouter} from 'vue-router'
import {ElMessage, FormInstance} from 'element-plus'
import {useUserStore} from "@/stores/user";
import request from "@/utils/request.js";


const FormRef = ref<FormInstance>()
const loading = ref(false)

const formData = reactive({
  username: 'zhang',
  email: 'zhang@test.com',
  password: '123456',
  repassword: '123456'
})


const store = useUserStore()
const router = useRouter()

function handleRegister() {
  FormRef.value?.validate(async (valid) => {
    if (valid) {
      if (formData.repassword !== formData.password) {
        ElMessage.error("两次输入密码不一致");
        return;
      }
      try {
        loading.value = true
        const resp = await request.post('/api/user/register', formData)
        const msg = resp.data
        loading.value = false
        ElMessage.success({
          message: msg,
          duration: 5 * 1000
        })
      } catch (error) {
        console.log(error)
        loading.value = false
        ElMessage.error("注册失败", error)
      }
    }
  })
}
</script>

<template>
  <div class="register">
    <el-card>
      <h2>注册</h2>
      <el-divider></el-divider>
      <el-form class="register-form"
               ref="FormRef"
               :model="formData"
               status-icon
      >
        <el-form-item label="" prop="username" required>
          <el-input v-model="formData.username" placeholder="用户名" prefix-icon="User"></el-input>
        </el-form-item>
        <el-form-item label="" prop="email" required>
          <el-input v-model="formData.email" placeholder="邮箱" prefix-icon="Message"></el-input>
        </el-form-item>

        <el-form-item label="" prop="password" required>
          <el-input type="password" v-model="formData.password" placeholder="密码" prefix-icon="Lock"></el-input>
        </el-form-item>

        <el-form-item label="" prop="password" required>
          <el-input type="password" v-model="formData.repassword" placeholder="密码" prefix-icon="Lock"></el-input>
        </el-form-item>


        <el-form-item>
          <el-button
              :loading="loading"
              class="login-button"
              type="primary"
              native-type="submit"
              @click.native.prevent="handleRegister"
              style="width: 100%"
              block
          >注册
          </el-button>
        </el-form-item>
      </el-form>

    </el-card>

  </div>
</template>

<style scoped>
.register {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

.register-form {
  width: 400px;
}
</style>