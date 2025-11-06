<template>
  <div class="terminal-wrapper" :style="wrapperStyle" @click="handleWrapperClick">
    <div ref="terminalRef" class="terminal" :style="terminalStyle" @click="focusInput">
      <div v-if="!isLogin" class="output-area">
        <p>Promptly v0.1 (web tty)</p>
        <p>login: <span>{{ usernameTemp }}</span></p>
        <p>password: <span>{{ passwordDisplay }}</span></p>
        <!-- 文本输出 -->
        <template v-for="output in outputList" :key="index">
          <div class="terminal-row text-output" :class="output.status">
            <span class="output-text">{{ output.text }}</span>
          </div>
        </template>
      </div>
      <!-- 输出区域 -->
      <div v-else class="output-area">
        <p>Promptly v0.1 (web tty)</p>
        <p>welcome {{ user?.username }}</p>
        <p>type 'help' to get some sugguesion</p>
        <template v-for="(output, index) in outputList" :key="index">
          <!-- 命令行输出 -->
          <div v-if="output.type === 'command'" class="terminal-row command-line">
            <span class="prompt">{{ prompt }}</span>
            <span class="command-text">{{ output.text }}</span>
          </div>

          <!-- 文本输出 -->
          <div v-else-if="output.type === 'text'" class="terminal-row text-output" :class="output.status">
            <span class="output-text">{{ output.text }}</span>
          </div>

          <!-- 组件输出 -->
          <div v-else-if="output.type === 'component'" class="terminal-row component-output">
            <component :is="output.component" v-bind="output.props" />
          </div>

          <!-- 可折叠输出 -->
          <div v-else-if="output.collapsible" class="terminal-row collapsible-output">
            <div class="collapsible-header">
              <span class="prompt">{{ prompt }}</span>
              <span class="command-text">{{ output.text }}</span>
            </div>
            <div class="collapsible-content">
              <template v-if="output.result_list" v-for="res in output.result_list">
                <div class="terminal-row result-item" :class="res.status">
                  <span class="output-text">{{ res.text }}</span>
                </div>
              </template>
            </div>
          </div>
        </template>
      </div>

      <!-- 输入行 -->
      <div class="input-line">
        <span v-if="!isLogin" class="prompt">{{ "[login]" + loginStep + prompt + ": " }}</span>
        <span v-else class="prompt">{{ prompt }}</span>
        <span class="text" style="white-space: pre;">{{ userInput.slice(0, cursorPosition) }}</span>
        <span class="cursor" :class="{ blink: isBlinking }">_</span>
        <span class="text" style="white-space: pre;">{{ userInput.slice(cursorPosition) }}</span>
        <input ref="hiddenInput" v-model="userInput" class="hidden-input" @blur="onBlur" @keydown="handleKeyDown"
          @input="handleInput">
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, nextTick } from 'vue'
import { useUserStore } from '../composables/useUserStore';

interface Props {
  user?: promptly.UserType
  onLogin: (username: string, password: string) => Promise<any>
  fullScreen?: boolean
  onSubmit?: (input: string) => Promise<void>
}

const props = withDefaults(defineProps<Props>(), {
  fullScreen: false
})

const userInput = ref<string>('')
const hiddenInput = ref<HTMLInputElement | null>(null)
const terminalRef = ref<HTMLElement | null>(null)
const outputList = ref<promptly.OutputType[]>([])
const isBlinking = ref<boolean>(true)
const cursorPosition = ref<number>(0)

const usernameTemp = ref<string>('')
const passwordTemp = ref<string>('')
const loginStep = ref<'username' | 'password'>('username')

const isLogin = computed(() => {
  return props.user?.username ? true : false
})

const prompt = computed(() => {
  let promptFmt = props.user?.prompt
  let username = props.user?.username

  if (promptFmt && username) {
    return promptFmt.replace('{:username}', username)
  }
  return '>'
})

const wrapperStyle = computed(() => props.fullScreen ? {
  width: '100vw',
  height: '100vh',
  position: 'fixed' as const,
  top: 0,
  left: 0,
  zIndex: 9999
} : {
  width: 'inherit',
  height: 'inherit',
  position: 'fixed' as const
})

const terminalStyle = computed(() => ({
  width: '100%',
  height: '100%',
  backgroundColor: '#1a1a1a',
  color: '#ffffff',
  fontFamily: '"Consolas", "Monaco", "Courier New", monospace',
  fontSize: '14px',
  padding: '16px',
  overflow: 'auto',
  boxSizing: 'border-box' as const
}))

const terminal = {
  print: (output: Omit<promptly.OutputType, 'type'>) => {
    const newOutput: promptly.OutputType = {
      type: 'text',
      ...output
    }
    outputList.value.push(newOutput)
    nextTick(() => {
      scrollToBottom()
    })
  },

  printCommand: (command: string, result: promptly.OutputType[] = []) => {
    outputList.value.push({
      type: 'command',
      text: command,
      result_list: result,
      status: 'info'
    })
    nextTick(() => {
      scrollToBottom()
    })
  },

  clear: () => {
    outputList.value = []
  },

  error: (message: string) => {
    terminal.print({
      text: message,
      status: 'error'
    })
  },

  success: (message: string) => {
    terminal.print({
      text: message,
      status: 'success'
    })
  },

  warning: (message: string) => {
    terminal.print({
      text: message,
      status: 'warning'
    })
  },

  info: (message: string) => {
    terminal.print({
      text: message,
      status: 'info'
    })
  }
}

const scrollToBottom = () => {
  if (terminalRef.value) {
    terminalRef.value.scrollTop = terminalRef.value.scrollHeight
  }
}

const focusInput = () => {
  hiddenInput.value?.focus()
}

const onBlur = () => {
  setTimeout(() => focusInput(), 100)
}

const handleWrapperClick = () => {
  focusInput()
}

const handleSubmit = async () => {
  const command = userInput.value.trim()
  if (!command) return

  // 添加命令到输出
  terminal.printCommand(command)

  // 清空输入框
  userInput.value = ''

  // 执行命令
  if (props.onSubmit) {
    try {
      await props.onSubmit(command)
    } catch (error) {
      terminal.error(`Error executing command: ${error}`)
    }
  }
}

const handleKeyDown = async (event: KeyboardEvent) => {
  if (event.key === 'Enter') {
    event.preventDefault()
    if (!isLogin.value) {
      if (loginStep.value === 'username') {
        // 用户名输入完成，切换到密码输入
        usernameTemp.value = userInput.value
        userInput.value = ''
        cursorPosition.value = 0
        loginStep.value = 'password'
        const inputSpans = document.querySelectorAll<HTMLElement>('.text')
        for (const inputSpan of inputSpans) {
          inputSpan.style.visibility = 'hidden'
        }
      } else {
        // 密码输入完成
        passwordTemp.value = userInput.value
        userInput.value = ''
        cursorPosition.value = 0
        // 这里可以添加登录逻辑
        const res = await props.onLogin(usernameTemp.value, passwordTemp.value)
        if (res.code === 0) {
          const userStore = useUserStore()
          userStore.setLoginUser(res.data)
          terminal.clear()
        } else {
          terminal.error("Login incorrect")
          loginStep.value = 'username'
          usernameTemp.value = ''
          passwordTemp.value = ''
        }
        const inputSpans = document.querySelectorAll<HTMLElement>('.text')
        for (const inputSpan of inputSpans) {
          inputSpan.style.visibility = 'visible'
        }
      }
      return
    }
    handleSubmit()
    return
  }

  if (!isLogin.value && loginStep.value === 'password') {
    cursorPosition.value = 0
  }

  const target = event.target as HTMLInputElement
  let pos = target.selectionStart

  console.log(pos)
  if (pos == null) {
    cursorPosition.value = 0
    return
  }
  if (event.key === 'ArrowLeft') {
    pos -= 1
  }
  if (event.key === 'ArrowRight') {
    pos += 1
  }
  cursorPosition.value = Math.max(pos, 0)
}

const handleInput = (event: Event) => {
  // 实时更新登录显示
  if (!isLogin.value) {
    if (loginStep.value === 'username') {
      usernameTemp.value = userInput.value
    } else {
      passwordTemp.value = userInput.value
      cursorPosition.value = 0
      return
    }
  }
  const target = event.target as HTMLInputElement
  let pos = target.selectionStart
  cursorPosition.value = Math.max(pos ? pos : 0, 0)
}

const passwordDisplay = computed(() => {
  return Array(passwordTemp.value.length + 1).join('*')
})

onMounted(() => {
  focusInput()
  setInterval(() => {
    isBlinking.value = !isBlinking.value
  }, 500)
})

// 暴露 terminal 对象供父组件使用
defineExpose({
  terminal
})
</script>

<style lang="css" scoped>
.terminal-wrapper {
  display: flex;
  flex-direction: column;
}

.terminal {
  display: flex;
  flex-direction: column;
  border: 1px solid #333;
  border-radius: 4px;
}

.output-area {
  flex: 1;
  margin-bottom: 8px;
}

.terminal-row {
  margin-bottom: 4px;
  word-wrap: break-word;
  white-space: pre-wrap;
}

.command-line {
  color: #00ff00;
}

.prompt {
  color: #00ff00;
  margin-right: 8px;
  font-weight: bold;
}

.command-text {
  color: #ffffff;
}

.text-output {
  padding: 2px 0;
}

.output-text {
  line-height: 1.4;
}

.text-output.info {
  color: #ffffff;
}

.text-output.success {
  color: #00ff00;
}

.text-output.warning {
  color: #ffff00;
}

.text-output.error {
  color: #ff0000;
}

.text-output.system {
  color: #00ffff;
}

.component-output {
  padding: 4px 0;
}

.collapsible-output {
  margin-bottom: 8px;
}

.collapsible-header {
  cursor: pointer;
  user-select: none;
  padding: 4px 0;
}

.collapsible-header:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.collapsible-content {
  padding-left: 20px;
  border-left: 2px solid #333;
  margin-top: 4px;
}

.result-item {
  padding: 2px 0;
}

.input-line {
  display: flex;
  align-items: center;
  padding: 4px 0;
  border-top: 1px solid #333;
  margin-top: 8px;
}

.input-line .text {
  color: #ffffff;
}

.cursor {
  color: #ffffff;
  font-weight: bold;
  margin-left: 2px;
}

.cursor.blink {
  opacity: 0;
}

.hidden-input {
  position: absolute;
  left: -9999px;
  opacity: 0;
  pointer-events: none;
}

/* 滚动条样式 */
.terminal::-webkit-scrollbar {
  width: 8px;
}

.terminal::-webkit-scrollbar-track {
  background: #1a1a1a;
}

.terminal::-webkit-scrollbar-thumb {
  background: #555;
  border-radius: 4px;
}

.terminal::-webkit-scrollbar-thumb:hover {
  background: #777;
}
</style>
