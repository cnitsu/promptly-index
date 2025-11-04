<script setup lang="ts">
import { ref, onMounted } from 'vue'

const userInput = ref<string>('')
const hiddenInput = ref<HTMLInputElement | null>(null)
const isBlinking = ref<boolean>(true)
const focusInput = () => {
  hiddenInput.value?.focus()
}
const onBlur = () => {
  setTimeout(() => focusInput(), 100);
}
onMounted(() => {
  focusInput()
  setInterval(() => {
    isBlinking.value = !isBlinking.value
  }, 500)
})
</script>

<template>
  <div class="terminal" @click="focusInput">
    <div class="input-line" id="istream">
      <span class="prompt">&gt;</span>
      <span class="text">{{ userInput }}</span>
      <span class="cursor" :class="{ blink: isBlinking }">_</span>
      <input ref="hiddenInput" v-model="userInput" class="hidden-input" @blur="onBlur">
    </div>
    <div class="output-line" id="ostream">
    </div>
  </div>
</template>

<style lang="css" scoped></style>
