<script setup>
import {onMounted, reactive, ref} from "vue";
import request from "@/utils/request.js";
import router from "@/router/index.js";
import {formatDate} from "@/utils/format.js";
import MkEditor from "@/components/MarkDownEditor/index.vue";
import {ElMessage} from "element-plus";


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

const question = reactive({title: '', answers: [], context: ''})
const answerContext = ref()


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

function gotoDeleteQuestion() {
  // 删除问题
  return request.post("/api/question/delete", null,{
    params: {
      "id": question.id
    }
  }).then(function (resp) {
    ElMessage.success(resp.data)
    router.push({path: '/'})
  })
}

function gotoQuestionEdit() {
  // 编辑问题
  console.log('gotoQuestionEdit')
  router.push({path: '/edit', query: {'id': question.id}})
}

function postAnswer() {
  // 回答问题
  return request.post("/api/answer/create", {
    question_id: question.id,
    context: answerContext.value
  }).then(function (resp) {
    ElMessage.success(resp.data)
    router.go(0)
  })
}

function gotoDeleteAnswer(id) {
  // 删除回答
  return request.post("/api/answer/delete", null,{
    params:{"id": id}}
  ).then(function (resp) {
    ElMessage.success(resp.data)
    router.go(0)
  })
}


</script>

<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="12">
      <el-divider content-position="left">问题</el-divider>
      <el-card v-if="question.title" class="box-card" shadow="never">
        <div slot="header" class="clearfix">
          <span>{{ question.title }} </span>
          <span class="header_name" style="margin-right: 5px; float: right;">
              <span @click="gotoQuestionEdit">修改</span> ｜ <span  @click="gotoDeleteQuestion">删除</span>
          </span>
        </div>
        <el-divider/>
        <MkEditor :viewer="true" v-model="question.context"/>
      </el-card>

      <el-divider content-position="left">所有回答</el-divider>
      <el-card v-for="answer in question.answers" style="margin-top: 5px; " class="box-card"
               shadow="hover">
        <div slot="header" class="clearfix">
          <span>{{ answer.author.user_name }}| {{ formatDate(answer.created_at) }}</span>
          <span class="header_name" style="margin-right: 5px; float: right;"
             @click="gotoDeleteAnswer(answer.id)">删除</span>

          <el-divider/>

          <MkEditor :viewer="true" v-model="answer.context"/>
        </div>
      </el-card>

      <el-divider content-position="left">我来回答</el-divider>
      <el-card class="box-card" shadow="never">
        <el-form ref="form" :model="question">
          <MkEditor :options="editorOptions"
                    height="300px"
                    initialEditType="markdown"
                    previewStyle="vertical" v-model="answerContext"/>
          <br/>
          <el-form-item>
            <el-button type="primary" @click="postAnswer">提交</el-button>
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

.header_name {
  float: right;
  font-size: 13px;
  font-weight: 400;
  margin: 0 15px 0 0;
  line-height: 34px;
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
  cursor: pointer;
}

.header_name > a {
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
}

</style>