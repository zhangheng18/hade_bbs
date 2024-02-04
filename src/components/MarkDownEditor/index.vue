<script setup lang="ts">
import {defineEmits, defineProps, onMounted, ref,watchEffect} from 'vue';
import Editor from '@toast-ui/editor';
import {EditorType,PreviewStyle as PreviewStyleType} from "@toast-ui/editor";

import '@toast-ui/editor/dist/toastui-editor.css';

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
    type: String as ()=>PreviewStyleType,
    required: false,
    default: 'tab',
  },

});
const emit = defineEmits(['update:modelValue']);
const editor  = ref()
onMounted(() => {
  const e = new Editor({
    el: editor.value,
    height: props.height,
    previewStyle:props.previewStyle,
    initialEditType: props.initialEditType,
    events: {
      change: () => emit('update:modelValue', e.getMarkdown()),

    },
    ...props.options,
  });
});

</script>

<template>
  <div>
    <div ref="editor"/>
  </div>
</template>

<style scoped>

</style>