<script setup>

import MkEditor from "@/components/MarkDownEditor/index.vue";
import {onMounted, reactive} from "vue";
import router from "@/router/index.js";
import request from "@/utils/request.js";
import {ElMessage} from "element-plus";

const editorOptions = {
  minHeight: '600px',
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
const question = reactive({title: '', context: ''})


onMounted(async () => {
  let id = parseInt(router.currentRoute.value.query.id)
  if (id) {
    await getDetail(id)
  }
});


async function getDetail(id) {
  // 获取详情
  try {
    const resp = await request.get("/api/question/detail",
        {
          params: {
            "id": id
          }
        })
    Object.assign(question, resp.data)
  } catch (e) {
    console.log(e)
  }
}

function onSubmit() {
  // 更新问题
  return request.post("/api/question/edit", {
    id: question.id,
    title: question.title,
    context: question.context
  }).then(function (resp) {
    ElMessage.success(resp.data)
    router.push({path: '/detail', query: {'id': question.id}})
  })
}

</script>

<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="12">
      <el-card class="box-card" shadow="never">
        <el-form ref="form" :model="question" label-width="70px">
          <el-form-item label="问题标题">
            <el-input v-model="question.title">{{ question.title }}</el-input>
          </el-form-item>
          <el-form-item label="问题描述" v-if="question.context">
            <MkEditor style="width: 100%" :options="editorOptions"
                      height="500px"
                      initialEditType="wysiwyg"
                      previewStyle="vertical" v-model="question.context"/>

          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">更新</el-button>
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