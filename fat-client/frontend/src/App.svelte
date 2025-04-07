<script lang="ts">
  import './app.css';
  import Sidebar from "./Sidebar.svelte";
  import Greet from "./Greet.svelte";
</script>

<main class="h-full">
  <div class="flex">
    <div class="w-64 flex-none"><Sidebar /></div>
    <div class="w-64 flex-1"><Greet /></div>
  </div>
</main>
