<template>
  <div id="app">
    <header>
      <h1>🧪 TestHub</h1>
      <nav>
        <router-link to="/">Home</router-link> |
        <router-link to="/testItemEdit">試験項目を登録する</router-link>
        <router-link to="/testItemList">試験項目一覧</router-link>
      </nav>
    </header>

    <main>
      <router-view />
    </main>
  </div>
</template>

<script setup>
</script>

<style scoped>
#app {
  font-family: 'Segoe UI', sans-serif;
  padding: 1rem;
  max-width: 800px;
  margin: 0 auto;
}

header {
  background-color: #3eaf7c;
  color: white;
  padding: 1rem;
  text-align: center;
}

footer {
  margin-top: 2rem;
  text-align: center;
  color: #888;
  font-size: 0.9rem;
}
</style>
