<script setup>
import MkEditor from '@/components/MarkDownEditor/index.vue'
import {reactive} from "vue";

const question = reactive({
  title: '',
  content: ''
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
    ['scrollSync'],
  ]
}

function onSubmit() {
  console.log(question)
}

</script>

<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="12">
      <el-card class="box-card" shadow="never">
        <el-form ref="form" :model="question">
          <el-form-item >
            <el-input v-model="question.title"></el-input>
          </el-form-item>

          <MkEditor :options="editorOptions"
                  height="600px"
                  initialEditType="wysiwyg"
                  previewStyle="tab" />
          <br />
          <el-form-item>
            <el-button type="primary" @click="onSubmit">立即提问</el-button>
            <el-button>取消</el-button>
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