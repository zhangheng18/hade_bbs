<script setup>
import MkEditor from '@/components/MarkDownEditor/index.vue'
import {reactive} from "vue";
import {ElMessage} from "element-plus";
import request from "@/utils/request.js";
import router from "@/router/index.js";

const question = reactive({
  title: 'test',
  context: 'context'
})


const editorOptions = {
  minHeight: '300px',
  language: 'zh-CN',
  useCommandShortcut: true,
  usageStatistics: false,
  hideModeSwitch: true,
  toolbarItems: [
    ['heading', 'bold', 'italic', 'strike'],
    ['hr', 'quote'],
    ['ul', 'ol', 'task',],
    ['table', 'image', 'link'],
    ['code', 'codeblock'],
  ]
}

async function onSubmit() {
  if (question.title === '' || question.context === '') {
    ElMessage.error('问题标题和内容不能为空')
    return
  }

  const resp = await request({
    method: 'POST',
    url: "/api/question/create",
    data: question,
  })
  router.push({path: '/'})
  ElMessage.success(resp.data)

}

</script>

<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="12">
      <el-card class="box-card" shadow="never">
        <el-form ref="form" :model="question">
          <el-form-item>
            <el-input v-model="question.title" placeholder="问题"></el-input>
          </el-form-item>

          <MkEditor :options="editorOptions"
                    height="600px"
                    initialEditType="markdown"
                    previewStyle="vertical" v-model="question.context"/>
          <br/>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">立即提问</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<style scoped>
.text {
  font-size: 14px;
}

.item {
  padding: 18px 0;
}

</style>