<script setup lang="ts">

</script>

<template>
  <div class="terminal" @click="">
    <div class="output">
      <p>Promptly v0.1 (web tty)</p>
      <p>login: </p>
      <p>password: </p>
      <p>> welcome {username}</p>
      <p>> type 'help' to get some sugguesion</p>
    </div>
  </div>
</template>

<style lang="css" scoped>
.output p {
  flex: 1;
  overflow-y: auto;
  white-space: pre-wrap;
  width: 100%;
}
</style>
