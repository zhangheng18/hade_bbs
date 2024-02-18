<script setup lang="ts">
import {onMounted, ref} from 'vue';
import Editor, {EditorType, PreviewStyle as PreviewStyleType} from '@toast-ui/editor';

import '@toast-ui/editor/dist/toastui-editor.css';
import '@toast-ui/editor/dist/toastui-editor-viewer.css';

import defaultOptions from './default-options.js';

const props = defineProps({
  modelValue: {
    type: String,
    required: false,
    default: '',
  },
  options: {
    type: Object,
    required: false,
    default: () => defaultOptions,
  },
  initialEditType: {
    type: String as () => EditorType,
    required: false,
    default: 'markdown',
  },
  height: {
    type: String,
    required: false,
    default: '300px'
  },
  language: {
    type: String,
    required: false,
    default: 'zh-CN' // https://github.com/nhn/tui.editor/blob/master/docs/en/i18n.md
  },
  previewStyle: {
    type: String as () => PreviewStyleType,
    required: false,
    default: 'tab',
  },
  viewer: {
    type: Boolean,
    required: false,
    default: false,
  }
});
const emit = defineEmits(['update:modelValue']);
const editor = ref()

onMounted(() => {
  let e = null
  if (props.viewer){
    e = new Editor.factory({
      el: editor.value,
      viewer: props.viewer,
      initialValue: props.modelValue,
    });
  }else{
     e = new Editor({
      el: editor.value,
      height: props.height,
      previewStyle: props.previewStyle,
      initialEditType: props.initialEditType,
      initialValue: props.modelValue,
      events: {
        change: () => emit('update:modelValue', e.getHTML()),
      },
      ...props.options,
    });
  }
  console.log('ver',props.viewer,e, props.modelValue)

});

</script>

<template>
  <div>
    <div ref="editor" style="width: 100%"/>
  </div>
</template>

<style scoped>


:deep(.toastui-editor-contents) {
  font-family: -apple-system,BlinkMacSystemFont,Helvetica Neue,PingFang SC,Microsoft YaHei,Source Han Sans SC,Noto Sans CJK SC,WenQuanYi Micro Hei,sans-serif;
}
</style>