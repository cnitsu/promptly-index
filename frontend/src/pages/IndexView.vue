<template>
  <Terminal ref="terminalRef" :user="loginUser" full-screen :on-submit="submitCommand" />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import Terminal from '../components/Terminal.vue';
import { executeCommand } from '../core/executor';
import { useUserStore } from '../composables/useUserStore';
import { storeToRefs } from 'pinia';

const terminalRef = ref()
const userStore = useUserStore()

const { loginUser } = storeToRefs(userStore)

const submitCommand = async (input: string) => {
  const terminal = terminalRef.value.terminal
  executeCommand(input, terminal)
}

onMounted(() => {
  userStore.getAndSetLoginUser()
})
</script>
