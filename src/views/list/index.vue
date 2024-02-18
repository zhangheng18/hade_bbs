<script setup>
import {computed, ref} from 'vue'
import MkEditor from "@/components/MarkDownEditor/index.vue";
import request from "@/utils/request.js";
import router from "@/router/index.js";
import {formatDate} from "../../utils/format.js";

const loading = ref(false)
const noMore = ref(false)

const disabled = computed(() => {
  console.log("disabled", loading.value || noMore.value)
  return loading.value || noMore.value
})

const questions = ref([])
const page = ref(1)
const size = ref(6)



function getQuestions(){
  if (noMore.value === true) {
    return
  }
  request.get('/api/question/list', {
    params: {
      page: page.value,
      size: size.value,
    }
  }).then(function (response) {
    const qa = response.data
    if (qa === null || qa.length === 0 || qa.length < size.value) {
      noMore.value = true
    }
    questions.value = questions.value.concat(qa)
    page.value = page.value + 1
    console.log("qa",qa,"page:",page.value,"noMore:", noMore.value, "questions:", questions.value.length)
    loading.value = false
  }).catch(function (error) {
    console.log(error)
  })

}
function load() {
  if (noMore.value === true) {
    return
  }
  loading.value = true
  setTimeout(() => {
    getQuestions()
  }, 600)

}

function gotoDetail(id) {
  router.push({path: '/detail', query:{'id': id}})
}

</script>

<template>
  <el-row type="flex" justify="center" align="middle">
    <el-col :span="12" >
      <div class="infinite-list-wrapper" style="overflow-y:scroll;height: calc( 100vh - 72px);
">
        <ul
            class="list"
            v-infinite-scroll="load"
            :infinite-scroll-disabled="disabled"
            :infinite-scroll-immediate="true"
        >
          <el-card v-for="question in questions" class="box-card" shadow="hover">
            <div slot="header" class="clearfix">
              <span>{{ question.title }}</span>
            </div>
            <div class="text item">
              <MkEditor :viewer="true" v-model="question.context"/>
            </div>
            <div class="bottom clearfix">
              <time class="time">{{formatDate(question.created_at)}} ｜ {{question.author.user_name}}  | {{question.answer_num}} 回答</time>
              <el-button link:="true" class="button" @click="gotoDetail(question.id)">去看看</el-button>
            </div>
          </el-card>
        </ul>
        <p v-if="loading" class="loading_tips">加载中...</p>
        <p v-if="disabled && !loading " class="loading_tips">没有更多了</p>
      </div>
    </el-col>
  </el-row>
</template>

<style scoped>


::-webkit-scrollbar {
  display: none;
}

.list {
  padding: 0;
}

.loading_tips {
  text-align: center;
  font-size: 13px;
  color: #999;
}

.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

:deep(.carousel) {
  text-align: center;
}

.box-card {
  margin-top: 10px;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}


.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.button {
  padding: 1px 5px;
  float: right;
}


.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both
}

</style>