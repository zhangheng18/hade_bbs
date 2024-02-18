<script setup>
import {useUserStore} from "@/stores/user.js";
import router from "@/router/index.js";
import {ElMessage} from "element-plus";

const store = useUserStore()

function handleLogout() {
  store.logout().then(() => {
        router.push({path: '/'})
        ElMessage.success('登出成功')
      }
  ).catch((e) => {
    console.log('登出失败,', e)
    ElMessage.error('登出失败', e)
  })
}

</script>

<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="12">
      <el-row>
        <el-col :span="8">
          <div class="hade_title"><a href="/" id="home-link" title="hadeBBS">hadeBBS</a></div>
        </el-col>
        <el-col class="" :span="10" :offset="6">

          <div class="header_name">
            <router-link class="to-link" :to="{path: '/create'}"><a title="question">我要提问</a></router-link>
            | {{ store.name }} | <a href="#" @click="handleLogout" title="logout">登出</a>
          </div>
        </el-col>
      </el-row>
    </el-col>
  </el-row>
</template>

<style scoped>
.hade_title {
  float: left;
  vertical-align: top;
  font-size: 24px;
  font-weight: 400;
  margin: 0 15px 0 0;
  line-height: 34px;
  padding-top: 10px;
}

.header_name {
  float: right;
  font-size: 13px;
  font-weight: 400;
  margin: 0 15px 0 0;
  line-height: 34px;
  padding-top: 10px;
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
}

.header_name > a {
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
}

.hade_title > a {
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
}

:deep(.el-header) {
  background-color: #409EFF;
  text-align: center;
  line-height: 80px;
}

</style>