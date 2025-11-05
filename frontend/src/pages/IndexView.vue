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
  terminalRef.value.terminal.error("fuck")
  return executeCommand(input)
}

onMounted(() => {
  userStore.getAndSetLoginUser()
})
</script>
