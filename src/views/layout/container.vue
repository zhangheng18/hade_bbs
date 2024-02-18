<script setup>
import BBSHeader from '@/views/layout/header.vue'

</script>

<template>
  <el-container>
    <el-header>
      <BBSHeader></BBSHeader>
    </el-header>
    <el-main>
        <router-view></router-view>
    </el-main>
  </el-container>
</template>

<style scoped>

::-webkit-scrollbar {
  display: none;
}



:deep(.el-header) {
  background-color: #fff;
  text-align: center;
  line-height: 1.6;
  z-index: 1000;
  height: 52px;
}


:deep(.el-main) {
  background-color: #E9EEF3;
  height: calc(100vh - 52px);
}

body > :deep(.el-container) {
  margin-bottom: 40px;
}

</style>