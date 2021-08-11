<template>
  <div>
    <form>
      <h1>Welcome to <span>Sim</span>Vino</h1>
      <div>
        <label for="full_name">full name</label>
        <br />
        <input name="full_name" required type="text" placeholder="full name" />
      </div>
      <div>
        <label for="email">email</label>
        <br />
        <input name="email" required type="email" placeholder="email" />
      </div>
      <div>
        <label for="password">password</label>
        <br />
        <input
          required
          name="password"
          type="password"
          placeholder="password"
        />
      </div>
      <button type="submit">submit</button>
    </form>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
span {
  color: #0d6efd;
}
</style>